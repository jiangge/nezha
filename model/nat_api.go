package model

type NATForm struct {
	Name     string `json:"name,omitempty"`
	ServerID uint64 `json:"server_id,omitempty"`
	Host     string `json:"host,omitempty"`
	Domain   string `json:"domain,omitempty"`
}