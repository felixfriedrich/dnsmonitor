package main

import (
	"dnsmonitor/pkg/alerting"
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/dns"
	"dnsmonitor/pkg/monitor"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	okExitCode                      = 0
	fileDoesntExistExitCode         = 1
	wrongCombinationOfFlagsExitCode = 2
	configErrorExitCode             = 3
)

func main() {
	flags := configuration.ParseFlags()
	ok, exitCode := sanityCheckFlags(flags)
	if !ok {
		os.Exit(exitCode)
	}

	c, err := configuration.CreateConfig(flags)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(configErrorExitCode)
	}

	for _, config := range c.Monitors {

		m, err := monitor.CreateMonitor(*config, alerting.NewMailAlerting(config.Alerting.Mail), alerting.New(config), dns.New())
		if err != nil {
			log.Error(err)
		}
		go m.Run(config.Interval, config.Silent)
	}

	select {} // Make this program not terminate in order to keep the go routines running
}

func sanityCheckFlags(flags configuration.Flags) (bool, int) {
	if flags.Version {
		fmt.Printf("dnsmonitor v%s\n", configuration.Version)
		return false, okExitCode
	}

	if flags.ConfigFile != "" {
		_, err := os.Stat(flags.ConfigFile)
		if os.IsNotExist(err) {
			fmt.Println("config file", flags.ConfigFile, "doesn't exist")
			return false, fileDoesntExistExitCode
		}
	}

	if flags.ConfigFile != "" && len(flags.Domains) > 0 {
		fmt.Println("Provide either -configfile OR -domain")
		return false, wrongCombinationOfFlagsExitCode
	}

	return true, -1
}
