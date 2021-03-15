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

type NationalIdDetails struct {
	// National ID for which details were requested
	Id string `json:"id"`
	// The scheme under which the National ID is defined
	Scheme string `json:"scheme"`
	// Name by which a party is known and which is usually used to identify that party
	InstitutionName string `json:"institution_name"`
	// A free text description of the branch as provided by the financial institution to which it belongs; this is currently provided for entries with a BIC when the financial institution concerned wants to provide this extra information
	BranchInformation string       `json:"branch_information,omitempty"`
	Address *Address               `json:"address,omitempty"`
	ContactDetails *ContactDetails `json:"contact_details,omitempty"`
	// Status of the entity in the office hierarchy
	OfficeType string `json:"office_type,omitempty"`
}
