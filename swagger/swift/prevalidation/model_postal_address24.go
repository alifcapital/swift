/*
 * Payment Pre-validation API
 *
 * New version of the Bank Account Verification Service (the BAV), BAV v2, for Users and Data Providers. Additionally to BAV v1, depending on the jurisdiction where the account is held and the market practices in use, BAV v2 brings in the validation of the account type and if the creditor name partially matches the account holder name.   **v2.1.1Release Notes**  * Addition of Beneficiary Account Verification   *  Beneficiary Account Verification has been updated to version 2, which includes features to better support domestic systems, future support of ISO20022 addresses and support for systems that may provide answers based on past transactions.  * Schema Organisation8 updated to OrganisationIdentification8
 *
 * API version: 2.1.1
 * Contact: developer-support@swift.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package prevalidation

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress24 struct {
	AddressType *AddressType3Choice `json:"address_type,omitempty"`
	Department string `json:"department,omitempty"`
	SubDepartment string `json:"sub_department,omitempty"`
	StreetName string `json:"street_name,omitempty"`
	BuildingNumber string `json:"building_number,omitempty"`
	BuildingName string `json:"building_name,omitempty"`
	Floor string `json:"floor,omitempty"`
	PostBox string `json:"post_box,omitempty"`
	Room string `json:"room,omitempty"`
	PostCode string `json:"post_code,omitempty"`
	TownName string `json:"town_name,omitempty"`
	TownLocationName string `json:"town_location_name,omitempty"`
	DistrictName string `json:"district_name,omitempty"`
	CountrySubDivision string `json:"country_sub_division,omitempty"`
	Country string `json:"country,omitempty"`
	AddressLine []string `json:"address_line,omitempty"`
}