package swift_sdk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

const CRLF = "\r\n"

// VerificationRequestSchema - this data send to endpoint to get information about SWIFT account
// not all fields required, same fields that not going to be used at all not implemented to save some time
// docs: https://developer.swift.com/reference#operation/VerifyAccount
type VerificationRequestSchema struct {
	/**
	The identifier of the request, assigned by the sender of the request.
	It will allow the requester to correlate the request and response.
	Specifies a character string with a maximum length of 50 characters.
	*/
	CorrelationIdentifier string `json:"correlation_identifier"`

	/**
	required
	Enum:"BENR" "PAYM" "RFPP"
	Transaction Context:
	'BENR' - Beneficiary registration.
	The beneficiary is being verified before being added to a master data of beneficiary accounts
	with which the debtor has business relationships.

	'PAYM' - The account verification is performed in the context of a Payment transaction,
	before it is issued for execution.

	'RFPP' - A request for payment transaction,
	the debtor account is being verified before the request is issued or the direct debit mandate is created.
	*/
	Context string `json:"context"`

	/**
	UETR of the transaction that is going to be sent to this account.
	Universally Unique IDentifier (UUID) version 4, as described in IETC RFC 4122
	"Universally Unique IDentifier (UUID) URN Namespace".
	*/
	Uetr string `json:"uetr"`

	/**
	required
	Identifies the account targeted by the transaction. Specifies a character string with a maximum length of 34 characters.
	*/
	CreditorAccount string `json:"creditor_account"`

	/**
	required
	Name by which the creditor is known. Specifies a character string with a maximum length of 140 characters.
	*/
	CreditorName string `json:"creditor_name"`

	/**
	Information that locates and identifies the address of the creditor, as defined by postal services.
	*/
	CreditorAddress CreditorAddress `json:"creditor_address"`

	/**
	Unique and unambiguous way to identify a creditor.
	*/
	CreditorOrganisationIdentification CreditorOrganisationIdentification `json:"creditor_organisation_identification"`

	/**
	Set of elements used to identify the financial institution servicing an account for the creditor.
	*/
	CreditorAgent CreditorAgent `json:"creditor_agent"`

	/**
	Max35Text
	Identifies a specific branch of the creditor.
	*/
	CreditorAgentBranchIdentification string `json:"creditor_agent_branch_identification"`
}

// CreditorAddress - Information that locates and identifies the address of the creditor, as defined by postal services.
type CreditorAddress struct {
	/**
	/^[A-Z]{2,2}$/
	Nation with its own government. Code to identify a country, a dependency,
	or another area of particular geopolitical interest,
	on the basis of country names obtained from the United Nations (ISO 3166, Alpha-2 code).
	*/
	Country string `json:"country"`

	/**
	Max70Text
	Information that locates and identifies a specific address,
	as defined by postal services, presented in free format text.
	*/
	AddressLine []string `json:"address_line"`

	/**
	Identifier consisting of a group of letters and/or numbers
	that is added to a postal address to assist the sorting of mail.
	Specifies a character string with a maximum length of 16 characters.
	*/
	PostCode string `json:"post_code"`

	/**
	Max35Text
	Name of a built-up area, with defined boundaries, and a local government.
	*/
	TownName string `json:"town_name"`
}

// CreditorOrganisationIdentification - Unique and unambiguous way to identify a creditor.
type CreditorOrganisationIdentification struct {
	AnyBic string `json:"any_bic"`

	/**
	TODO:
	Array of object (GenericOrganisationIdentification1)
	Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	other
	*/
}

// CreditorAgent - Set of elements used to identify the financial institution servicing an account for the creditor.
type CreditorAgent struct {
	/**
	^[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}$
	Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362
	"Banking - Banking telecommunication messages - Business identifier code (BIC)".
	*/
	Bicfi string `json:"bicfi"`

	/**
	TODO:
	Unique identification, as assigned by a clearing system, to unambiguously identify a member of the clearing system.
	clearing_system_member_identification
	*/
}

// NewVerificationRequestSchema - init new request instance
func NewVerificationRequestSchema(
	correlationIdentifier string,
	verificationCtx verificationContext,
	uetr,
	creditorAccount,
	creditorName,
	country,
	addressLine,
	postCode,
	townName,
	anyBic,
	bicfi,
	creditorAgentBranchIdentification string,
) *VerificationRequestSchema {
	return &VerificationRequestSchema{
		CorrelationIdentifier: correlationIdentifier,
		Context:               string(verificationCtx),
		Uetr:                  uetr,
		CreditorAccount:       creditorAccount,
		CreditorName:          creditorName,
		CreditorAddress: CreditorAddress{
			Country:     country,
			AddressLine: strings.Split(addressLine, ","),
			PostCode:    postCode,
			TownName:    townName,
		},
		CreditorOrganisationIdentification: CreditorOrganisationIdentification{AnyBic: anyBic},
		CreditorAgent:                      CreditorAgent{Bicfi: bicfi},
		CreditorAgentBranchIdentification:  creditorAgentBranchIdentification,
	}
}

