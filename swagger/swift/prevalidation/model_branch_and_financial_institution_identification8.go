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

// Unique and unambiguous identification of a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification8 struct {
	FinancialInstitutionIdentification *FinancialInstitutionIdentification24 `json:"financial_institution_identification"`
	BranchIdentification *BranchData3 `json:"branch_identification,omitempty"`
	ContactData *Contact10 `json:"contact_data,omitempty"`
	ConnectivityData *GenericConnectivityChoice `json:"connectivity_data,omitempty"`
	OfficeType string `json:"office_type,omitempty"`
}
