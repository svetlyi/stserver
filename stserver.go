package main

import (
	"flag"
	"fmt"
	"github.com/svetlyi/stserver/server"
	"github.com/svetlyi/stserver/tools"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	rootDir string
	port    int
	prefix  string
	address string
)

const maxPort = 65535
const minOpenPort = 49152
const minPort = 1024

const defaultPrefix = "/"

func init() {
	flag.StringVar(&rootDir, "r", "", "root dir")
	flag.StringVar(&prefix, "pr", defaultPrefix, "prefix (localhost:8080/prefix)")
	flag.StringVar(&address, "a", "", "address")
	flag.IntVar(&port, "p", tools.GetRandomDynamicPort(maxPort, minOpenPort), "port (default - random)")
}

func main() {
	flag.Parse()
	logger := tools.CreateLogger()

	if rootDir == "" {
		logger.Fatal("root dir is necessary")
	}
	var err error
	if rootDir, err = filepath.Abs(filepath.Clean(rootDir)); err != nil {
		logger.Fatal("could not get absolute path for", rootDir)
	}

	if port < minPort || port > maxPort {
		logger.Fatalf("port must be between %d and %d", minPort, maxPort)
	}

	if prefix != defaultPrefix {
		prefix = fmt.Sprintf("/%s/", strings.Trim(prefix, "/"))
	}

	logger.Infof("starting serving folder %s on port %d; prefix: %s", rootDir, port, prefix)
	tools.ListAddresses(port, logger)

	http.Handle(prefix, http.StripPrefix(prefix, server.NewFileServerHandler(rootDir, logger)))
	if err = http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil); err != nil {
		log.Fatalln("could not start server", err)
	}
}
