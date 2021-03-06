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

type CountryCurrencies struct {
	// The ISO currency name
	Name string `json:"name"`
	// The ISO 3-letters currency code
	Iso3aCode string `json:"iso_3a_code"`
	// The ISO 3-digits currency code
	Iso3nCode string `json:"iso_3n_code,omitempty"`
}
