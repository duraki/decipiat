package cli

import (
	"flag"
	"fmt"
)

type Config struct {
	Host    *string
	Port    *int
	Ssl     *bool
	Version *bool
}

var (
	options = Config{
		Host:    flag.String("host", "127.0.0.1", "Set hostname"),
		Port:    flag.Int("port", 8080, "Set port"),
		Ssl:     flag.Bool("ssl", true, "Append SSL/TLS config"),
		Version: flag.Bool("version", false, "Display version and author information and exit."),
	}
)

func usage() {
	fmt.Printf("Usage: ./decipiat up [OPTIONS]\n\n")
	fmt.Printf("\tOffensive Phishing-as-a-Service environment.\n")
	fmt.Printf("\tDeveloped by:\tErhad Husovic -> (Github: XdaemonX)\n")
	fmt.Printf("\t\t\tHalis Duraki  -> (Github: duraki)\n\n")
	flag.PrintDefaults()
}

func ParseConfiguration() *Config {
	flag.Usage = usage
	flag.Parse()

	return &options
}