/*
 * SWIFTRef API
 *
 * SWIFTRef API
 *
 * API version: 1.6.0
 * Contact: developer-support@swift.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swiftref

type Ssi struct {
	// The owner of the SSI Nostro Account
	OwnerBic string `json:"owner_bic"`
	// Name by which an institution is known and which is usually used to identify that institution
	InstitutionName string `json:"institution_name,omitempty"`
	// A free text description of the branch as provided by the financial institution to which it belongs; for the time being this will be provided only for entries with a BIC and only when the financial institution concerned wants to provide this extra information; the information is sourced from the BIC Directory
	BranchInformation string `json:"branch_information,omitempty"`
	Address *Address         `json:"address,omitempty"`
	// The ISO 4217 currency code of the requested SSIs
	CurrencyCode string `json:"currency_code"`
	// Indicates whether there is a direct account relationship between the owner of the SSI and the correspondent; if false, then this means at least the first intermediary must be present
	Direct bool                         `json:"direct"`
	Correspondent *CorrespondentBic     `json:"correspondent"`
	FirstIntermediary *IntermediaryBic  `json:"first_intermediary,omitempty"`
	SecondIntermediary *IntermediaryBic `json:"second_intermediary,omitempty"`
}