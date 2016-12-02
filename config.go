package main

import (
	"flag"
	"os"
	"strconv"
)

// Config : object to configure server
type Config struct {
	Port    int64
	FileDir string
}

const defaultPort = 4000
const defaultFileDir = "vendor/github.com/jpopesculian/wa-client/dist"

func loadConfig() Config {
	return Config{
		Port:    loadPortConfig(),
		FileDir: loadFileDirConfig(),
	}
}

func loadPortConfig() int64 {
	defaultPort := parsePortString(os.Getenv("PORT"), defaultPort)
	return *flag.Int64("port",
		defaultPort,
		"port to serve from")
}

func loadFileDirConfig() string {
	return *flag.String("files",
		defaultFileDir,
		"directory to serve static files from")
}

func parsePortString(str string, defaultVal int64) int64 {
	val, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		val = defaultVal
	}
	return val
}
