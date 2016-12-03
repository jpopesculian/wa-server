package config

import (
	"flag"
	"os"
	"strconv"
)

var (
	Port            *int64
	StaticFileDir   *string
	UserServiceAddr *string
)

var (
	defaultPort            int64 = 4000
	defaultStaticFileDir         = "vendor/github.com/jpopesculian/wa-client/dist"
	defaultUserServiceAddr       = "http://127.0.0.1:4000"
)

func init() {
	var err error

	if newDefaultPort, present := os.LookupEnv("PORT"); present {
		if defaultPort, err = strconv.ParseInt(newDefaultPort, 10, 16); err != nil {
			panic(err)
		}
	}

	Port = flag.Int64("port", defaultPort, "Port to serve from")
	StaticFileDir = flag.String("static-files", defaultStaticFileDir, "Static file directory")
	UserServiceAddr = flag.String("user-service", defaultUserServiceAddr, "Address of User Service")
}

func Parse() {
	flag.Parse()
}
