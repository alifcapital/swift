/*
 * Payment Pre-validation API
 *
 * New version of the Bank Account Verification Service (the BAV), BAV v2, for Users and Data Providers. Additionally to BAV v1, depending on the jurisdiction where the account is held and the market practices in use, BAV v2 brings in the validation of the account type and if the creditor name partially matches the account holder name.   **v2.1.1Release Notes**  * Addition of Beneficiary Account Verification   *  Beneficiary Account Verification has been updated to version 2, which includes features to better support domestic systems, future support of ISO20022 addresses and support for systems that may provide answers based on past transactions.  * Schema Organisation8 updated to OrganisationIdentification8
 *
 * API version: 2.1.1
 * Contact: developer-support@swift.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package prevalidation
// AccountValidationResponse4Code : *`MTCH`-Matched validation check. *`MBAM`-The beneficiary bank has determined that the provided creditor name closely resembles the account holder name. However, it is not an exact match. The beneficiary bank can choose to disclose the actual account holder name for the payment sender to verify or update its records. *`NMTC`-Unmatched validation check. *`NOAP`-Validation check is not applicable. *`NOTC`-Validation check has not been carried out.
type AccountValidationResponse4Code string

// List of AccountValidationResponse4Code
const (
	MTCH_AccountValidationResponse4Code AccountValidationResponse4Code = "MTCH"
	MBAM_AccountValidationResponse4Code AccountValidationResponse4Code = "MBAM"
	NMTC_AccountValidationResponse4Code AccountValidationResponse4Code = "NMTC"
	NOAP_AccountValidationResponse4Code AccountValidationResponse4Code = "NOAP"
	NOTC_AccountValidationResponse4Code AccountValidationResponse4Code = "NOTC"
)
