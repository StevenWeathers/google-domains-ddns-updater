package main

import (
    "fmt"
    "strings"
    "os"
    "net/http"
	"log"
	"io/ioutil"
    "github.com/rdegges/go-ipify"
    "github.com/robfig/cron"
)

func postToGoogleDns(finalEndpoint string) {
    resp, err := http.Get(finalEndpoint)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}

func attemptIpAddressUpdates(hostnames []string, gdduEndpoint string) {
    ip, err := ipify.GetIp()

    if err != nil {
        log.Println(err)
    } else {
        for _, hostname := range hostnames {
            var finalEndpoint = fmt.Sprintf("%s?hostname=%s&myip=%s", gdduEndpoint, hostname, ip)
            postToGoogleDns(finalEndpoint)
        }
    }
}

func main() {
    var domains string = os.Getenv("DOMAINS")
    var username string = os.Getenv("USERNAME")
    var password string = os.Getenv("PASSWORD")
    var cadence string = os.Getenv("CADENCE")
   
    if domains == "" {        
        log.Fatalln("DOMAINS not defined, must be comma separated list")
    }

    if username == "" {
        log.Fatalln("USERNAME not defined")
    }

    if password == "" {
        log.Fatalln("PASSWORD not defined")
    }

    if cadence == "" {
        log.Fatalln("CADENCE not defined")
    }

    var gdduEndpoint string = fmt.Sprintf("https://%s:%s@domains.google.com/nic/update", username, password)
    var hostnames []string = strings.Split(domains, ",")

    attemptIpAddressUpdates(hostnames, gdduEndpoint) // run at start for good measure
    c := cron.New()
    c.AddFunc(cadence, func() { attemptIpAddressUpdates(hostnames, gdduEndpoint) })
    c.Start()
    select{} // This guarantees this program never exits so cron can keep running as per the cron interval.
}