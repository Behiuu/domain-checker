package checkdomain

import (
	"fmt"
	"net/http"
	"time"
)

func CheckDomain(domain string) (bool, int, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(domain)
	if err != nil {
		return false, 0, fmt.Errorf("%s is DOWN , error code: %v", domain, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, resp.StatusCode, nil
	}
	return false, resp.StatusCode, fmt.Errorf("%s is DOWN, status CODE %d", domain, resp.StatusCode)
}
