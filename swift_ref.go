// SWIFTRef API (1.6.0)
package swift

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Address - address information of financial institution by BIC number
type Address struct {
	AddressLines       []string `json:"address_lines"`
	PostOfficeBox      string   `json:"post_office_box"`
	TownName           string   `json:"town_name"`
	CountrySubdivision string   `json:"country_subdivision"`
	PostCode           string   `json:"post_code"`
	// CountryName - country name in English of the institution/branch as indicated in the ISO 3166 list
	CountryName string `json:"country_name"`
	// CountryCode - ISO 3166-1 alpha-2 code of the country of the institution/branch
	CountryCode string `json:"country_code"`
}

// ContactDetails - contact of financial institution by BIC number
type ContactDetails struct {
	PhoneNumber  string `json:"phone_number"`
	FaxNumber    string `json:"fax_number"`
	EmailAddress string `json:"email_address"`
	WebAddress   string `json:"web_address"`
}

// DetailsOfBic - returned in response to request of details with only one value, BIK number
// https://developer.swift.com/reference#tag/bics
type DetailsOfBic struct {
	// The BIC of the institution on which more information is requested
	Bic string `json:"bic"`
	// Name by which a party is known and which is usually used to identify that party
	InstitutionName string `json:"institution_name"`
	// A free text description of the branch as provided by the financial institution
	// to which it belongs; this is currently provided for entries with a BIC when
	// the financial institution concerned wants to provide this extra information
	BranchInformation string `json:"branch_information"`
	// Address of BIC owner
	Address Address `json:"address"`
	// Contact info of bic owner
	ContactDetails ContactDetails `json:"contact_details"`
	// Status of the entity in the office hierarchy
	OfficeType string `json:"office_type"`
}

// GetDetailsOfBic - implements https://developer.swift.com/reference#tag/bics specifications
// bic - target bic
// xApiKey - {{consumer-key}}
func GetDetailsOfBic(ctx context.Context, bic, xApiKey string, e env) (*DetailsOfBic, error) {
	// create a new http-request instance
	req, err := http.NewRequest(http.MethodGet, bicDetailsUrl(e, bic), nil)
	if err != nil {
		return nil, err
	}
	// set credential
	req.Header.Set("x-api-key", xApiKey)

	// set any given context
	req = req.WithContext(ctx)

	// make request to remote endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	var b bytes.Buffer
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle not 200 error codes
	if resp.StatusCode != http.StatusOK {
		return nil, NewHttpError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response and return details
	details := DetailsOfBic{}
	if err := json.Unmarshal(b.Bytes(), &details); err != nil {
		return nil, err
	}
	return &details, nil
}

// ------------------------------------------------------------------

// Status - part of response schema of SWIFT-ref API
type Status struct {
	Http             string `json:"http"`
	Code             string `json:"code"`
	UserMessage      string `json:"user_message"`
	DeveloperMessage string `json:"developer_message"`
	MoreInfo         string `json:"more_info"`
}

// IBANResponseSchema - response schema after on success response
type IBANResponseSchema struct {
	IBAN string `json:"iban"`
	// The code `IVAL` if the IBAN is valid
	Validity string `json:"validity"`
}

func (irs *IBANResponseSchema) Valid() bool {
	return irs.Validity == "IVAL"
}

// CheckValidityOfIBAN - check whether an IBAN is valid, that is its country code, structure, length, and checksum are valid.
// https: //developer.swift.com/reference#operation/getIbanValidity
func CheckValidityOfIBAN(ctx context.Context, iban, xApiKey string, e env) (*IBANResponseSchema, error) {
	// create a new http-request instance
	req, err := http.NewRequest(http.MethodGet, validIBANUrl(e, iban), nil)
	if err != nil {
		return nil, err
	}
	// set credential
	req.Header.Set("x-api-key", xApiKey)

	// make request to remote endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	var b bytes.Buffer
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle not 200 error codes
	if resp.StatusCode != http.StatusOK {
		return nil, NewHttpError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response and return responseSchema
	responseSchema := IBANResponseSchema{}
	if err := json.Unmarshal(b.Bytes(), &responseSchema); err != nil {
		return nil, err
	}
	return &responseSchema, nil
}

// ------------------------------------------------------------------

// IBANComponents - components of IBAN
type IBANComponents struct {
	IBAN          string `json:"iban"`
	CountryCode   string `json:"country_code"`
	Checksum      string `json:"checksum"`
	BankId        string `json:"bank_id"`
	BranchId      string `json:"branch_id"`
	AccountNumber string `json:"account_number"`
	Length        int    `json:"length"`
}

// GetDetailsForIBAN - For a given IBAN, you can obtain the components of the IBAN.
// https://developer.swift.com/reference#operation/getIbanDetails_v2
func GetDetailsForIBAN(ctx context.Context, iban, xApiKey string, e env) (*IBANComponents, error) {
	// create a new http-request instance
	req, err := http.NewRequest(http.MethodGet, detailsOfIBANUrl(e, iban), nil)
	if err != nil {
		return nil, err
	}
	// set credential
	req.Header.Set("x-api-key", xApiKey)

	// make request to remote endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	var b bytes.Buffer
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle not 200 error codes
	if resp.StatusCode != http.StatusOK {
		return nil, NewHttpError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response and return components
	components := IBANComponents{}
	if err := json.Unmarshal(b.Bytes(), &components); err != nil {
		return nil, err
	}
	return &components, nil
}

// ------------------------------------------------------------------

type bicByIBAN struct {
	Bic string `json:"bic"`
}

func GetBicByIBAN(ctx context.Context, iban, xApiKey string, e env) (string, error) {
	// create a new http-request instance
	req, err := http.NewRequest(http.MethodGet, bicByIBANUrl(e, iban), nil)
	if err != nil {
		return "", err
	}
	// set credential
	req.Header.Set("x-api-key", xApiKey)

	// make request to remote endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read response body
	var b bytes.Buffer
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	// handle not 200 error codes
	if resp.StatusCode != http.StatusOK {
		return "", NewHttpError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response and return bic
	bic := bicByIBAN{}
	if err := json.Unmarshal(b.Bytes(), &bic); err != nil {
		return "", err
	}
	return bic.Bic, nil
}
