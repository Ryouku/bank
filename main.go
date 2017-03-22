package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/bvnk/bank/configuration"
)

const (
	// This is the FQDN from the certs generated
	CONN_HOST = "localhost"
	CONN_PORT = "3300"
	CONN_TYPE = "tcp"
	HTTP_PORT = "8443"
)

var logPath *string

func main() {
	argClientServer := flag.String("mode", "server", "Mode to run the service in")
	configPath := flag.String("configPath", "/etc/bvnk/config.json", "Config path absolute location. Default /etc/bvnk/config.json")
	logPath = flag.String("logPath", "/var/log/bvnk/bank.log", "Log path absolute location. Default /var/log/bvnk/bank.log")
	flag.Parse()

	configuration.SetConfigPath(*configPath)

	err := parseArguments(*argClientServer)
	if err != nil {
		log.Fatalf("Error starting, err: %v\n", err)
	}
	os.Exit(0)
}

func parseArguments(arg string) (err error) {

	switch arg {
	case "http":
		err := RunHttpServer()
		if err != nil {
			log.Fatalf("Could not start HTTP server. " + err.Error())
		}
		break
	case "client":
		// Run client for bank system
		runClient("tls")
		break
	case "clientNoTLS":
		// Run client for bank system
		runClient("no-tls")
		break
	case "server":
		// Run server for bank system
		for {
			runServer("tls")
		}
	case "serverNoTLS":
		// Run server for bank system
		for {
			runServer("no-tls")
		}
	default:
		return errors.New("No valid option chosen. Valid options: client, clientNoTLS, server, serverNoTLS")
	}

	return
}

// Simple log function for logging to a file
func bLog(logLevel int, message string, functionName string) (err error) {
	f, err := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	// Check logLevel
	if logLevel > 4 {
		// Default to highest available to avoid returning errors
		logLevel = 4
	}

	// Load app config
	Config, err := configuration.LoadConfig()
	if err != nil {
		return errors.New("main.bLog: " + err.Error())
	}
	// Check log level based on config
	// logLevel is an int: 0 debug, 1 info, 2 warning, 3 error, 4 critical
	// List of colours: https://radu.cotescu.com/coloured-log-outputs/
	// Default Blue
	colourBegin := "\033[0;34m"
	switch Config.LogLevel {
	case "critical":
		if logLevel < 4 {
			return
		}
		// High intensity red
		colourBegin = "\033[0;91m"
		break
	case "error":
		if logLevel < 3 {
			return
		}
		// Red
		colourBegin = "\033[0;31m"
		break
	case "warning":
		if logLevel < 2 {
			return
		}
		// Yellow
		colourBegin = "\033[0;33m"
		break
	case "info":
		if logLevel < 1 {
			return
		}
		// Cyan
		colourBegin = "\033[0;36m"
		break
	case "debug":
		// Log everything
		break
	}

	colourEnd := "\033[39m"
	// Construct message
	log.Printf("%s%s :: %s%s", colourBegin, message, functionName, colourEnd)
	return
}

func trace() (funcTrace string) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	// Go two levels deep
	funcTrace = ""
	for i := 0; i < 2; i++ {
		if pc[i] != 0 {
			f := runtime.FuncForPC(pc[i])
			file, line := f.FileLine(pc[i])
			funcTrace = funcTrace + fmt.Sprintf("%s:%d %s | ", file, line, f.Name())
		}
	}
	return
}
