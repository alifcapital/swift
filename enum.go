package swift

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
