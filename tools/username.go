package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

var (
	results = make(map[string]string)
)

type SocialNetwork struct {
	Name                   string `json:"name"`
	CheckURI               string `json:"check_uri"`
	AccountExistenceString string `json:"account_existence_string"`
}

type SocialNetworksList struct {
	Sites []SocialNetwork `json:"sites"`
}

func checkUsernameExists(username string, socialNetwork SocialNetwork) bool {
	baseURL := socialNetwork.CheckURI
	checkURL := strings.Replace(baseURL, "{account}", url.QueryEscape(username), -1)

	resp, err := http.Get(checkURL)
	if err != nil {
		log.Printf("Error checking %s: %v", socialNetwork.Name, err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body for %s: %v", socialNetwork.Name, err)
			return false
		}
		if strings.Contains(string(body), socialNetwork.AccountExistenceString) {
			return true
		}
	}
	return false
}

func getSocialNetworksList() []SocialNetwork {
	filepath := "./tools/social-networks.json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading social networks list: %v", err)
	}

	var webAccountsList SocialNetworksList
	err = json.Unmarshal(data, &webAccountsList)
	if err != nil {
		log.Fatalf("Error parsing social networks list: %v", err)
	}

	return webAccountsList.Sites
}

func ProcessUsername(username string) (map[string]bool, error) {
	socialNetworks := getSocialNetworksList()

	var wg sync.WaitGroup
	results := make(map[string]bool)
	mu := &sync.Mutex{}

	for _, socialNetwork := range socialNetworks {
		wg.Add(1)
		go func(socialNetwork SocialNetwork) {
			defer wg.Done()
			exists := checkUsernameExists(username, socialNetwork)
			mu.Lock()
			results[socialNetwork.Name] = exists
			mu.Unlock()
		}(socialNetwork)
	}

	wg.Wait()

	return results, nil
}