// ------------------------------------------------------------------

// VerificationResponse - RESPONSE SCHEMA: application/json
type VerificationResponse struct {
	/**
	The identifier of the request, assigned by the sender of the request.
	It will allow the requester to correlate the request and response.
	Specifies a character string with a maximum length of 50 characters.
	*/
	CorrelationIdentifier string   `json:"correlation_identifier"`
	Response              response `json:"response"`
}

type response struct {
	/**
	'PASS' - Passed. It is possible to credit a transaction (as per the input request) to an account
	identified with this account identification and creditor information provided in the response.

	'INCO' - Incomplete. The account is in the beneficiary bank's books however the creditor
	name has not been checked at this stage and will be checked at payment transaction time.

	'FAIL' - Failed. The service provider would not be able to execute a transaction (as per the input request)
	to an account identified by the creditor account and creditor details provided in the request.
	*/
	AccountValidationStatus string `json:"account_validation_status"`

	/**
	'NMTC' - Mismatch. The creditor information passed in the request did not match the account holder details.
	'MTCH' - Match. The creditor information passed in the request matches the account holder details.
	*/
	CreditorAccountMatch string `json:"creditor_account_match"`

	/**
	if creditor_account_match is MISMATCH, creditor_name_match must contain NOT_CHECKED.
	if the creditor_account_match is NO, creditor_address_match must contain NOT_APPLICABLE.
	if the creditor_account_match is NO, creditor_organization_identification_match must contain NOT_APPLICABLE.

	'NOTC' - Not Checked. The service provider did not verify the creditor
	information from the request however this will be checked at payment processing time.

	'NOAP' - Not Applicable. Return this value only if the creditor information is not checked at this stage
	and will not be checked at payment processing time OR the creditor information was not passed in the request.

	'NMTC' - Mismatch. The creditor information passed in the request did not match the account holder details.

	'MTCH' - Match. The creditor information passed in the request matches the account holder details.
	*/
	CreditorNameMatch string `json:"creditor_name_match"`

	/**
	if creditor_account_match is MISMATCH, creditor_name_match must contain NOT_CHECKED.
	if the creditor_account_match is NO, creditor_address_match must contain NOT_APPLICABLE.
	if the creditor_account_match is NO, creditor_organization_identification_match must contain NOT_APPLICABLE.

	'NOTC' - Not Checked. The service provider did not verify the creditor information
	from the request however this will be checked at payment processing time.

	'NOAP' - Not Applicable. Return this value only if the creditor information is not checked at this stage
	and will not be checked at payment processing time OR the creditor information was not passed in the request.

	'NMTC' - Mismatch. The creditor information passed in the request did not match the account holder details.

	'MTCH' - Match. The creditor information passed in the request matches the account holder details.
	*/
	CreditorAddressMatch string `json:"creditor_address_match"`

	/**
	if creditor_account_match is MISMATCH, creditor_name_match must contain NOT_CHECKED.
	if the creditor_account_match is NO, creditor_address_match must contain NOT_APPLICABLE.
	if the creditor_account_match is NO, creditor_organization_identification_match must contain NOT_APPLICABLE.

	'NOTC' - Not Checked. The service provider did not verify the creditor information from the request
	however this will be checked at payment processing time.

	'NOAP' - Not Applicable. Return this value only if the creditor information is not checked at this
	stage and will not be checked at payment processing time OR the creditor information was not passed in the request.

	'NMTC' - Mismatch. The creditor information passed in the request did not match the account holder details.

	'MTCH' - Match. The creditor information passed in the request matches the account holder details.
	*/
	CreditorOrganisationIdentificationMatch string `json:"creditor_organisation_identification_match"`
}

// ------------------------------------------------------------------

