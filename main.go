package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rdegges/go-ipify"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

var (
	version    = "dev"
	embedUseOS = true
)

// ServerConfig holds server global config values
type ServerConfig struct {
	// port the application server will listen on
	ListenPort string
	// the app version
	Version string
	// PathPrefix allows the application to be run on a shared domain
	PathPrefix string
}

type server struct {
	config *ServerConfig
	router *mux.Router
}

var jsonFilePath string

// getJSONFilePath gets the jsonFilePath string
func getJSONFilePath() string {
	return jsonFilePath
}

// setJSONFilePath sets the jsonFilePath string
func setJSONFilePath(path string) {
	jsonFilePath = path
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

// respondWithJSON takes a payload and writes the response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	log.Println("Google Domains DDNS Updater version " + version)

	embedUseOS = len(os.Args) > 1 && os.Args[1] == "live"

	InitConfig()

	pathPrefix := viper.GetString("http.path_prefix")
	router := mux.NewRouter()

	if pathPrefix != "" {
		router = router.PathPrefix(pathPrefix).Subrouter()
	}

	setJSONFilePath(viper.GetString("config.json_path"))
	var cadence = viper.GetString("config.cadence")

	c := cron.New()
	c.AddFunc(cadence, attemptIPAddressUpdates)
	c.Start()

	s := &server{
		config: &ServerConfig{
			ListenPort: viper.GetString("http.port"),
			Version:    version,
			PathPrefix: pathPrefix,
		},
		router: router,
	}

	s.routes()

	srv := &http.Server{
		Handler: s.router,
		Addr:    fmt.Sprintf(":%s", s.config.ListenPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Access the WebUI via 127.0.0.1:" + s.config.ListenPort)

	log.Fatal(srv.ListenAndServe())
}
