# dnsmonitor

* [![Build Status](https://travis-ci.org/felixfriedrich/dnsmonitor.svg?branch=master)](https://travis-ci.org/felixfriedrich/dnsmonitor)
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