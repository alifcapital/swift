package swift

import "fmt"

// env - sandbox (testing) or production
// choose from options below (const list)
type env string

// verificationContext - `what is verification intention?`
// see descriptions and choose form options below
type verificationContext string

const (
	// SandBoxEnv - testing environment, but can be used in production for some services
	// from docs:  ... KYC Registry API, Banking Analytics API, Banking Premium API and Compliance Analytics API, Sandbox CAN be used for UAT.
	// https://developer.swift.com/support
	SandBoxEnv = env("sandbox")
	// ProductionEnv - production environment
	ProductionEnv = env("production")

	// VerCtxBENR - Beneficiary registration. The beneficiary is being verified before being added to a
	//master data of beneficiary accounts with which the debtor has business relationships.
	VerCtxBENR = verificationContext("BENR")
	// VerCtxPAYM - The account verification is performed in the context of a Payment transaction,
	//before it is issued for execution.
	VerCtxPAYM = verificationContext("PAYM")
	// VerCtxRFPP - A request for payment transaction
	//(the debtor account is being verified before the request is issued or the direct debit mandate is created)
	VerCtxRFPP = verificationContext("RFPP")
)

// ------------------------------------------------------------------

func accessTokenUrl(e env) string {
	if e == ProductionEnv {
		return "https://api.swift.com/oauth2/v1/token"
	}
	return "https://sandbox.swift.com/oauth2/v1/token"
}

func revokeTokenUrl(e env) string {
	if e == ProductionEnv {
		return "https://api.swift.com/oauth2/v1/revoke"
	}
	return "https://sandbox.swift.com/oauth2/v1/revoke"
}

func verificationUrl(e env) string {
	if e == ProductionEnv {
		return "https://api.swift.com/swift-preval-pilot/v1/accounts/verification"
	}
	return "https://sandbox.swift.com/swift-preval-pilot/v1/accounts/verification"
}

func bicDetailsUrl(e env, bic string) string {
	if e == ProductionEnv {
		return fmt.Sprintf("https://api.swift.com/swiftref-api/bics/%s", bic)
	}
	return fmt.Sprintf("https://sandbox.swift.com/swiftref-api/bics/%s", bic)
}

func validIBANUrl(e env, iban string) string {
	if e == ProductionEnv {
		return fmt.Sprintf("https://api.swift.com/swiftref-api/ibans/%s/validity", iban)
	}
	return fmt.Sprintf("https://sandbox.swift.com/swiftref-api/ibans/%s/validity", iban)
}

func detailsOfIBANUrl(e env, iban string) string {
	if e == ProductionEnv {
		return fmt.Sprintf("https://api.swift.com/swiftref-api/ibans/%s", iban)
	}
	return fmt.Sprintf("https://sandbox.swift.com/swiftref-api/ibans/%s", iban)
}

func bicByIBANUrl(e env, iban string) string {
	if e == ProductionEnv {
		return fmt.Sprintf("https://api.swift.com/swiftref-api/ibans/%s/bic", iban)
	}
	return fmt.Sprintf("https://sandbox.swift.com/swiftref-api/ibans/%s/bic", iban)
}