// NewVerificationRequest - constructs new request for verifying SWIFT account
// docs: https://developer.swift.com/reference#operation/VerifyAccount
// lauAppId - "001"
// lauReqNonce - a random value generated by the client consumer
// lauVersion - "1.0"
// lauHmacKey - key for generating hmac signature
// appApiKey - part of `LAUSigned` header
// xbic - part of `LAUSigned` header && x-bik header
// xApiKey - {{consumer-key}} for
func NewVerificationRequest(
	verReqSchema *VerificationRequestSchema,
	lauAppId string,
	lauReqNonce string,
	lauVersion string,
	lauHmacKey string,
	appApiKey string,
	xbic string,
	xApiKey string,
	e env,
) (*http.Request, error) {
	// request content-type is application/json
	// so convert to json
	reqSchema, err := json.Marshal(verReqSchema)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	_, err = b.Write(reqSchema)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, getVerificationUrl(e), &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	/**
	required
	string^(?:[a-zA-Z0-9]{1,16})$
	Default: "001"
	ID that identifies the application generating the API and used by the gpi Connector
	to retrieve the related LAU keys, required for consumers using gpi connector only.
	*/
	lauAppId = strings.TrimSpace(lauAppId)
	if lauAppId == "" {
		lauAppId = "001"
	}
	req.Header.Set("LAUApplicationID", lauAppId)

	/**
	required
	Default: "1.0"
	Version of the LAUSigned header.
	Mandatory. "1.0" for this first release, required for consumers using gpi connector only.
	*/
	lauVersion = strings.TrimSpace(lauVersion)
	if lauVersion == "" {
		lauVersion = "1.0"
	}
	req.Header.Set("LAUVersion", lauVersion)

	/**
	required
	Example: "2018-03-23T15:56:26.728Z"
	Timestamp in UTC of the API call in the format YYYY-MM-DDTHH:MM:SS.sssZ,
	required for consumers using gpi connector only.
	*/
	lauCallTime := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	req.Header.Set("LAUCallTime", lauCallTime)

	/**
	required
	Example: "e802ab96-bb3a-4965-9139-5214b9f0f074"
	A random value generated by the client consumer. Provided with the request and
	copied by the gpi Connector on the response, required for consumers using gpi connector only.
	*/
	lauReqNonce = strings.TrimSpace(lauReqNonce)
	if lauReqNonce == "" {
		lauReqNonce = uuid.New().String()
	}
	req.Header.Set("LAURequestNonce", lauReqNonce)

	/**
	required
	Example: "(ApplAPIKey=yVGhKiV5z1ZGdaqFXoZ8AiSA9n5CrY6B),(x-bic=cclabeb0)"
	Includes the signed headers to be passed to SWIFT in the request.
	The first is the application API Key (ApplAPIKey).
	The second parameter is the x-bic. It identifies the destination of the request.
	*/
	lauSigned := fmt.Sprintf("(ApplAPIKey=%s),(x-bic=%s))", appApiKey, xbic)
	req.Header.Set("LAUSigned", lauSigned)

	// `signature.sigBytes = 16` for some mtf reason...
	lauSignature := makeLauSignature(lauAppId, lauCallTime, lauReqNonce, lauSigned, lauVersion, req.URL.RequestURI(), lauHmacKey, b.Bytes())
	req.Header.Set("LAUSignature", lauSignature)

	// authentication
	req.Header.Set("x-api-key", xApiKey)

	req.Header.Set("x-bic", xbic)

	return req, nil
}

// makeLauSignature - generates HMAC-SHA256 signature according to docs
func makeLauSignature(
	lauAppId, lauCallTime, lauReqNonce, lauSigned, lauVersion, verificationUrl, lauHmacKey string,
	reqSchemaJson []byte,
) string {
	var sb strings.Builder
	sb.WriteString("LAUApplicationID:")
	sb.WriteString(lauAppId)
	sb.WriteString(CRLF)
	sb.WriteString("LAUCallTime:")
	sb.WriteString(lauCallTime)
	sb.WriteString(CRLF)
	sb.WriteString("LAURequestNonce:")
	sb.WriteString(lauReqNonce)
	sb.WriteString(CRLF)
	sb.WriteString("LAUSigned:")
	sb.WriteString(lauSigned)
	sb.WriteString(CRLF)
	sb.WriteString("LAUVersion:")
	sb.WriteString(lauVersion)
	sb.WriteString(CRLF)
	sb.WriteString(verificationUrl)
	sb.WriteString(CRLF)

	// ... from request-body
	sb.Write(reqSchemaJson)

	// generate hmac signature
	h := hmac.New(sha256.New, []byte(lauHmacKey))
	h.Write([]byte(sb.String()))
	sum := h.Sum(nil)

	// `signature.sigBytes = 16` for some mtf reason...
	return base64.StdEncoding.EncodeToString(sum[:len(sum)/2])
}

// ------------------------------------------------------------------

// PerformPreValidationCheck - makes an http call to SWIFT api, `VerificationResponse` returned
// caller responsible for further processing the response to determine `success`
// lauAppId - "001"
// lauReqNonce - a random value generated by the client consumer
// lauVersion - "1.0"
// lauHmacKey - key for generating hmac signature
// appApiKey - part of `LAUSigned` header
// xbic - part of `LAUSigned` header && x-bik header
// xApiKey - {{consumer-key}} for
func PerformPreValidationCheck(
	verReq *VerificationRequestSchema,
	ctx context.Context,
	lauAppId string,
	lauReqNonce string,
	lauVersion string,
	lauHmacKey string,
	appApiKey string,
	xbic string,
	xApiKey string,
	e env,
) (*VerificationResponse, error) {
	// initialize httpRequest instance and set given context
	httpReq, err := NewVerificationRequest(verReq, lauAppId, lauReqNonce, lauVersion, lauHmacKey, appApiKey, xbic, xApiKey, e)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	// perform request with default client
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response body
	var b bytes.Buffer
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	// check for status-code being 200
	if resp.StatusCode != http.StatusOK {
		return nil, NewSWIFTError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response to appropriate structure
	verResp := VerificationResponse{}
	if err := json.Unmarshal(b.Bytes(), &verResp); err != nil {
		return nil, err
	}
	return &verResp, nil
}
