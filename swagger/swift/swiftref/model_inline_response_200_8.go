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

type InlineResponse2008 struct {
	Status *Status `json:"status,omitempty"`
	// The corresponding BIC for a given LEI
	Bic string `json:"bic"`
}
