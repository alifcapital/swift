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
// Market1Code : > Specifies the business market. *`PMTS_payments`-Specifies the payments market. *`TRFN_trade_finance`-Specifies the trade finance market. *`TREA_treasury`-Specifies the treasury market.
type Market1Code string

// List of Market1Code
const (
	PMTS_PAYMENTS_Market1Code Market1Code = "PMTS_payments"
	TRFN_TRADE_FINANCE_Market1Code Market1Code = "TRFN_trade_finance"
	TREA_TREASURY_Market1Code Market1Code = "TREA_treasury"
)
