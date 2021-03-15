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
// PaymentPurposeCodeValidationStatus1Code : The list of unique status identifiers for the Payment Purpose Code Validation service. *`FVAL_valid`-The payment information instructed for validation it is valid in the instructed payment context. *`UNRE_unrecognised`-The payment information instructed for validation has not been found in the SWIFT repository for the instructed payment context. *`CNPR_could_not_perform`-The validation of the instructed payment information could not be performed in the instructed payment context typically due to missing payment context information. *`NAVL_not_available`-The validation of the instructed payment information is not available in the instructed payment context, typically due to payment data collection process being in progress and insufficient data coverage.
type PaymentPurposeCodeValidationStatus1Code string

// List of PaymentPurposeCodeValidationStatus1Code
const (
	FVAL_VALID_PaymentPurposeCodeValidationStatus1Code PaymentPurposeCodeValidationStatus1Code = "FVAL_valid"
	UNRE_UNRECOGNISED_PaymentPurposeCodeValidationStatus1Code PaymentPurposeCodeValidationStatus1Code = "UNRE_unrecognised"
	CNPR_COULD_NOT_PERFORM_PaymentPurposeCodeValidationStatus1Code PaymentPurposeCodeValidationStatus1Code = "CNPR_could_not_perform"
	NAVL_NOT_AVAILABLE_PaymentPurposeCodeValidationStatus1Code PaymentPurposeCodeValidationStatus1Code = "NAVL_not_available"
)
