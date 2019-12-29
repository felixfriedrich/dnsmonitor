# dnsmonitor

* [![Build Status](https://github.com/felixfriedrich/dnsmonitor/workflows/check-commit/badge.svg)](https://github.com/felixfriedrich/dnsmonitor/actions)
* [![Go Report Card](https://goreportcard.com/badge/github.com/felixfriedrich/dnsmonitor)](https://goreportcard.com/report/github.com/felixfriedrich/dnsmonitor)

## Usage

* Configure mail client
```
This application is configured via the environment. The following environment
variables can be used:

KEY                         TYPE       DEFAULT      REQUIRED    DESCRIPTION
DNSMONITOR_MAIL_HOST        String     127.0.0.1
DNSMONITOR_MAIL_PORT        Integer    25
DNSMONITOR_MAIL_USERNAME    String                  true
DNSMONITOR_MAIL_PASSWORD    String                  true
DNSMONITOR_MAIL_FROM        String                  true
DNSMONITOR_MAIL_TO          String                  true
```

* Run application
```
$ ./bin/dnsmonitor -domain www.google.com
```

* Command line flags
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


## Links
* For generating mocks: https://github.com/maxbrunsfeld/counterfeiter
* Repo layout: https://github.com/golang-standards/project-layout
