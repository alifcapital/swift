package swift

import "context"

// Status - part of response schema of SWIFT-ref API
type Status struct {
	Http             string `json:"http"`
	Code             string `json:"code"`
	UserMessage      string `json:"user_message"`
	DeveloperMessage string `json:"developer_message"`
	MoreInfo         string `json:"more_info"`
}

// IBANResponseSchema - response schema after on success response
// on validating IBAN
type IBANResponseSchema struct {
	IBAN     string `json:"iban"`
	Validity string `json:"validity"`
	Status   Status `json:"status"`
}

// CheckValidityOfIBAN - check whether an IBAN is valid, that is its country code, structure, length, and checksum are valid.
// https: //developer.swift.com/reference#operation/getIbanValidity
func CheckValidityOfIBAN(ctx context.Context, iban string, e env) (*IBANResponseSchema, error) {
	return nil, nil
}
