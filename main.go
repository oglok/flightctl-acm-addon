package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oglok/flightctl-acm-addon/pkg/initfc"
)

type Device struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
}

type DeviceList struct {
	APIVersion string   `json:"apiVersion"`
	Items      []Device `json:"items"`
	Kind       string   `json:"kind"`
}

func decodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func main() {
	configMap := initfc.LoadFlightCtlConfigMap()

	certData, _ := decodeBase64(configMap["authentication.client-certificate-data"])
	keyData, _ := decodeBase64(configMap["authentication.client-key-data"])
	caData, _ := decodeBase64(configMap["service.certificate-authority-data"])
	server := configMap["service.server"]

	cert, _ := tls.X509KeyPair(certData, keyData)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caData)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		},
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(string(server) + "/api/v1/devices")

	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var devices DeviceList
	json.Unmarshal(body, &devices)

	for _, device := range devices.Items {
		fmt.Println("Device Name:", device.Metadata.Name)
	}
}
