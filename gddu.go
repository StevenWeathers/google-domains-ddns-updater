package main

import (
    "fmt"
    "os"
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
    "github.com/rdegges/go-ipify"
    "github.com/robfig/cron"
    "github.com/gorilla/mux"
)

var jsonFilePath string

type Hostnames struct {
    Hostnames []Hostname `json:"hostnames"`
}

type Hostname struct {
    Domain string `json:"domain"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func postToGoogleDns(username string, password string, domain string, ip string) {
    var endpoint string = fmt.Sprintf("https://%s:%s@domains.google.com/nic/update?hostname=%s&myip=%s", username, password, domain, ip)
    resp, err := http.Get(endpoint)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}

func attemptIpAddressUpdates(hostnames Hostnames) {
    ip, err := ipify.GetIp()

    if err != nil {
        log.Println(err)
    } else {
        for _, hostname := range hostnames.Hostnames {
            postToGoogleDns(hostname.Username, hostname.Password, hostname.Domain, ip)
        }
    }
}

func getJsonFile(path string) []byte {
    hostnamesFile, err := os.Open(path)

    if err != nil {
        log.Fatalln(err)
    }

    defer hostnamesFile.Close() // defer the closing of our jsonFile so that we can parse it later on

    byteValue, readErr := ioutil.ReadAll(hostnamesFile)

    if readErr != nil {
        log.Fatalln(err)
    }

    return byteValue
}

func writeJsonFile(payload interface{}) {
    b, _ := json.MarshalIndent(payload, "", " ")

    _ = ioutil.WriteFile(jsonFilePath, b, 0644)
}

func getHostnamesFromJson() Hostnames {
    var hostnames Hostnames
    var byteValue = getJsonFile(jsonFilePath)
	json.Unmarshal(byteValue, &hostnames)

    return hostnames
}

// func respondWithError(w http.ResponseWriter, code int, message string) {
//     respondWithJSON(w, code, map[string]string{"error": message})
// }

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

// @TODO - add error handling
func GetHostnames(w http.ResponseWriter, r *http.Request) {
    var hostnames = getHostnamesFromJson()
    
    // if err != nil {
    //     respondWithError(w, http.StatusInternalServerError, err.Error())
    //     return
    // }

    respondWithJSON(w, http.StatusOK, hostnames)
}

// @TODO - add error handling
func GetHostname(w http.ResponseWriter, r *http.Request) {
    var hostnames = getHostnamesFromJson()
    params := mux.Vars(r)
    for _, hostname := range hostnames.Hostnames {
        if hostname.Domain == params["domain"] {
            json.NewEncoder(w).Encode(hostname)
            return
        }
    }
    respondWithJSON(w, http.StatusOK, &Hostname{})
}

// @TODO - add error handling
func CreateHostname(w http.ResponseWriter, r *http.Request) {
    var hostnames = getHostnamesFromJson()
    var hostname Hostname
    _ = json.NewDecoder(r.Body).Decode(&hostname)
    hostnames.Hostnames = append(hostnames.Hostnames, hostname)

    writeJsonFile(hostnames)
    respondWithJSON(w, http.StatusOK, hostnames)
}

// @TODO - add error handling
func UpdateHostname(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var hostnames = getHostnamesFromJson()
    var updatedHostname Hostname
    _ = json.NewDecoder(r.Body).Decode(&updatedHostname)
    updatedHostname.Domain = params["domain"]

    for index, hostname := range hostnames.Hostnames {
        if hostname.Domain == updatedHostname.Domain {
            hostnames.Hostnames[index] = updatedHostname
        }
    }

    writeJsonFile(hostnames)
    respondWithJSON(w, http.StatusOK, hostnames)
}

// @TODO - add error handling
func DeleteHostname(w http.ResponseWriter, r *http.Request) {
    var hostnames = getHostnamesFromJson()
    params := mux.Vars(r)
    for index, hostname := range hostnames.Hostnames {
        if hostname.Domain == params["domain"] {
            hostnames.Hostnames = append(hostnames.Hostnames[:index], hostnames.Hostnames[index+1:]...)
            break
        }
    }
    writeJsonFile(hostnames)
    respondWithJSON(w, http.StatusOK, hostnames)
}

func main() {
    var cadence string = os.Getenv("CADENCE")
    if cadence == "" {
        log.Fatalln("CADENCE not defined")
    }

    jsonFilePath = "/data/hostnames.json" // @TODO - read this from env
    
    var hostnames = getHostnamesFromJson()

    attemptIpAddressUpdates(hostnames) // run at start for good measure, @todo add endpoint to maunally trigger this as well
    c := cron.New()
    c.AddFunc(cadence, func() { attemptIpAddressUpdates(hostnames) })
    c.Start()
    
    router := mux.NewRouter()
    router.HandleFunc("/hostnames", GetHostnames).Methods("GET")
    router.HandleFunc("/hostnames/{domain}", GetHostname).Methods("GET")
    router.HandleFunc("/hostnames", CreateHostname).Methods("POST")
    router.HandleFunc("/hostnames/{domain}", UpdateHostname).Methods("PUT")
    router.HandleFunc("/hostnames/{domain}", DeleteHostname).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router)) // @TODO - abstract port to env
}