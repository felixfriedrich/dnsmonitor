# dnsmonitor

This application monitors domain(s) for changing DNS records and is able to alert on changes via [E-Mail and SMS](../../wiki/Alerting).

* [![Build Status](https://github.com/felixfriedrich/dnsmonitor/workflows/check-commit/badge.svg)](https://github.com/felixfriedrich/dnsmonitor/actions)
* [![Go Report Card](https://goreportcard.com/badge/github.com/felixfriedrich/dnsmonitor)](https://goreportcard.com/report/github.com/felixfriedrich/dnsmonitor)

## Usage

Run application:
```
$ make build && ./bin/dnsmonitor -domain www.google.com
```

Command line flags:
```
Usage of ./bin/dnsmonitor:
  -dns string
    	DNS server (default "8.8.8.8")
  -domain value
    	domain to check. Can be used multiple times.
  -interval int
    	interval in seconds (default 1)
  -mail
    	send mail if DNS record changes
  -silent
    	silence output
  -version
    	print version
```

[Alerting](../../wiki/Alerting) options.

## Links
* For generating mocks: https://github.com/maxbrunsfeld/counterfeiter
* Repo layout: https://github.com/golang-standards/project-layout
