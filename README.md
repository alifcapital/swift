## SWIFT SDK in GO
This package is written to cover low level stuff for integration with [SWIFT Payments](https://swift.com).
All code is written according to specification from [Developers API](https://developer.swift.com/api), 
checkout the link to discover more. For the moment package implements following:
- [x] [Authentication](https://developer.swift.com/oauth-reference#section/Authentication)
- [x] [Pre-Validation](https://developer.swift.com/api/prevalidation/overview)
- [ ] [SWIFTRef](https://developer.swift.com/api/swiftref/overview)

## INSTALL
```bash
go get github.com/alifcapital/swift_sdk
```

## DOCS
Official Swift Developer Portal https://developer.swift.com/reference. See tests and implementation specifics to discover more.

## TESTS
```bash
// provide `tests.json` file (example: `tests.json.example`) and run
go test -cover

// to see coverage details
go tool cover -html tests.cover
```

## Contributing
Pull requests are welcome. For any changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
