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
// AccountValidationSource1Code : * `ACSV`-The responder is the financial institution servicing this account, basing its decision on its account holding system. *`HIST`-The responder is SWIFT, basing its decision on SWIFT past transations (successful or failed) made to this account. *`OBSP`-The responder is a third-party system to which the account servicer delegated implementation of the API, basing its decision on data fed by the account servicer. The data feed frequency can vary from one implementation to another.
type AccountValidationSource1Code string

// List of AccountValidationSource1Code
const (
	ACSV_AccountValidationSource1Code AccountValidationSource1Code = "ACSV"
	HIST_AccountValidationSource1Code AccountValidationSource1Code = "HIST"
	OBSP_AccountValidationSource1Code AccountValidationSource1Code = "OBSP"
)
