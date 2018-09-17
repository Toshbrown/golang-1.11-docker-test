package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	getExternalIP()
}

func getExternalIP() (string, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	response, err := netClient.Get("http://whatismyip.akamai.com/")
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return "", err
	}
	ip, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return "", err
	}
	response.Body.Close()
	fmt.Println("External IP found " + string(ip))
	return string(ip), nil
}
