package app

import (
	"flag"
	"fmt"
	"strings"
)

const (
	PortFlag                = "port"
	AllowedOriginFlag       = "allowed-origin"
	PrintUnknownOriginsFlag = "print-unknown-origins"
)

type allowedOrigin []string

func (i *allowedOrigin) String() string {
	return fmt.Sprint(*i)
}

func (i *allowedOrigin) Set(value string) error {
	for _, origin := range strings.Split(value, ",") {
		*i = append(*i, origin)
	}
	return nil
}

var port = "8080"
var allowedOrigins allowedOrigin
var printUnknownOrigins = false

func ParseFlags() {
	flag.StringVar(&port, PortFlag, port, "Port to listen on")
	flag.Var(&allowedOrigins, AllowedOriginFlag, "Add an allowed origin (can be specified multiple times)")
	flag.BoolVar(&printUnknownOrigins, "print-unknown-origins", printUnknownOrigins, "Print all unknown origin addresses")

	flag.Parse()
}

func GetPort() string {
	return port
}

func GetAllowedOrigins() map[string]struct{} {
	allowedOriginsSet := make(map[string]struct{}, len(allowedOrigins))

	for _, origin := range allowedOrigins {
		allowedOriginsSet[origin] = struct{}{}
	}

	return allowedOriginsSet
}

func GetAllowedOriginsAsString() string {
	return strings.Join(allowedOrigins, ", ")
}

func PrintUnknownOrigins() bool {
	return printUnknownOrigins
}
