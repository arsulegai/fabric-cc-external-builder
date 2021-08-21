package chaincode

const (
	Success = 0
	Failure = 1

	// `metadata.json` file

	MetadataFile = "metadata.json"

	// `connection.json` file

	ConnectionFile = "connection.json"

	// `META-INF` folder

	MetaInf = "META-INF"

	// StateDB folder

	StateDB = "statedb"

	// Release directory

	ReleaseDir = "chaincode/server/"

	// ChaincodeServer/External builder

	External = "external"

	// Environment variables for chaincode TLS

	CoreChaincodeAddress            = "CORE_CHAINCODE_ADDRESS"
	CoreChaincodeTimeout            = "CORE_CHAINCODE_TIMEOUT"
	CoreChaincodeTlsClientKey       = "CORE_CHAINCODE_TLS_CLIENT_KEY"
	CoreChaincodeTlsClientCert      = "CORE_CHAINCODE_TLS_CLIENT_CERT"
	CoreChaincodeTlsRequired        = "CORE_CHAINCODE_TLS_REQUIRED"
	CoreChaincodeClientAuthRequired = "CORE_CHAINCODE_CLIENT_AUTH_REQUIRED"
	CoreChaincodeRootCert           = "CORE_CHAINCODE_ROOT_CERT"
)
