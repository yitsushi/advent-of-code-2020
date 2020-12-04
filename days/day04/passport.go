package day04

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
)

type passport struct {
	ID             string
	CountryID      int64
	BirthYear      int64
	IssueYear      int64
	ExpirationYear int64
	Height         struct {
		Value int64
		Unit  string
	}
	HairColor string
	EyeColor  string

	rawFields map[string]string
}

func parsePassport(blob []string) passport {
	p := passport{
		rawFields: map[string]string{},
	}

	for _, data := range blob {
		parts := strings.SplitN(data, ":", 2)
		key, value := parts[0], parts[1]

		p.ParseField(key, value)
	}

	return p
}

// ParseField parses and populates a field based on a given key/value pair.
func (p *passport) ParseField(key, value string) {
	p.rawFields[key] = value

	switch key {
	case "byr":
		p.BirthYear, _ = strconv.ParseInt(value, 10, 64)
	case "iyr":
		p.IssueYear, _ = strconv.ParseInt(value, 10, 64)
	case "eyr":
		p.ExpirationYear, _ = strconv.ParseInt(value, 10, 64)
	case "hgt":
		result := regexp.MustCompile(`(\d+)(.*)`).FindStringSubmatch(value)
		p.Height.Value = bullshit.DropErrorInt64(strconv.ParseInt(result[1], 10, 64))
		p.Height.Unit = result[2]
	case "hcl":
		p.HairColor = value
	case "ecl":
		p.EyeColor = value
	case "pid":
		p.ID = value
	case "cid":
		p.CountryID, _ = strconv.ParseInt(value, 10, 64)
	}
}

// Validate check only if fields are defined.
// Ignore CountryID for Part 1.
func (p *passport) ValidateFake() bool {
	requiredFields := []string{
		"byr", "iyr", "eyr", "hgt",
		"hcl", "ecl", "pid",
	}

	for _, key := range requiredFields {
		if _, ok := p.rawFields[key]; !ok {
			return false
		}
	}

	return true
}

// Validate a passport.
func (p *passport) Validate() bool {
	return (p.validateBirthYear() &&
		p.validateIssueYear() &&
		p.validateExpirationYear() &&
		p.validateHeight() &&
		p.validateHairColor() &&
		p.validateEyeColor() &&
		p.validateID() &&
		p.validateCountryID())
}

func (p passport) validateBirthYear() bool {
	return p.BirthYear >= 1920 && p.BirthYear <= 2002
}

func (p passport) validateIssueYear() bool {
	return p.IssueYear >= 2010 && p.IssueYear <= 2020
}

func (p passport) validateExpirationYear() bool {
	return p.ExpirationYear >= 2020 && p.ExpirationYear <= 2030
}

func (p passport) validateHeight() bool {
	if p.Height.Unit == "cm" {
		return p.Height.Value >= 150 && p.Height.Value <= 193
	}

	if p.Height.Unit == "in" {
		return p.Height.Value >= 59 && p.Height.Value < 76
	}

	return false
}

func (p passport) validateHairColor() bool {
	return bullshit.DropErrorBoolean(regexp.MatchString(`^#[a-z0-9]{6}$`, p.HairColor))
}

func (p passport) validateEyeColor() bool {
	validColors := []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}

	for _, color := range validColors {
		if p.EyeColor == color {
			return true
		}
	}

	return false
}

func (p passport) validateID() bool {
	return bullshit.DropErrorBoolean(regexp.MatchString(`^[0-9]{9}$`, p.ID))
}

func (p passport) validateCountryID() bool {
	return true
}
