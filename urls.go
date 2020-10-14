package swift_sdk

func getAccessTokenUrl(e env) string {
	if e == ProductionEnv {
		return "https://api.swift.com/oauth2/v1/token"
	}
	return "https://sandbox.swift.com/oauth2/v1/token"
}

func getRevokeTokenUrl(e env) string {
	if e == ProductionEnv {
		return "https://api.swift.com/oauth2/v1/revoke"
	}
	return "https://sandbox.swift.com/oauth2/v1/revoke"
}

func getVerificationUrl(e env) string {
	if e == ProductionEnv {
		return "https://api.swift.com/swift-preval-pilot/v1/accounts/verification"
	}
	return "https://sandbox.swift.com/swift-preval-pilot/v1/accounts/verification"
}
