package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetHostnames route handler gets the hostnames json array
// and responds with the JSON output
// @TODO - add error handling
func GetHostnames(w http.ResponseWriter, r *http.Request) {
	var hostnames = getHostnamesFromJSON()

	// if err != nil {
	//     respondWithError(w, http.StatusInternalServerError, err.Error())
	//     return
	// }

	respondWithJSON(w, http.StatusOK, hostnames)
}

// GetHostname route handler gets the hostname by Domain
// and responds with the JSON output
// @TODO - add error handling
func GetHostname(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var hostname = getHostnameByDomain(params["domain"])

	respondWithJSON(w, http.StatusOK, hostname)
}

// CreateHostname route handler attempts to add the hostname to
// hostnames array and responds with the updated hostnames JSON array output
// @TODO - add error handling
func CreateHostname(w http.ResponseWriter, r *http.Request) {
	var hostname Hostname
	_ = json.NewDecoder(r.Body).Decode(&hostname)
	var hostnames = createHostname(hostname)

	respondWithJSON(w, http.StatusOK, hostnames)
}

// UpdateHostname route handler attempts to update the hostname in
// hostnames array and responds with the updated hostnames JSON array output
// @TODO - add error handling
func UpdateHostname(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var updatedHostname Hostname
	_ = json.NewDecoder(r.Body).Decode(&updatedHostname)
	updatedHostname.Domain = params["domain"]

	var hostnames = updateHostname(updatedHostname)
	
	respondWithJSON(w, http.StatusOK, hostnames)
}

// DeleteHostname route handler attempts to delete the hostname from
// hostnames array and responds with the updated hostnames JSON array output
// @TODO - add error handling
func DeleteHostname(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var domain = params["domain"]

	var hostnames = deleteHostname(domain)

	respondWithJSON(w, http.StatusOK, hostnames)
}

// TriggerUpdate route handler triggers the attemptIPAddressUpdates job
// and responds with a string message to let the caller know it was triggered
// @TODO - add error handling
func TriggerUpdate(w http.ResponseWriter, r *http.Request) {
	attemptIPAddressUpdates()

	// if err != nil {
	//     respondWithError(w, http.StatusInternalServerError, err.Error())
	//     return
	// }

	fmt.Fprintf(w, "Update Triggered")
}
