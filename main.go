package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type domainList struct {
	Name      string      `json:"name"`
	DeletedAt interface{} `json:"deleted_at"`
	Comment   string      `json:"comment"`
	CreatedAt time.Time   `json:"created_at"`
	Version   int         `json:"version"`
	Locked    bool        `json:"locked"`
	UpdatedAt time.Time   `json:"updated_at"`
	ServiceID string      `json:"service_id"`
}

type fastlyService struct {
	Version    int       `json:"version"`
	UpdatedAt  time.Time `json:"updated_at"`
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	Versions   versions  `json:"versions"`
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Comment    string    `json:"comment"`
	CustomerID string    `json:"customer_id"`
}

type versions []struct {
	Comment   string      `json:"comment"`
	Active    bool        `json:"active"`
	Testing   bool        `json:"testing"`
	ServiceID string      `json:"service_id"`
	CreatedAt time.Time   `json:"created_at"`
	Deployed  bool        `json:"deployed"`
	Number    int         `json:"number"`
	Locked    bool        `json:"locked"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
	Staging   bool        `json:"staging"`
}

type backend struct {
	Locked              bool        `json:"locked"`
	ServiceID           string      `json:"service_id"`
	OverrideHost        interface{} `json:"override_host"`
	MaxTLSVersion       interface{} `json:"max_tls_version"`
	MaxConn             int         `json:"max_conn"`
	ConnectTimeout      int         `json:"connect_timeout"`
	Ipv4                string      `json:"ipv4"`
	Version             int         `json:"version"`
	MinTLSVersion       interface{} `json:"min_tls_version"`
	ErrorThreshold      int         `json:"error_threshold"`
	SslClientKey        interface{} `json:"ssl_client_key"`
	Shield              string      `json:"shield"`
	Comment             string      `json:"comment"`
	AutoLoadbalance     bool        `json:"auto_loadbalance"`
	UseSsl              bool        `json:"use_ssl"`
	FirstByteTimeout    int         `json:"first_byte_timeout"`
	SslCheckCert        bool        `json:"ssl_check_cert"`
	BetweenBytesTimeout int         `json:"between_bytes_timeout"`
	Hostname            interface{} `json:"hostname"`
	DeletedAt           interface{} `json:"deleted_at"`
	Healthcheck         interface{} `json:"healthcheck"`
	Name                string      `json:"name"`
	Port                int         `json:"port"`
	SslCaCert           interface{} `json:"ssl_ca_cert"`
	SslClientCert       interface{} `json:"ssl_client_cert"`
	Address             string      `json:"address"`
	Weight              int         `json:"weight"`
	SslCiphers          interface{} `json:"ssl_ciphers"`
	RequestCondition    string      `json:"request_condition"`
	UpdatedAt           time.Time   `json:"updated_at"`
	Ipv6                interface{} `json:"ipv6"`
	CreatedAt           time.Time   `json:"created_at"`
	SslCertHostname     interface{} `json:"ssl_cert_hostname"`
	ClientCert          interface{} `json:"client_cert"`
	SslSniHostname      interface{} `json:"ssl_sni_hostname"`
	SslHostname         interface{} `json:"ssl_hostname"`
}

func main() {
	filePath := "./services-20200226.json"
	fmt.Printf("// reading file %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}
	// fmt.Println(file)
	fmt.Println("// defining array of struct FastlyService")
	var svc []fastlyService

	err2 := json.Unmarshal(file, &svc)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}

	fmt.Println("// loop over array of structs of FastlyService")
	for k := range svc {
		ver := strconv.Itoa(svc[k].Version)
		// fmt.Printf("Service '%s' Name: '%s' Updated: '%s' Active: '%v'\n", svc[k].ID, svc[k].Name, svc[k].UpdatedAt, ver)
		// fmt.Printf("The group is '%s' and description: %s\n", svc[k].UpdatedAt, svc[k].Comment)
		// fmt.Printf("Service '%s' Name: '%s' Updated: '%s' Comment: '%s'\n", svc[k].ID, svc[k].Name, svc[k].UpdatedAt, svc[k].Comment)
		// fmt.Printf("%s, %s, %s, %s, %v\n", svc[k].Name, svc[k].ID, svc[k].UpdatedAt, svc[k].Comment, ver)
		// fmt.Printf("%#v", svc)
		var Name = (svc[k].Name)
		// fmt.Println(Name)
		//  ___Prepare to call API to gdt info on services
		var url = fmt.Sprint("https://api.fastly.com/service/" + svc[k].ID + "/version/" + ver)
		// getDomain(url, Name)
		getBackend(url, Name)
		time.Sleep(3 + time.Second)
	}
}

// Name is the fastly service name
var Name string
var url string

func getDomain(url string, Name string) {
	// Create request object.
	req, err := http.NewRequest("GET", url+"/domain", nil)

	// Set the header in the request.
	req.Header.Set("Fastly-Key", os.Getenv("FASTLY_API_TOKEN"))

	// Execute the request.
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error getting request request.\n[ERRO] -", err)
	}
	defer resp.Body.Close()

	// fmt.Println(" Got response from server", resp.Body)
	// read json http response
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var jsonData []domainList

	err = json.Unmarshal(body, &jsonData) // here!

	if err != nil {
		fmt.Println("errored on unmarshal")
		panic(err)
	}

	// test struct data
	//  fmt.Println(jsonData)
	// fmt.Println("___  Service___")
	for n := range jsonData {
		fmt.Println(jsonData[n].ServiceID, jsonData[n].Name, Name, jsonData[n].UpdatedAt)
	}
	fmt.Println("--Done--")

}

func getBackend(url string, Name string) {
	// fmt.Println(url)
	// Create request object.
	req, err := http.NewRequest("GET", url+"/backend", nil)

	// Set the header in the request.
	req.Header.Set("Fastly-Key", os.Getenv("FASTLY_API_TOKEN"))

	// Execute the request.
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error getting request request.\n[ERRO] -", err)
	}
	defer resp.Body.Close()

	// fmt.Println(" Got response from server", resp.Body)
	// read json http response
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var jsonData []backend

	err = json.Unmarshal(body, &jsonData) // here!

	if err != nil {
		fmt.Println("errored on unmarshal")
		panic(err)
	}

	// test struct data
	//  fmt.Println(jsonData)
	// fmt.Println("___  Backend  ___")
	for n := range jsonData {
		fmt.Printf("%s, %s, %s, %s, %s\n", Name, jsonData[n].ServiceID, jsonData[n].Name, jsonData[n].Shield, jsonData[n].UpdatedAt)
	}

}
