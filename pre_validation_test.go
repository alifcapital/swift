package swift

import (
	"context"
	"crypto/hmac"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newVerificationReqSchema(t *testing.T, uetr string) *VerificationRequestSchema {
	reqBody := NewVerificationRequestSchema(
		"SCENARIO1-CORRID-001",
		VerCtxPAYM,
		uetr,
		"GB3112000000001987426375",
		"John Doe",
		"GB",
		"London",
		"",
		"",
		"BANABEBB",
		"BANABEBB",
		"",
	)
	return reqBody
}

// ------------------------------------------------------------------

func TestNewVerificationRequestSchema(t *testing.T) {
	uetr := uuid.New().String()
	reqSchema := newVerificationReqSchema(t, uetr)

	expected := VerificationRequestSchema{
		CorrelationIdentifier: "SCENARIO1-CORRID-001",
		Context:               string(VerCtxPAYM),
		Uetr:                  uetr,
		CreditorAccount:       "GB3112000000001987426375",
		CreditorName:          "John Doe",
		CreditorAddress: CreditorAddress{
			Country:     "GB",
			AddressLine: []string{"London"},
			PostCode:    "",
			TownName:    "",
		},
		CreditorOrganisationIdentification: CreditorOrganisationIdentification{AnyBic: "BANABEBB"},
		CreditorAgent:                      CreditorAgent{Bicfi: "BANABEBB"},
		CreditorAgentBranchIdentification:  "",
	}

	assert.Equal(t, &expected, reqSchema)
}

func TestNewVerificationRequest(t *testing.T) {
	uetr := uuid.New().String()
	reqSchema := newVerificationReqSchema(t, uetr)

	conf := newTestsConfig()
	lauReqNonce := uuid.New().String()
	httpReq, err := NewVerificationRequest(
		reqSchema,
		conf.LauAppId,
		lauReqNonce,
		conf.LauVersion,
		conf.LauHmacKey,
		conf.AppApiKey,
		conf.Xbic,
		conf.XApiKey,
		SandBoxEnv,
	)
	assert.Nil(t, err)

	reqSchemaJson, err := json.Marshal(reqSchema)
	assert.Nil(t, err)
	signature := makeLauSignature(
		conf.LauAppId,
		httpReq.Header.Get("LAUCallTime"),
		lauReqNonce,
		httpReq.Header.Get("LAUSigned"),
		conf.LauVersion,
		"/swift-preval-pilot/v1/accounts/verification",
		conf.LauHmacKey,
		reqSchemaJson,
	)
	assert.True(t, hmac.Equal([]byte(signature), []byte(httpReq.Header.Get("LAUSignature"))))

	assert.Equal(t, conf.LauAppId, httpReq.Header.Get("LAUApplicationID"))
	assert.Equal(t, conf.LauVersion, httpReq.Header.Get("LAUVersion"))
	assert.Equal(t, lauReqNonce, httpReq.Header.Get("LAURequestNonce"))
}

func TestPerformPreValidationCheck(t *testing.T) {
	uetr := uuid.New().String()
	verReq := newVerificationReqSchema(t, uetr)
	ctx := context.Background()
	config := newTestsConfig()
	e := SandBoxEnv

	verResp, err := PerformPreValidationCheck(
		verReq,
		ctx,
		config.LauAppId,
		uuid.New().String(),
		config.LauVersion,
		config.LauHmacKey,
		config.AppApiKey,
		config.Xbic,
		config.XApiKey,
		e,
	)
	assert.Nil(t, err)

	assert.Equal(t, verReq.CorrelationIdentifier, verResp.CorrelationIdentifier)
	assert.Equal(t, "PASS", verResp.Response.AccountValidationStatus)
	assert.Equal(t, "MTCH", verResp.Response.CreditorAccountMatch)
	assert.Equal(t, "NOAP", verResp.Response.CreditorNameMatch)
	assert.Equal(t, "NOAP", verResp.Response.CreditorAddressMatch)
	assert.Equal(t, "NOAP", verResp.Response.CreditorOrganisationIdentificationMatch)
}
