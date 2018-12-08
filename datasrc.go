package main

import (
    "os"
    "log"
    "io/ioutil"
    "encoding/json"
)

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

    _ = ioutil.WriteFile(getJsonFilePath(), b, 0644)
}

func getHostnamesFromJson() Hostnames {
    var hostnames Hostnames
    var byteValue = getJsonFile(getJsonFilePath())
	json.Unmarshal(byteValue, &hostnames)

    return hostnames
}