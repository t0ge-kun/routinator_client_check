package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	Connections int    `json:"connections"`
	Description string `json:"description,omitempty"`
}

type Rtr struct {
	Clients map[string]Client `json:"clients"`
}

type Response struct {
	Rtr Rtr `json:"rtr"`
}

type Config struct {
	ApiURL  string `json:"api_url"`
	ApiPort string `json:"api_port"`
}

type Output struct {
	Clients      map[string]Client `json:"clients"`
	TotalConnect int               `json:"total_connections"`
}

type IPApiResponse struct {
	Org string `json:"org"`
}

func getOrganization(ip string) (string, error) {
	// Check if the IP address is a local address
	if isLocalAddress(ip) {
		return "local address", nil
	}

	resp, err := http.Get("https://ipinfo.io/" + ip + "/org")
	if err != nil || resp.StatusCode != http.StatusOK {
		return getOrganizationFromIPApi(ip)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return getOrganizationFromIPApi(ip)
	}

	org := strings.TrimSpace(string(body))
	if org == "" || org == "Unknown" {
		return getOrganizationFromIPApi(ip)
	}

	return org, nil
}

func getOrganizationFromIPApi(ip string) (string, error) {
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return "Unknown", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Unknown", err
	}

	var ipApiResponse IPApiResponse
	err = json.Unmarshal(body, &ipApiResponse)
	if err != nil {
		return "Unknown", err
	}

	return ipApiResponse.Org, nil
}

func isLocalAddress(ip string) bool {
	ipNet := net.ParseIP(ip)
	if ipNet == nil {
		return false
	}

	localBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
	}

	for _, block := range localBlocks {
		_, subnet, _ := net.ParseCIDR(block)
		if subnet.Contains(ipNet) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Get("http://" + config.ApiURL + ":" + config.ApiPort + "/api/v1/status")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	filtered := make(map[string]Client)
	totalConnections := 0
	for ip, client := range data.Rtr.Clients {
		if client.Connections > 0 {
			org, err := getOrganization(ip)
			if err != nil {
				fmt.Println(err)
				return
			}
			client.Description = org
			filtered[ip] = client
			totalConnections += client.Connections
		}
	}

	outputData := Output{
		Clients:      filtered,
		TotalConnect: totalConnections,
	}

	output, err := json.Marshal(outputData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
