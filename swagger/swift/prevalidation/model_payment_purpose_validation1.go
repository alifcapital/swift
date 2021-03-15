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

// This represents the details of Payment Purpose Validation service request. 
type PaymentPurposeValidation1 struct {
	PurposeCode string `json:"purpose_code,omitempty"`
	CountryCode string `json:"country_code"`
	PurposeDescription string `json:"purpose_description,omitempty"`
	ClearingSystemIdentification string `json:"clearing_system_identification,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}
