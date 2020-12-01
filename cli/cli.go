package cli

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Config struct handles options that will be passed to the web server
type Config struct {
	Host     *string
	Port     *int
	Ssl      *bool
	CertPath *string
	KeyPath  *string
	Version  *bool
}

var (
	options = Config{
		Host:     flag.String("host", "127.0.0.1", "Set hostname"),
		Port:     flag.Int("port", 8080, "Set port"),
		Ssl:      flag.Bool("ssl", true, "Append SSL/TLS config"),
		Version:  flag.Bool("version", false, "Display version and author information and exit."),
		CertPath: flag.String("certPath", "", "Path to certificate file"),
		KeyPath:  flag.String("keyPath", "", "Path to key file"),
	}
)

// Usage prints details about decipiat; should be unexported really
func Usage() {
	fmt.Printf("Usage: ./decipiat up [OPTIONS]\n\n")
	fmt.Printf("\tOffensive Phishing-as-a-Service environment.\n")
	fmt.Printf("\tDeveloped by:\n")
	fmt.Printf("\t\t\tErhad Husovic -> (Github: XdaemonX)\n")
	fmt.Printf("\t\t\tHalis Duraki  -> (Github: duraki)\n\n")
	flag.PrintDefaults()

	fmt.Printf("\n\n")
}

// ParseConfiguration is the main function that will be called to parse flags
func ParseConfiguration() *Config {
	flag.Usage = Usage
	flag.Parse()

	return &options
}

func InitializeLogging() {
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = "02.01.2006 15:05:04"
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
}
