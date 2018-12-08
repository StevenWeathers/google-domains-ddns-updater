package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
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

// @TODO - add error handling
func TriggerUpdate(w http.ResponseWriter, r *http.Request) {
    var hostnames = getHostnamesFromJson()

    attemptIpAddressUpdates(hostnames)
    
    // if err != nil {
    //     respondWithError(w, http.StatusInternalServerError, err.Error())
    //     return
    // }

    fmt.Fprintf(w, "Update Triggered")
}