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

type InlineResponse2001 struct {
	Status *Status `json:"status,omitempty"`
	// The BIC associated with the IBAN
	Bic string `json:"bic"`
}
