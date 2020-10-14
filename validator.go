package swift_sdk

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidAnyBic                   = errors.New("bic number is not valid")
	ErrCreditorAccountLengthOutOfRange = errors.New("credit account length out of 1 to 34 range")
	ErrCreditorNameLengthOutOfRange    = errors.New("credit name length out of 1 to 140 range")
	ErrInvalidCountryCode              = errors.New("county code must contain only 2 uppercase symbols")

	reAnyBic      = regexp.MustCompile(`^[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}$`)
	reCountryCode = regexp.MustCompile(`^[A-Z]{2,2}$`)
)

// ValidateBic - validates BIC format
// BIC - code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority
func ValidateBic(anyBic string) error {
	if !reAnyBic.MatchString(anyBic) {
		return ErrInvalidAnyBic
	}
	return nil
}

// ValidateBicFi - same as `ValidateBic`
func ValidateBicFi(bicFi string) error {
	return ValidateBic(bicFi)
}

// ValidateCreditorAccount - strings, between 1 and 34 chars
func ValidateCreditorAccount(creditorAccount string) error {
	creditorAccount = strings.TrimSpace(creditorAccount)
	length := utf8.RuneCountInString(creditorAccount)
	if 1 > length || length > 34 {
		return ErrCreditorAccountLengthOutOfRange
	}
	return nil
}

// ValidateCreditorName - strings, between 1 and 140 chars
func ValidateCreditorName(creditorAccount string) error {
	creditorAccount = strings.TrimSpace(creditorAccount)
	length := utf8.RuneCountInString(creditorAccount)
	if 1 > length || length > 140 {
		return ErrCreditorNameLengthOutOfRange
	}
	return nil
}

// ValidateCountryCode - ISO 3166, Alpha-2 code
func ValidateCountryCode(countryCode string) error {
	if !reCountryCode.MatchString(countryCode) {
		return ErrInvalidCountryCode
	}
	return nil
}
