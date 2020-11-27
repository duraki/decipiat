package core

import (
	"bytes"
	"os/exec"
	"strings"
)

// GetDomainPermutations will return permutations of certain domain generated by dnstwist
func GetDomainPermutations(word string) []string {
	var out bytes.Buffer

	cmd := exec.Command("dnstwist", "--format", "list", word)
	cmd.Stdout = &out
	_ = cmd.Run()

	words := strings.Split(out.String(),"\n")

	return words
}
