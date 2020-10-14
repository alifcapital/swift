package swift_sdk

// env - sandbox (testing) or production
// choose from options below (const list)
type env string

// verificationContext - `what is verification intention?`
// see descriptions and choose form options below
type verificationContext string

const (
	// only available options by the time of
	SandBoxEnv    = env("sandbox")
	ProductionEnv = env("production")

	//'BENR' - Beneficiary registration. The beneficiary is being verified before being added to a
	//master data of beneficiary accounts with which the debtor has business relationships.
	VerCtxBENR = verificationContext("BENR")
	// 'PAYM' - The account verification is performed in the context of a Payment transaction,
	//before it is issued for execution.
	VerCtxPAYM = verificationContext("PAYM")
	// RFPP' - A request for payment transaction
	//(the debtor account is being verified before the request is issued or the direct debit mandate is created)
	VerCtxRFPP = verificationContext("RFPP")
)
