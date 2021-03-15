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

type InlineResponse20019 struct {
	Status *Status `json:"status,omitempty"`
	// The LEI that was validated
	Lei string `json:"lei"`
	// The code VLEI if the LEI is valid
	Validity string `json:"validity"`
}
