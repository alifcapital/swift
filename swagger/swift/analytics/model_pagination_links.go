/*
 * SWIFT Banking Analytics
 *
 * SWIFT Watch Banking Analytics API enables institutions to retrieve their  own SWIFT traffic data and the SWIFT totals, extending to the level of value and currency per market.   **v1.0.11 Release Notes**   * The Test Environment url updated to https://api-pilot.swift.com/bi/banking-analytics/v1  * PaginationLinks Object is defined as optional, as well as property next.   * Response examples were udpated to align with the API contract.   * Example for Market parameter in the request fixed
 *
 * API version: 1.0.11
 * Contact: developer-support@swift.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package analytics

// Links to help to find the self, next, previous, first and last pages of the response
type PaginationLinks struct {
	// The URI to the current page in the response
	Self string `json:"self"`
	// The URI to the next page of response data
	Next string `json:"next,omitempty"`
	// The URI to the previous page of response data
	Previous string `json:"previous,omitempty"`
	// The URI to the first page in the response
	First string `json:"first,omitempty"`
	// The URI to the last page in the response
	Last string `json:"last,omitempty"`
}
