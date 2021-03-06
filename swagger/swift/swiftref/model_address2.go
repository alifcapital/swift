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

type Address2 struct {
	StreetAddress []string `json:"street_address,omitempty"`
	City string `json:"city,omitempty"`
	Region string `json:"region,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	CountryName string `json:"country_name"`
}
