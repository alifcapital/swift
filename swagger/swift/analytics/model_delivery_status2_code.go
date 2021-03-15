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
// DeliveryStatus2Code : > Specifies the status of the delivery of a message, in coded form. *`ABRT_aborted`-Sending of message was aborted by sender. *`DLVR_delivered`-Message was delivered to receiver. *`NACK_nacked`-Message was not acknowledged by receiver. *`RFUS_refused_by_infrastructure`-Message was refused by the infrastructure (e.g. SWIFT)
type DeliveryStatus2Code string

// List of DeliveryStatus2Code
const (
	ABRT_ABORTED_DeliveryStatus2Code DeliveryStatus2Code = "ABRT_aborted"
	DLVR_DELIVERED_DeliveryStatus2Code DeliveryStatus2Code = "DLVR_delivered"
	NACK_NACKED_DeliveryStatus2Code DeliveryStatus2Code = "NACK_nacked"
	RFUS_REFUSED_BY_INFRASTRUCTURE_DeliveryStatus2Code DeliveryStatus2Code = "RFUS_refused_by_infrastructure"
)
