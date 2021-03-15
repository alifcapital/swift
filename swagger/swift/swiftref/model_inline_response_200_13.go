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

type InlineResponse20013 struct {
	Status *Status `json:"status,omitempty"`
	// The National ID that was validated
	NationalId string `json:"national_id"`
	// The scheme under which the National ID is defined
	Scheme string `json:"scheme"`
	// Contains the code VNID (Valid National ID) if the National ID is valid in the provided scheme or country
	Validity string `json:"validity"`
}
