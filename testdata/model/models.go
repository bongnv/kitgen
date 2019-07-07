package model

// Profile ...
type Profile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

// Address ...
type Address struct {
	ID       string `json:"id"`
	Location string `json:"location,omitempty"`
}
