package models

type DomainType string

const (
	Bitsquatting DomainType = "Bitsquatting"
	Vowels = "Vowels"
	Repetition = "Repetition"
	Omission = "Omission"
	Homograph = "Homograph"
	Hyphenation = "Hyphenation"
)

type Domain struct {
	Name 			string
	Type 			DomainType
	Available		[]string
}

type AllDomains struct {
	Domains []Domain
}
