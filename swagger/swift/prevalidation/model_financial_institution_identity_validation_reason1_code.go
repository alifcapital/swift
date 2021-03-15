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
// FinancialInstitutionIdentityValidationReason1Code : The list of unique identifiers which describe the reasons for Financial Institution Identity validation status. *`VBIC_valid_bic`-The instructed financial institution identifier subject to financial institution identity validation is a valid BIC for the instructed payment context. *`IBIC_invalid_bic`-The instructed financial institution identifier subject to financial institution identity validation is not a  valid BIC for the instructed payment context. *`INCT_inconsistent_country`-The instructed country code does not match the country of the payment information that is subject to validation. *`CCTY_consistent_country`-The instructed country code matches the country of the payment information that is subject to validation. *`VCSM_valid_clearing_system_member_identification`-The instructed financial institution identifier subject to financial institution identity validation is a valid clearing system member identifier for the instructed payment context. *`UCSM_unrecognized_clearing_system_member_identification`-The instructed financial institution identifier subject to financial institution identity validation has not been found in the SWIFT directory for the instructed payment context. *`ICTY_invalid_country`-Instructed country code is not a valid country code on the basis of country names obtained from the United Nations (ISO 3166, Alpha-2 code) *`MCTY_missing_country_code`-Country code is missing from the payment context in order for the validation of the instructed payment information to be performed.
type FinancialInstitutionIdentityValidationReason1Code string

// List of FinancialInstitutionIdentityValidationReason1Code
const (
	VBIC_VALID_BIC_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "VBIC_valid_bic"
	IBIC_INVALID_BIC_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "IBIC_invalid_bic"
	INCT_INCONSISTENT_COUNTRY_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "INCT_inconsistent_country"
	CCTY_CONSISTENT_COUNTRY_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "CCTY_consistent_country"
	VCSM_VALID_CLEARING_SYSTEM_MEMBER_IDENTIFICATION_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "VCSM_valid_clearing_system_member_identification"
	UCSM_UNRECOGNIZED_CLEARING_SYSTEM_MEMBER_IDENTIFICATION_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "UCSM_unrecognized_clearing_system_member_identification"
	ICTY_INVALID_COUNTRY_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "ICTY_invalid_country"
	MCTY_MISSING_COUNTRY_CODE_FinancialInstitutionIdentityValidationReason1Code FinancialInstitutionIdentityValidationReason1Code = "MCTY_missing_country_code"
)