package core

import (
	"math/rand"
	"unicode"
	"strings"
	"github.com/duraki/decipiat/models"

	_ "github.com/sirupsen/logrus"
	"fmt"
	"net"
	"sync"
)


func countChar(word string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range []rune(word) {
		count[r]++
	}
	return count
}

func bitsquatting(domain string) []models.Domain {
	results := []models.Domain{}
	masks := []int32{1,2,4,8,16,32,64,128}

	for i, c := range domain {
		for m := range masks {
			b := rune(int(c)^m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				d := fmt.Sprintf("%s%c%s", domain[:i], b, domain[i+1:])
				results = append(results, models.Domain{  Name: d, Type: models.Bitsquatting})
			}
		}
	}

	return results
}

func vowels(domain string) []models.Domain {
	results := []models.Domain{}
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	runes := []rune(domain)

	for i := 0; i < len(runes); i++ {
		for _, v := range vowels {
			switch runes[i] {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				if runes[i] != v {
					d := fmt.Sprintf("%s%c%s", string(runes[:i]), v, string(runes[i+1:]))
					results = append(results, models.Domain{  Name: d, Type: models.Vowels})
				}
			default:
			}
		}
	}
	return results
}

func repetition(domain string) []models.Domain {
	results := []models.Domain{}
	count := make(map[string]int)
	for i, c := range domain {
		if unicode.IsLetter(c) {
			result := fmt.Sprintf("%s%c%c%s", domain[:i], domain[i], domain[i], domain[i+1:])
			// remove duplicates
			count[result]++
			if count[result] < 2 {
				results = append(results, models.Domain{  Name: result, Type: models.Repetition})
			}
		}
	}
	return results
}

func omission(domain string) []models.Domain {
	results := []models.Domain{}
	for i := range domain {
		d := fmt.Sprintf("%s%s", domain[:i], domain[i+1:])
		results = append(results, models.Domain{  Name: d, Type: models.Omission})
	}
	return results
}

func homograph(domain string) []models.Domain {
	glyphs := map[rune][]rune{
		'a': {'à', 'á', 'â', 'ã', 'ä', 'å', 'ɑ', 'а', 'ạ', 'ǎ', 'ă', 'ȧ', 'α', 'ａ'},
		'b': {'d', 'ʙ', 'Ь', 'ɓ', 'Б', 'ß', 'β', 'ᛒ', '\u1E05', '\u1E03', '\u1D6C'}, // 'lb', 'ib'
		'c': {'ϲ', 'с', 'ƈ', 'ċ', 'ć', 'ç', 'ｃ'},
		'd': {'b', 'ԁ', 'ժ', 'ɗ', 'đ'}, // 'cl', 'dl', 'di'
		'e': {'é', 'ê', 'ë', 'ē', 'ĕ', 'ě', 'ė', 'е', 'ẹ', 'ę', 'є', 'ϵ', 'ҽ'},
		'f': {'Ϝ', 'ƒ', 'Ғ'},
		'g': {'q', 'ɢ', 'ɡ', 'Ԍ', 'Ԍ', 'ġ', 'ğ', 'ց', 'ǵ', 'ģ'},
		'h': {'һ', 'հ', '\u13C2', 'н'}, // 'lh', 'ih'
		'i': {'1', 'l', '\u13A5', 'í', 'ï', 'ı', 'ɩ', 'ι', 'ꙇ', 'ǐ', 'ĭ'},
		'j': {'ј', 'ʝ', 'ϳ', 'ɉ'},
		'k': {'κ', 'κ'}, // 'lk', 'ik', 'lc'
		'l': {'1', 'i', 'ɫ', 'ł'},
		'm': {'n', 'ṃ', 'ᴍ', 'м', 'ɱ'}, // 'nn', 'rn', 'rr'
		'n': {'m', 'r', 'ń'},
		'o': {'0', 'Ο', 'ο', 'О', 'о', 'Օ', 'ȯ', 'ọ', 'ỏ', 'ơ', 'ó', 'ö', 'ӧ', 'ｏ'},
		'p': {'ρ', 'р', 'ƿ', 'Ϸ', 'Þ'},
		'q': {'g', 'զ', 'ԛ', 'գ', 'ʠ'},
		'r': {'ʀ', 'Г', 'ᴦ', 'ɼ', 'ɽ'},
		's': {'Ⴝ', '\u13DA', 'ʂ', 'ś', 'ѕ'},
		't': {'τ', 'т', 'ţ'},
		'u': {'μ', 'υ', 'Ս', 'ս', 'ц', 'ᴜ', 'ǔ', 'ŭ'},
		'v': {'ѵ', 'ν', '\u1E7F', '\u1E7D'},      // 'v̇'
		'w': {'ѡ', 'ա', 'ԝ'}, // 'vv'
		'x': {'х', 'ҳ', '\u1E8B'},
		'y': {'ʏ', 'γ', 'у', 'Ү', 'ý'},
		'z': {'ʐ', 'ż', 'ź', 'ʐ', 'ᴢ'},
	}
	doneCount := make(map[rune]bool)
	results := []models.Domain{}
	runes := []rune(domain)
	count := countChar(domain)

	for i, char := range runes {
		// perform attack against single character
		for _, glyph := range glyphs[char] {
			d := fmt.Sprintf("%s%c%s", string(runes[:i]), glyph, string(runes[i+1:]))
			results = append(results, models.Domain{  Name: d, Type:models.Homograph})
		}
		// determine if character is a duplicate
		// and if the attack has already been performed
		// against all characters at the same time
		if count[char] > 1 && doneCount[char] != true {
			doneCount[char] = true
			for _, glyph := range glyphs[char] {
				result := strings.Replace(domain, string(char), string(glyph), -1)
				results = append(results, models.Domain{  Name: result, Type:models.Homograph})
			}
		}
	}
	return results
}

