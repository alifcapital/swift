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
// Region1Code : > Specifies the geographical region, in coded form. *`AFRC_africa`-African region. *`APAC_asia_pacific`-Asia-Pacific region. *`CLAM_central_and_latin_america`-Central and Latin-American region. *`EURO_euro_zone`-Region of European countries that use the Euro currency. *`MDEA_middle_east`-Middle eastern region. *`NEUR_non_euro_zone`-Region of European countries that don't use the Euro currency. *`NOAM_north_america`-North American region.
type Region1Code string

// List of Region1Code
const (
	AFRC_AFRICA_Region1Code Region1Code = "AFRC_africa"
	APAC_ASIA_PACIFIC_Region1Code Region1Code = "APAC_asia_pacific"
	CLAM_CENTRAL_AND_LATIN_AMERICA_Region1Code Region1Code = "CLAM_central_and_latin_america"
	EURO_EURO_ZONE_Region1Code Region1Code = "EURO_euro_zone"
	MDEA_MIDDLE_EAST_Region1Code Region1Code = "MDEA_middle_east"
	NEUR_NON_EURO_ZONE_Region1Code Region1Code = "NEUR_non_euro_zone"
	NOAM_NORTH_AMERICA_Region1Code Region1Code = "NOAM_north_america"
)