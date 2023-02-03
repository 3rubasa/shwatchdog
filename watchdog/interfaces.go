package watchdog

type OSServicesProvider interface {
	Reboot() error
}
