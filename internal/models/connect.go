package models

import "net"

type Connect struct {
	Connection      net.Conn `json:"-"`
	Hostname        string   `json:"hostname"`
	Username        string   `json:"username"`
	UserID          string   `json:"userId"`
	OSName          string   `json:"os_name"`
	MacAddress      string   `json:"mac_address"`
	LocalIPAddress  string   `json:"localIpAddress"`
	PublicIpAddress string   `json:"publicIpAddress"`
	Port            string   `json:"port"`
	FetchedUnix     int64    `json:"fetchedUnix"`
}

type Message struct {
	Command   string
	Data      []byte
	MasterKey []byte
	Error     Error
}

type Error struct {
	HasError bool
	Message  string
}

type Credentials struct {
	OriginUrl     string `json:"origin_url" db:"origin_url"`
	UsernameValue string `json:"usernam_value" db:"username_value"`
	PasswordValue string `json:"password_value" db:"password_value"`
}
