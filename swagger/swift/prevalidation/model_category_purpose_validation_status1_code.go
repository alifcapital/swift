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
// CategoryPurposeValidationStatus1Code : The list of unique identifiers which describe the status of the Category Purpose validation. *`FVAL_valid`-The payment information instructed for validation it is valid in the instructed payment context. *`UNRE_unrecognised`-The payment information instructed for validation has not been found in the SWIFT repository for the instructed payment context. *`NAVL_not_available`-The validation of the instructed payment information is not available in the instructed payment context, typically due to payment data collection process being in progress and insufficient data coverage. *`CNPR_could_not_perform`-The validation of the instructed payment information could not be performed in the instructed payment context typically due to missing payment context information.
type CategoryPurposeValidationStatus1Code string

// List of CategoryPurposeValidationStatus1Code
const (
	FVAL_VALID_CategoryPurposeValidationStatus1Code CategoryPurposeValidationStatus1Code = "FVAL_valid"
	UNRE_UNRECOGNISED_CategoryPurposeValidationStatus1Code CategoryPurposeValidationStatus1Code = "UNRE_unrecognised"
	NAVL_NOT_AVAILABLE_CategoryPurposeValidationStatus1Code CategoryPurposeValidationStatus1Code = "NAVL_not_available"
	CNPR_COULD_NOT_PERFORM_CategoryPurposeValidationStatus1Code CategoryPurposeValidationStatus1Code = "CNPR_could_not_perform"
)
