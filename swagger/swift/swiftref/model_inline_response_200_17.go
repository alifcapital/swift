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

type InlineResponse20017 struct {
	Status *Status `json:"status,omitempty"`
	// The currency code that was validated
	CurrencyCode string `json:"currency_code"`
	// The code VCUC if the currency code is valid
	Validity string `json:"validity"`
}
