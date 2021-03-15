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
// AccountValidationCheckContext1Code : *`BENR`-Request is in the scope of a beneficiary registration validation. *`PAYM`-Request is in the scope of the payment initiation. *`RFPP`-Request is in the scope of a request for a payment initiation.
type AccountValidationCheckContext1Code string

// List of AccountValidationCheckContext1Code
const (
	BENR_AccountValidationCheckContext1Code AccountValidationCheckContext1Code = "BENR"
	PAYM_AccountValidationCheckContext1Code AccountValidationCheckContext1Code = "PAYM"
	RFPP_AccountValidationCheckContext1Code AccountValidationCheckContext1Code = "RFPP"
)