package network

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func GetMacAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var address []string
	for _, i := range interfaces {
		a := i.HardwareAddr.String()
		if a != "" {
			address = append(address, a)
		}
	}
	if len(address) == 0 {
		return "", nil
	}
	return address[0], nil
}

func GetLocalIP() net.IP {
	conn, err := net.Dial(`udp`, `8.8.8.8:80`)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func CreateConnection(address, port string) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetPublicIpAddress() string {
	resp, err := http.Get("https://ramziv.com/ip") // https://ip.beget.ru
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
