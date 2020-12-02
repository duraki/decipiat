package core

import (
	"log"
	"os"
	"os/exec"
)

/*
GenerateCertificate will generate certificate based on ProxyDomain field from type ModlishkaConfig struct from models/modlishka.go.
Currently, it requires to have certbot binary installed.
*/
func GenerateCertificate(domain string) (string, string, error) {
	certbot, err := exec.LookPath("certbot")
	if err != nil {
		return "", "", err
	}
	log.Printf("Certbot found at %s\n", certbot)
	cmd := exec.Cmd{Path: certbot,
		Args: []string{"sudo", "certonly", "--manual", "--preferred-challenges=dns", "--server", "https://acme-v02.api.letsencrypt.org/directory", "--agree-tos", "-d", "*." + domain, "--email", "noreply@live.com"},
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	return "", "", nil
}
