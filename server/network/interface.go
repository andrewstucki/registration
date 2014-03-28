package network

import (
	"fmt"
)

const (
	SecurityNone   = 0
	SecurityWEP40  = 1
	SecurityWEP104 = 2
)

type WirelessInterface struct {
	InterfaceName string
	Ssid          string
	Channel       uint
}

func Start(intName, security, netName string, netChannel uint, password string) string {
	var ip string
	var err error
	var ifaceName string
	if intName == "" {
		interfaces := getInterfaces()
		ifaceName = interfaces[0]
	} else {
		ifaceName = intName
	}

	var sec uint
	if security == "WEP40" {
		sec = SecurityWEP40
		//check length of password stuff here
	} else if security == "WEP104" {
		sec = SecurityWEP104
		//check length of password stuff here
	} else {
		sec = SecurityNone
	}

	ip, err = makeNetwork(ifaceName, netName, sec, netChannel, password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Attached network '%s', host ip now '%s'!\n", netName, ip)
	return ip
}

func Cleanup() {
	//todo
}