func hyphenation(domain string) []models.Domain {
	var results []models.Domain
	for i := 1; i < len(domain); i++ {
		if (rune(domain[i]) != '-' || rune(domain[i]) != '.') && (rune(domain[i-1]) != '-' || rune(domain[i-1]) != '.') {
			d := fmt.Sprintf("%s-%s", domain[:i], domain[i:])
			results = append(results, models.Domain{  Name: d, Type: models.Hyphenation})
		}
	}
	return results
}

func isAvailable(domain *models.Domain, wg *sync.WaitGroup) {
	defer wg.Done()
	ip_records, err := net.LookupIP(domain.Name)
	if err != nil {
		domain.Available = []string{"Unavailable"}
	} else {
		for _, rec := range ip_records {
			domain.Available = append(domain.Available, rec.String())
		}
	}
}

func GenerateSimilar(domain string, n int, types string) *models.AllDomains {
	var res [][]models.Domain
	var collected []models.Domain

	if types == "all" {
		res = append(res, bitsquatting(domain))
		res = append(res, vowels(domain))
		res = append(res, repetition(domain))
		res = append(res, omission(domain))
		res = append(res, homograph(domain))
		res = append(res, hyphenation(domain))
	} else {
		typesToCollect := strings.Split(types, "+")

		for _, v := range typesToCollect {
			switch v {
			case "bitsquatting":
				res = append(res, bitsquatting(domain))
			case "vowels":
				res = append(res, vowels(domain))
			case "repetition":
				res = append(res, repetition(domain))
			case "omission":
				res = append(res, omission(domain))
			case "homograph":
				res = append(res, homograph(domain))
			case "hyphenation":
				res = append(res, hyphenation(domain))
			}
		}
	}

	totalItems := 0

	for i := range res {
		totalItems += len(res[i])
	}

	if n > totalItems || n == 0 {
		n = totalItems
	}

	// Pick n elements from the results
	for i := 0; i < n; i++ {
		randomArrayNumber := rand.Intn(len(res))
		randomDomainNumber := rand.Intn(len(res[randomArrayNumber]))
		collected = append(collected, res[randomArrayNumber][randomDomainNumber])
	}

	// Check if domain is available
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go isAvailable(&collected[i], &wg)
	}

	wg.Wait()
	

	return &models.AllDomains{Domains: collected}
}