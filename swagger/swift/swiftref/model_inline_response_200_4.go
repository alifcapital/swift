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

type InlineResponse2004 struct {
	Status *Status `json:"status,omitempty"`
	// The corresponding IBAN
	Iban string `json:"iban"`
}