package watchdog

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/3rubasa/shagent/config"
	"github.com/3rubasa/shagent/watchdog/mockinternetchecker"
	"github.com/3rubasa/shagent/watchdog/mockosservicesprovider"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var conf string = `{
	"inet_checker": {
		"enabled": true,
		"url":"http://google.com",
		"long_period": 1,
		"short_period": 1
	},
	"vpn_checker": {
		"enabled": false,
		"svc_name": "openvpn",
		"long_period": 1,
		"short_period": 300
	}
}`

func Test_InetNotAvail(t *testing.T) {
	if os.Getenv("SH_RUN_ALL_TESTS") != "1" {
		t.Skip("Long test, skipping due to SH_RUN_ALL_TEST != 1 ...")
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	osSvcs := mockosservicesprovider.NewMockOSServicesProvider(mockCtrl)
	inetChecker := mockinternetchecker.NewMockInternetChecker(mockCtrl)

	var cfg *config.WatchdogConfig
	err := json.NewDecoder(strings.NewReader(conf)).Decode(&cfg)
	assert.NoError(t, err)

	wd := New(cfg, osSvcs, inetChecker)

	inetChecker.EXPECT().IsInternetAvailable().Return(false, nil).Times(3)
	osSvcs.EXPECT().Reboot().Return(nil).Times(1)

	err = wd.Start()
	assert.NoError(t, err)

	time.Sleep(5 * time.Second)
}

func Test_InetAvail(t *testing.T) {
	if os.Getenv("SH_RUN_ALL_TESTS") != "1" {
		t.Skip("Long test, skipping due to SH_RUN_ALL_TEST != 1 ...")
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	osSvcs := mockosservicesprovider.NewMockOSServicesProvider(mockCtrl)
	inetChecker := mockinternetchecker.NewMockInternetChecker(mockCtrl)

	var cfg *config.WatchdogConfig
	err := json.NewDecoder(strings.NewReader(conf)).Decode(&cfg)
	assert.NoError(t, err)

	wd := New(cfg, osSvcs, inetChecker)

	inetChecker.EXPECT().IsInternetAvailable().Return(true, nil).MinTimes(1)
	osSvcs.EXPECT().Reboot().Return(nil).Times(0)

	err = wd.Start()
	assert.NoError(t, err)

	time.Sleep(5 * time.Second)
}

func Test_InetAvail_Error(t *testing.T) {
	if os.Getenv("SH_RUN_ALL_TESTS") != "1" {
		t.Skip("Long test, skipping due to SH_RUN_ALL_TEST != 1 ...")
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	osSvcs := mockosservicesprovider.NewMockOSServicesProvider(mockCtrl)
	inetChecker := mockinternetchecker.NewMockInternetChecker(mockCtrl)

	var cfg *config.WatchdogConfig
	err := json.NewDecoder(strings.NewReader(conf)).Decode(&cfg)
	assert.NoError(t, err)

	wd := New(cfg, osSvcs, inetChecker)

	inetChecker.EXPECT().IsInternetAvailable().Return(true, errors.New("dummy error")).Times(1)
	osSvcs.EXPECT().Reboot().Return(nil).Times(1)

	err = wd.Start()
	assert.NoError(t, err)

	time.Sleep(5 * time.Second)
}

func Test_InetAvail_Intermittent(t *testing.T) {
	if os.Getenv("SH_RUN_ALL_TESTS") != "1" {
		t.Skip("Long test, skipping due to SH_RUN_ALL_TEST != 1 ...")
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	osSvcs := mockosservicesprovider.NewMockOSServicesProvider(mockCtrl)
	inetChecker := mockinternetchecker.NewMockInternetChecker(mockCtrl)

	var cfg *config.WatchdogConfig
	err := json.NewDecoder(strings.NewReader(conf)).Decode(&cfg)
	assert.NoError(t, err)

	wd := New(cfg, osSvcs, inetChecker)

	counter := 0

	inetChecker.EXPECT().IsInternetAvailable().DoAndReturn(func() (bool, error) {
		switch counter {
		case 0:
			counter++
			return false, nil
		case 1:
			counter++
			return true, nil
		case 2:
			counter++
			return true, nil
		case 3:
			counter++
			return false, nil
		case 4:
			counter++
			return false, nil
		case 5:
			counter++
			return false, nil
		}

		return false, nil
	}).Times(6)

	osSvcs.EXPECT().Reboot().Return(nil).Times(1)

	err = wd.Start()
	assert.NoError(t, err)

	time.Sleep(10 * time.Second)
}
