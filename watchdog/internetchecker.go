package watchdog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/3rubasa/shwatchdog/config"
)

type InternetCheckerImpl struct {
	cfg *config.InetCheckerConfig
}

func NewInternetChecker(cfg *config.InetCheckerConfig) *InternetCheckerImpl {
	return &InternetCheckerImpl{
		cfg: cfg,
	}
}

func (ic *InternetCheckerImpl) IsInternetAvailable() (bool, error) {
	fmt.Printf("About to send request to check if Internet is available: %s \n", ic.cfg.URL)

	req, err := http.NewRequest(http.MethodGet, ic.cfg.URL, nil)
	if err != nil {
		fmt.Printf("Error while creating request: %s \n", err.Error())
		return false, err
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error while sending request: %s \n", err.Error())
		return false, nil
	}

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("response status is >= 400: %d", resp.StatusCode)
		fmt.Printf("Error: %s \n", err.Error())
		return false, nil
	}

	return true, nil
}
