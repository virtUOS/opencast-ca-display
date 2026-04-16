package endpoints

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type NetworkStatus struct {
	Interfaces []NetInterface `json:"interfaces"`
	Connected  bool           `json:"connected"`
	Hostname   string         `json:"hostname"`
}

type NetInterface struct {
	Name   string   `json:"name"`
	Adress []string `json:"addr"`
	MAC    string   `json:"mac_adress"`
	Flags  string   `json:"flags"`
}

func networkEndpoint(c *gin.Context) {
	var net_status NetworkStatus
	net_interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln("Network devices could not be loaded.")
		return
	}
	for _, net_inter := range net_interfaces {
		addrs, err := net_inter.Addrs()
		var addrs_str []string
		if err == nil {
			for _, a := range addrs {
				addrs_str = append(addrs_str, a.String())
			}
		}
		inter := NetInterface{Name: net_inter.Name, MAC: net_inter.HardwareAddr.String(), Adress: addrs_str, Flags: net_inter.Flags.String()}
		net_status.Interfaces = append(net_status.Interfaces, inter)
	}
	client := &http.Client{Timeout: time.Duration(localConfig.Timeout * int(time.Millisecond))}
	url := localConfig.Opencast.URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadGateway, nil)
		return
	}
	// req.SetBasicAuth(localConfig.Opencast.Username, localConfig.Opencast.Password)
	_, err = client.Do(req)
	if err != nil {
		net_status.Connected = false
	} else {
		net_status.Connected = true
	}
	net_status.Hostname, err = os.Hostname()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, net_status)
}
