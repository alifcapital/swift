package swift_sdk

import (
	"context"
	"crypto/hmac"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func newVerificationReqSchema(t *testing.T, uetr string) *VerificationRequestSchema {
	reqBody, err := NewVerificationRequestSchema(
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
	if err != nil {
		t.Fatal(err)
	}
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

	if !reflect.DeepEqual(&expected, reqSchema) {
		t.Fatalf("expected:\t%v\ngot:\t%v", expected, reqSchema)
	}
}

// ------------------------------------------------------------------

func TestNewVerificationRequest_Fails(t *testing.T) {
	uetr := uuid.New().String()
	_, err := NewVerificationRequestSchema(
		"",
		VerCtxPAYM,
		uetr,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	)
	if err == nil || !errors.Is(err, ErrCorrelationIdentifierOutOfRange) {
		t.Fatalf("expected err: ErrCorrelationIdentifierOutOfRange")
	}

	_, err = NewVerificationRequestSchema(
		"uuid-or-any-string-from-1-to-50-symbols",
		VerCtxPAYM,
		uetr,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	)
	if err == nil || !errors.Is(err, ErrCreditorAccountOutOfRange) {
		t.Fatalf("expected err: ErrCreditorAccountOutOfRange")
	}

	_, err = NewVerificationRequestSchema(
		"uuid-or-any-string-from-1-to-50-symbols",
		VerCtxPAYM,
		uetr,
		"ВНот-1-до-34-символов-ыц;Ы!ЫЦЧФБЬЫ",
		"ВНот-max-140-символов-ыц;Ы!ЫЦЧФБЬЫВНот-1-до-34-символов-ыц;Ы!ЫЦЧФБЬЫВНот-1-до-34-символов-ыц;Ы!ЫЦЧФБЬЫВНот-1-до-34-символов-ыц;Ы!ЫЦЧФБЬЫВНот-1-до-34-символов-ыц;Ы!ЫЦЧФБЬЫ",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	)
	if err == nil || !errors.Is(err, ErrCreditorNameOutOfRange) {
		t.Fatalf("expected err: ErrCreditorNameOutOfRange")
	}
}

// ------------------------------------------------------------------

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
	if err != nil {
		t.Fatal(err)
	}

	reqSchemaJson, err := json.Marshal(reqSchema)
	if err != nil {
		t.Fatal(err)
	}
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
	if !hmac.Equal([]byte(signature), []byte(httpReq.Header.Get("LAUSignature"))) {
		t.Fatalf("LAUSignature not correct")
	}

	if conf.LauAppId != httpReq.Header.Get("LAUApplicationID") {
		t.Fatalf("expected:\t%s\ngot:\t%s", "001", httpReq.Header.Get("LAUApplicationID"))
	}
	if conf.LauVersion != httpReq.Header.Get("LAUVersion") {
		t.Fatalf("expected:\t%s\ngot:\t%s", "001", httpReq.Header.Get("LAUVersion"))
	}
	if lauReqNonce != httpReq.Header.Get("LAURequestNonce") {
		t.Fatalf("expected:\t%s\ngot:\t%s", lauReqNonce, httpReq.Header.Get("LAURequestNonce"))
	}
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
	if err != nil {
		t.Fatal(err)
	}

	if verReq.CorrelationIdentifier != verResp.CorrelationIdentifier {
		t.Fatalf("expected:\t%s\ngot:\t%s", verReq.CorrelationIdentifier, verResp.CorrelationIdentifier)
	}
	if "PASS" != verResp.Response.AccountValidationStatus {
		t.Fatalf("expected:\t%s\ngot:\t%s", "PASS", verResp.Response.AccountValidationStatus)
	}
	if "MTCH" != verResp.Response.CreditorAccountMatch {
		t.Fatalf("expected:\t%s\ngot:\t%s", "MTCH", verResp.Response.CreditorAccountMatch)
	}
	if "NOAP" != verResp.Response.CreditorNameMatch {
		t.Fatalf("expected:\t%s\ngot:\t%s", "NOAP", verResp.Response.CreditorNameMatch)
	}
	if "NOAP" != verResp.Response.CreditorAddressMatch {
		t.Fatalf("expected:\t%s\ngot:\t%s", "NOAP", verResp.Response.CreditorAddressMatch)
	}
	if "NOAP" != verResp.Response.CreditorOrganisationIdentificationMatch {
		t.Fatalf("expected:\t%s\ngot:\t%s", "NOAP", verResp.Response.CreditorOrganisationIdentificationMatch)
	}
}
