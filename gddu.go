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
)

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
        for i := 0; i < len(hostnames.Hostnames); i++ {
            var hostname = hostnames.Hostnames[i]
            
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

func getHostnamesFromJson() Hostnames {
    var hostnames Hostnames
    var byteValue = getJsonFile("/data/hostnames.json")
	json.Unmarshal(byteValue, &hostnames)

    return hostnames
}

func main() {
    var cadence string = os.Getenv("CADENCE")
    if cadence == "" {
        log.Fatalln("CADENCE not defined")
    }
    
    var hostnames = getHostnamesFromJson()

    attemptIpAddressUpdates(hostnames) // run at start for good measure
    c := cron.New()
    c.AddFunc(cadence, func() { attemptIpAddressUpdates(hostnames) })
    c.Start()
    select{} // This guarantees this program never exits so cron can keep running as per the cron interval.
}