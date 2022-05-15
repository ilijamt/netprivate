netprivate
----------

![Build Status](https://github.com/ilijamt/netprivate/actions/workflows/go.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/stretchr/testify)](https://goreportcard.com/report/github.com/ilijamt/netprivate) [![PkgGoDev](https://pkg.go.dev/badge/github.com/stretchr/testify)](https://pkg.go.dev/github.com/ilijamt/netprivate)


A simple library that helps you check if an IP address is in the private defined range of the network. 

The library works with:

* IPv4
* IPv6

### Example usage

```go
if netprivate.Is(net.ParseIP("127.0.0.1")) {
    fmt.Println("this is a private network IP address")
}
```
