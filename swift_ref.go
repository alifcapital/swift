// SWIFTRef API (1.6.0)
package swift

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Address - address information
type Address struct {
	AddressLines       []string `json:"address_lines"`
	PostOfficeBox      string   `json:"post_office_box"`
	TownName           string   `json:"town_name"`
	CountrySubdivision string   `json:"country_subdivision"`
	PostCode           string   `json:"post_code"`
	// The country name in English of the institution/branch as indicated in the ISO 3166 list
	CountryName string `json:"country_name"`
	// The ISO 3166-1 alpha-2 code of the country of the institution/branch
	CountryCode string `json:"country_code"`
}

// ContactDetails - contact info
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
func GetDetailsOfBic(bic string, ctx context.Context, xApiKey string, e env) (*DetailsOfBic, error) {
	// create a new http-request instance
	req, err := http.NewRequest(http.MethodGet, getBikDetailsUrl(e, bic), nil)
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
		return nil, NewSWIFTError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response and return details
	details := DetailsOfBic{}
	if err := json.Unmarshal(b.Bytes(), &details); err != nil {
		return nil, err
	}
	return &details, nil
}
