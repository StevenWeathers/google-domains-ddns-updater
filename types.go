package main

// Hostnames contains hostname(s)
type Hostnames struct {
	Hostnames []Hostname `json:"hostnames"`
}

// Hostname each hostname is a domain with a user/pass to update DDNS
type Hostname struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}
