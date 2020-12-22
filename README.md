# dnsmonitor

This application monitors domain(s) for changing DNS records and is able to alert on changes via [E-Mail and SMS](../../wiki/Alerting).

* [![Build Status](https://github.com/felixfriedrich/dnsmonitor/workflows/check-commit/badge.svg)](https://github.com/felixfriedrich/dnsmonitor/actions)
* [![Go Report Card](https://goreportcard.com/badge/github.com/felixfriedrich/dnsmonitor)](https://goreportcard.com/report/github.com/felixfriedrich/dnsmonitor)
* [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=felixfriedrich_dnsmonitor&metric=alert_status)](https://sonarcloud.io/dashboard?id=felixfriedrich_dnsmonitor)
* https://trello.com/b/PGkhfOQE/dnsmonitor


## Usage

### Docker

```
$ docker pull ghcr.io/felixfriedrich/dnsmonitor-amd64:latest
$ docker run ghcr.io/felixfriedrich/dnsmonitor-amd64
```

See https://github.com/users/felixfriedrich/packages/container/package/dnsmonitor-amd64


### Build locally

```
$ make build && ./bin/dnsmonitor
```


### Use with config file

```
$ ./bin/dnsmonitor -configfile configs/default.yml
```

See [ConfigFile](../../wiki/ConfigFile) more information on config files.


### Use with command line flags
```
$ ./bin/dnsmonitor -domain www.google.com
$ docker run ghcr.io/felixfriedrich/dnsmonitor-amd64 dnsmonitor -domain www.google.com
```

#### All flags
```
Usage of ./bin/dnsmonitor:
  -configfile string
      config file
  -dns string
      DNS server (default "8.8.8.8")
  -domain value
      domain to check. Can be used multiple times.
  -interval int
      interval in seconds (default 300)
  -mail
      send mail if DNS record changes
  -silent
      silence output
  -sms
      send SMS if DNS record changes
  -vendor value
      Alerting vendor, e.g. 'messagebird'.
  -version
      print version
```

## More information
* [Alerting](../../wiki/Alerting)

## Links
* For generating mocks: https://github.com/maxbrunsfeld/counterfeiter
* Repo layout: https://github.com/golang-standards/project-layout
