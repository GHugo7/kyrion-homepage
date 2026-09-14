package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getNPMProxyHosts() ([]NPMProxyHost, error) {
	godotenv.Load()
	// Etape 1: login
	loginBody := map[string]string{
		"identity": os.Getenv("NPM_IDENTITY"),
		"secret":   os.Getenv("NPM_SECRET"),
	}
	jsonBody, _ := json.Marshal(loginBody)

	resp, err := http.Post("http://192.168.1.69:81/api/tokens", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &tokenResp)

	// Etape 2: GET url
	req, err := http.NewRequest("GET", "http://192.168.1.69:81/api/nginx/proxy-hosts", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tokenResp.Token)

	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()

	var hosts []NPMProxyHost
	bodyBytes2, _ := io.ReadAll(resp2.Body)
	json.Unmarshal(bodyBytes2, &hosts)

	return hosts, nil
}

func matchNPMUrl(containerName string, hosts []NPMProxyHost) string {
	for _, h := range hosts {
		if h.ForwardHost == containerName {
			return "https://" + h.DomainNames[0]
		}
	}
	return ""
}
