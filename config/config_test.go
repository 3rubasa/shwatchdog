package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input string = `{
    "watchdog":{
        "inet_checker": {
            "enabled": true,
            "url":"http://google.com",
            "long_period": 1800,
            "short_period": 300
        },
        "vpn_checker": {
            "enabled": true,
            "svc_name": "openvpn",
            "long_period": 300
        }
    } 
}`

func TestConfig(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "testCfgFile.json")
	assert.NoError(t, err)
	fname := file.Name()
	defer func() {
		os.Remove(fname)
	}()

	_, err = file.WriteString(input)
	assert.NoError(t, err)
	file.Close()

	cfg, err := ReadFromFile(fname)
	assert.NoError(t, err)

	assert.Equal(t, cfg.BusinessLogic.Consumer.Enabled, true)
	assert.Equal(t, cfg.BusinessLogic.Consumer.APIKeys, "consumer_key")
	assert.Equal(t, cfg.BusinessLogic.Consumer.Address, "api.thingspeak.com")
	assert.Equal(t, cfg.BusinessLogic.Consumer.Schema, "https")
	assert.Equal(t, cfg.BusinessLogic.Consumer.URI, "update")

	assert.Equal(t, cfg.Watchdog.InetChecker.Enabled, true)
	assert.Equal(t, cfg.Watchdog.InetChecker.URL, "http://google.com")
	assert.Equal(t, cfg.Watchdog.InetChecker.LongPeriod, 1800)
	assert.Equal(t, cfg.Watchdog.InetChecker.ShortPeriod, 300)
	assert.Equal(t, cfg.Watchdog.VPNChecker.Enabled, true)
	assert.Equal(t, cfg.Watchdog.VPNChecker.SvcName, "openvpn")
	assert.Equal(t, cfg.Watchdog.VPNChecker.LongPeriod, 300)

	assert.Equal(t, cfg.WeatherProvider.Enabled, true)

	assert.Equal(t, cfg.ForecastProvider.Enabled, true)
}
