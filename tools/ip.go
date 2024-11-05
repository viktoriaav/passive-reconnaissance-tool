package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPInfo struct {
	IP              string `json:"ip"`
	CountryName     string `json:"country_name"`
	CountryCode2    string `json:"country_code2"`
	ISP             string `json:"isp"`
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

func ProcessIPAddress(ipAddress string) (string, error) {
	url := "https://api.iplocation.net/?ip=" + ipAddress

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return "", err
	}

	if ipInfo.ResponseCode != "200" {
		return "", fmt.Errorf("API request failed with response code: %s - %s", ipInfo.ResponseCode, ipInfo.ResponseMessage)
	}

	result := fmt.Sprintf("ISP: %s\nCountry: %s (%s)\n", ipInfo.ISP, ipInfo.CountryName, ipInfo.CountryCode2)
	return result, nil
}
