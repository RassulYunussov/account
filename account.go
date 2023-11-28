package account

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"
)

var ErrInvalidAccountNumber = errors.New("Invalid account number")
var ErrIdOutOfRange = errors.New("Id out of range")
var ErrInvalidArgument = errors.New("Invalid input arguments")
var ErrCouldNotGenerate = errors.New("Cound not generate account number out of id")
var ErrCouldNotParseBigInteger = errors.New("Cound not parse big integer")

const (
	min_id              = int64(1)
	max_id              = int64(999999999999)
	country_length      = 2
	organization_length = 2
	magic_number        = int64(98)
	magic_modulator     = int64(97)
	blank_controlsum    = "00"
	account_suffix      = "AC"
	account_length      = 20
)

var modulator *big.Int

func init() {
	modulator = big.NewInt(magic_modulator)
}

// limitation up to 999 999 999 999 accounts
func CreateAccountNumber(country string, organization string, id int64) (*string, error) {
	if id < min_id || id > max_id {
		return nil, ErrIdOutOfRange
	}
	if len(country) != country_length || len(organization) != organization_length {
		return nil, ErrInvalidArgument
	}
	country = strings.ToUpper(country)
	organization = strings.ToUpper(organization)
	stringRepresentation := fmt.Sprintf("%012d", id)
	bigNumber := constructBigNumber(country, organization, stringRepresentation)
	controlSum, err := getControlSum(bigNumber)
	if err != nil {
		return nil, ErrCouldNotGenerate
	}
	result := fmt.Sprintf("%s%02d%s%s%s", country, *controlSum, organization, account_suffix, stringRepresentation)
	return &result, nil
}

func IsValid(account string) bool {
	if len(account) != account_length {
		return false
	}
	countryPrefix := account[:2]
	var countryPrefixConvertedToDigits string
	for _, ch := range countryPrefix {
		countryPrefixConvertedToDigits += convertCharToDecimal(ch)
	}
	controlSum := account[2:4]
	right := account[4:]
	var rightConvertedToDigits string
	for _, ch := range right {
		if unicode.IsDigit(ch) {
			rightConvertedToDigits += string(ch)
		} else {
			rightConvertedToDigits += convertCharToDecimal(ch)
		}
	}
	return checkControlSum(fmt.Sprintf("%s%s%s", rightConvertedToDigits, countryPrefixConvertedToDigits, controlSum))
}

func getBigIntFromString(bigNumberString string) (*big.Int, error) {
	bigNumber := big.Int{}
	if _, success := bigNumber.SetString(bigNumberString, 10); success {
		return &bigNumber, nil
	}
	return nil, ErrCouldNotParseBigInteger
}

func convertCharToDecimal(ch rune) string {
	ascii := int(ch) - 55
	return fmt.Sprintf("%d", ascii)
}

func getControlSum(bigNumberString string) (*int64, error) {
	bigInt, err := getBigIntFromString(bigNumberString)
	if err != nil {
		return nil, err
	}
	modulated := bigInt.Mod(bigInt, modulator)
	controlSum := magic_number - modulated.Int64()
	return &controlSum, nil
}

func checkControlSum(bigNumberString string) bool {
	bigInt, err := getBigIntFromString(bigNumberString)
	if err != nil {
		return false
	}
	modulated := bigInt.Mod(bigInt, modulator)
	return modulated.Int64() == 1
}

func constructBigNumber(country string, organization string, idString string) string {
	var prefix string
	for _, ch := range fmt.Sprintf("%s%s", organization, account_suffix) {
		prefix += convertCharToDecimal(ch)
	}
	output := idString
	for _, ch := range country {
		output += convertCharToDecimal(ch)
	}
	return fmt.Sprintf("%s%s%s", prefix, output, blank_controlsum)
}
