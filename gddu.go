package main

import (
    "os"
    "fmt"
    "log"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/rdegges/go-ipify"
    "github.com/robfig/cron"
    "github.com/gorilla/mux"
)

var jsonFilePath string

func getJsonFilePath() string {
    return jsonFilePath
}

func setJsonFilePath(path string) {
    jsonFilePath = path
}

func getEnv(key string, fallback string) string {
    var result = os.Getenv(key)
    if result == "" {
        result = fallback
    }
    return result
}

func postToGoogleDns(username string, password string, domain string, ip string) {
    log.Println("Domain: ", domain)
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

// func respondWithError(w http.ResponseWriter, code int, message string) {
//     respondWithJSON(w, code, map[string]string{"error": message})
// }

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func main() {
    setJsonFilePath(getEnv("JSONPATH", "/data/hostnames.json"))
    var cadence string = getEnv("CADENCE", "@hourly")
    var listenPort string = fmt.Sprintf(":%s", getEnv("PORT", "8000"))
    var css string = "/static/css/"
    var js string = "/static/js/"
    var entry string = "/static/index.html"
   
    var hostnames = getHostnamesFromJson()

    attemptIpAddressUpdates(hostnames) // run at start for good measure, @todo add endpoint to maunally trigger this as well
    c := cron.New()
    c.AddFunc(cadence, func() { attemptIpAddressUpdates(hostnames) })
    c.Start()
    
    router := mux.NewRouter()
    router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(css))))
    router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(js))))
    router.HandleFunc("/hostnames", GetHostnames).Methods("GET")
    router.HandleFunc("/hostnames/{domain}", GetHostname).Methods("GET")
    router.HandleFunc("/hostnames", CreateHostname).Methods("POST")
    router.HandleFunc("/hostnames/{domain}", UpdateHostname).Methods("PUT")
    router.HandleFunc("/hostnames/{domain}", DeleteHostname).Methods("DELETE")
    router.HandleFunc("/triggerUpdate", TriggerUpdate).Methods("GET")
    router.PathPrefix("/").HandlerFunc(IndexHandler(entry))

    srv := &http.Server{
        Handler:      router,
        Addr:         listenPort,
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}