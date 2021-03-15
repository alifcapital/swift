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
// GeoRoute1Code : > Specifies information about the geographic route the message is following. *`DMST_domestic`-Sender and receiver of the information are in the same country. *`INTL_international`-Sender and Receiver in different countries or the location of at least one of them is unknown. *`SWFT_swift`-Message involving SWIFT.
type GeoRoute1Code string

// List of GeoRoute1Code
const (
	DMST_DOMESTIC_GeoRoute1Code GeoRoute1Code = "DMST_domestic"
	INTL_INTERNATIONAL_GeoRoute1Code GeoRoute1Code = "INTL_international"
	SWFT_SWIFT_GeoRoute1Code GeoRoute1Code = "SWFT_swift"
)
