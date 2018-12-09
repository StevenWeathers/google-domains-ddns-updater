package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// getJSONFile retreives the hostnames JSON file from the filesystem
// and attempts to read it into a []byte and return the value
func getJSONFile() []byte {
	hostnamesFile, err := os.Open(getJSONFilePath())

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

// writeJSONFile takes an interface payload
// and writes it to the hostnames JSON file in the filesystem
func writeJSONFile(payload interface{}) {
	b, _ := json.MarshalIndent(payload, "", " ")

	_ = ioutil.WriteFile(getJSONFilePath(), b, 0644)
}

// getHostnamesFromJSON gets the hostnames JSON file from getJSONFile
// and converts the bytevalue to JSON
func getHostnamesFromJSON() Hostnames {
	var hostnames Hostnames
	var byteValue = getJSONFile()
	json.Unmarshal(byteValue, &hostnames)

	return hostnames
}

// getHostnameByDomain takes a domain string
// then retrieves the hostname that matches that domain
// from the hostnames JSON file and returns it
func getHostnameByDomain(domain string) *Hostname {
	var hostnames = getHostnamesFromJSON()
	var result *Hostname

	for _, hostname := range hostnames.Hostnames {
		if hostname.Domain == domain {
			result = &hostname
			break
		}
	}

	return result
}

// createHostname takes a hostname interface
// and writes it to the hostnames JSON file
// returning the updated hostnames
func createHostname(hostname Hostname) Hostnames {
	var hostnames = getHostnamesFromJSON()

	hostnames.Hostnames = append(hostnames.Hostnames, hostname)

	writeJSONFile(hostnames)

	return hostnames
}

// updateHostname takes a Hostname interface
// and retrieves the hostnames JSON, then updates the matching
// hostname with the passed values, returning the updated hostnames
func updateHostname(updatedHostname Hostname) Hostnames {
	var hostnames = getHostnamesFromJSON()

	for index, hostname := range hostnames.Hostnames {
		if hostname.Domain == updatedHostname.Domain {
			hostnames.Hostnames[index] = updatedHostname
			break
		}
	}

	writeJSONFile(hostnames)

	return hostnames
}

// deleteHostname takes a domain string
// and looks for matching hostname in hostnames JSON
// then removes it, writing the updated hostnames to JSON file
// and returns the updated hostnames
func deleteHostname(domain string) Hostnames {
	var hostnames = getHostnamesFromJSON()
	
	for index, hostname := range hostnames.Hostnames {
		if hostname.Domain == domain {
			hostnames.Hostnames = append(hostnames.Hostnames[:index], hostnames.Hostnames[index+1:]...)
			break
		}
	}

	writeJSONFile(hostnames)

	return hostnames
}