package main

import (
	"awesomeProject1/checkdomain"
	"awesomeProject1/sendalert"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type config struct {
	Domains []string `json:"domains"`
}

func loadDomains(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Failed to parse JSON: %v", err)
	}
	var cfg config
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Faild to parse JSON: %v", err)
	}
	return cfg.Domains
}
func main() {
	domains := loadDomains("domain.json")
 for {
	for _, domain := range domains {
		ok, status, err := checkdomain.CheckDomain(domain)
		if !ok {
			msg := fmt.Sprintf("ALERT: %v", err)
			fmt.Println(msg)
			sendalert.SendSlackAlert(msg)
		} else {
			fmt.Printf("%s is UP (status: %v)\n", domain, status)
		}
	}
	time.Sleep(1 * time.Minute)
   }
}
