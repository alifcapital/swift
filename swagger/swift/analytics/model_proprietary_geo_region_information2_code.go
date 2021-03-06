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
// ProprietaryGeoRegionInformation2Code : > Contains information about a country or geographical region in a proprietary, coded form. *`IMIR_imi_and_related_copies`-IMI and related copies. *`NOTA_not_available`-Not available. *`OTHR_other`-Other. *`SWFT_swift`-SWIFT
type ProprietaryGeoRegionInformation2Code string

// List of ProprietaryGeoRegionInformation2Code
const (
	IMIR_IMI_AND_RELATED_COPIES_ProprietaryGeoRegionInformation2Code ProprietaryGeoRegionInformation2Code = "IMIR_imi_and_related_copies"
	NOTA_NOT_AVAILABLE_ProprietaryGeoRegionInformation2Code ProprietaryGeoRegionInformation2Code = "NOTA_not_available"
	OTHR_OTHER_ProprietaryGeoRegionInformation2Code ProprietaryGeoRegionInformation2Code = "OTHR_other"
	SWFT_SWIFT_ProprietaryGeoRegionInformation2Code ProprietaryGeoRegionInformation2Code = "SWFT_swift"
)
