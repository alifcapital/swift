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
// CustomerRoute1Code : > Specifies the customer route of the message. *`SWFT_swift`-Traffic exchanged with SWIFT. *`ITER_inter_customer_group`-Traffic exchanged between different institutions. *`ITRA_intra_customer_group`-Traffic exchanged within the same institution group of branches.
type CustomerRoute1Code string

// List of CustomerRoute1Code
const (
	SWFT_SWIFT_CustomerRoute1Code CustomerRoute1Code = "SWFT_swift"
	ITER_INTER_CUSTOMER_GROUP_CustomerRoute1Code CustomerRoute1Code = "ITER_inter_customer_group"
	ITRA_INTRA_CUSTOMER_GROUP_CustomerRoute1Code CustomerRoute1Code = "ITRA_intra_customer_group"
)