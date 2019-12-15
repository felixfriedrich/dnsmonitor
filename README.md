# dnsmonitor

* [![Build Status](https://travis-ci.org/felixfriedrich/dnsmonitor.svg?branch=master)](https://travis-ci.org/felixfriedrich/dnsmonitor)
* [![Go Report Card](https://goreportcard.com/badge/github.com/felixfriedrich/dnsmonitor)](https://goreportcard.com/report/github.com/felixfriedrich/dnsmonitor)
* https://github.com/golang-standards/project-layout

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
  -DNS string
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