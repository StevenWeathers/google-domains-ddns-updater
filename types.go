package main

type Hostnames struct {
    Hostnames []Hostname `json:"hostnames"`
}

type Hostname struct {
    Domain string `json:"domain"`
    Username string `json:"username"`
    Password string `json:"password"`
}