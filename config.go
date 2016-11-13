package main

import (
	"flag"
	"os"
	"strconv"
)

func loadFlagsAndEnv(config map[string]interface{}) map[string]interface{} {
	defaultPort, err := strconv.ParseInt(os.Getenv("PORT"), 10, 16)
	if err != nil {
		defaultPort = 4000
	}
	config["port"] = *flag.Int64("port", defaultPort, "port to serve from")

	return config
}
