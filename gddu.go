package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
    "github.com/rdegges/go-ipify"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

var jsonFilePath string

// getJSONFilePath gets the jsonFilePath string
func getJSONFilePath() string {
	return jsonFilePath
}

// setJSONFilePath sets the jsonFilePath string
func setJSONFilePath(path string) {
	jsonFilePath = path
}

// getEnv gets environment variable matching key string
// and if it finds none uses fallback string
// returning either the matching or fallback string
func getEnv(key string, fallback string) string {
	var result = os.Getenv(key)

	if result == "" {
		result = fallback
	}

	return result
}

// postToGoogleDNS makes a call to GoogleDomains DDNS with provided hostname details
// and the current WAN IP as provided, then attempts to read the response from Google
// and logs the result or error
func postToGoogleDNS(username string, password string, domain string, ip string) {
	log.Println("Domain: ", domain)
	var endpoint = fmt.Sprintf("https://%s:%s@domains.google.com/nic/update?hostname=%s&myip=%s", username, password, domain, ip)

	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	log.Println(string(body))
}

// attemptIPAddressUpdates gets the hostnames from JSON file
// and loops through calling postToGoogleDNS to attempt
// to update the record on Google Domains DDNS
func attemptIPAddressUpdates() {
	var hostnames = getHostnamesFromJSON()
	ip, err := ipify.GetIp()

	if err != nil {
		log.Println(err)
	} else {
		for _, hostname := range hostnames.Hostnames {
			postToGoogleDNS(hostname.Username, hostname.Password, hostname.Domain, ip)
		}
	}
}

// func respondWithError(w http.ResponseWriter, code int, message string) {
//     respondWithJSON(w, code, map[string]string{"error": message})
// }

// responsdWithJSON takes a payload and writes the response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	setJSONFilePath(getEnv("JSONPATH", "/data/hostnames.json"))
	var cadence = getEnv("CADENCE", "@hourly")
	var listenPort = fmt.Sprintf(":%s", getEnv("PORT", "8000"))
	var staticDir = getEnv("STATICDIR", "/static")
	var css = fmt.Sprintf("%s/css/", staticDir)
	var js = fmt.Sprintf("%s/js/", staticDir)
	var entry = fmt.Sprintf("%s/index.html", staticDir)

	c := cron.New()
	c.AddFunc(cadence, attemptIPAddressUpdates)
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
		Handler: router,
		Addr:    listenPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
