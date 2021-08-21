package main

import (
	"errors"
	"external-builder/pkg/chaincode"
	"external-builder/pkg/env"
	"external-builder/pkg/logger"
	"github.com/otiai10/copy"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
)

const (
	BuilderName    = "Chaincode Server Builder"
	BuilderVersion = "1.0"
)

func init() {
	logger.InitLogger("build")
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(chaincode.Failure)
		}
		log.Infof("successfully built %s:%s",
			BuilderName, BuilderVersion)
		os.Exit(chaincode.Success)
	}()

	log.Debugf("builder called - step 2 of 4")
	// build CHAINCODE_SOURCE_DIR CHAINCODE_METADATA_DIR BUILD_OUTPUT_DIR
	// Input validation
	if len(os.Args) != 4 {
		err = errors.New("invalid input, expected 3 parameters")
		return
	}
	chaincodeSourceDir := os.Args[1]
	chaincodeMetadataDir := os.Args[2]
	buildOutputDir := os.Args[3]
	err = process(chaincodeSourceDir, chaincodeMetadataDir, buildOutputDir)
}

func process(chaincodeSourceDir, chaincodeMetadataDir, buildOutputDir string) error {
	// Process the chaincode, chaincode as a service
	chaincodeServer := chaincode.GetServer(chaincode.Build, chaincodeSourceDir, chaincodeMetadataDir, buildOutputDir)

	// Compose the connection file at this point in time with peer specific
	// environment variables
	connection, err := chaincodeServer.GetConnection()
	if err != nil {
		return err
	}
	err = updateConnection(&connection)
	if err != nil {
		return err
	}
	log.Debugf("successfully updated the connection file %+v", connection)
	// Copy the connection file to the output directory
	err = chaincodeServer.SaveConnection(connection)
	if err != nil {
		return err
	}
	// Copy the metadata indexes to the output directory if it is present
	metaIndexSrcDir := filepath.Join(chaincodeSourceDir, chaincode.MetaInf)
	metaIndexDstDir := filepath.Join(buildOutputDir, chaincode.MetaInf)
	_, err = os.Stat(metaIndexSrcDir)
	if !os.IsNotExist(err) {
		err = copy.Copy(metaIndexSrcDir, metaIndexDstDir)
		if err != nil {
			return err
		}
	}
	// Copy the metadata json file to the output directory
	err = copy.Copy(chaincodeMetadataDir, buildOutputDir)
	if err != nil {
		return err
	}
	return nil
}

func updateConnection(connection *chaincode.Connection) error {
	// Fill in connection file from env
	address, err := env.GetEnvOrError(chaincode.CoreChaincodeAddress)
	if err != nil {
		return err
	}
	timeout, err := env.GetEnvOrError(chaincode.CoreChaincodeTimeout)
	if err != nil {
		return err
	}
	tlsRequired, err := env.GetEnvOrError(chaincode.CoreChaincodeTlsRequired)
	if err != nil {
		return err
	}
	tlsRequiredBool, err := strconv.ParseBool(tlsRequired)
	if err != nil {
		return err
	}
	var clientAuthRequiredBool = false
	var rootCert, clientKey, clientCert string
	if tlsRequiredBool {
		// Read the cert file if TLS is set
		rootCert, err = env.GetEnvOrError(chaincode.CoreChaincodeRootCert)
		if err != nil {
			return err
		}
		// Read client key and client cert
		clientAuthRequired, err := env.GetEnvOrError(chaincode.CoreChaincodeClientAuthRequired)
		if err != nil {
			return err
		}
		clientAuthRequiredBool, err = strconv.ParseBool(clientAuthRequired)
		if err != nil {
			return err
		}
		if clientAuthRequiredBool {
			// Must have a client key and client cert
			clientKey, err = env.GetEnvOrError(chaincode.CoreChaincodeTlsClientKey)
			if err != nil {
				return err
			}
			clientCert, err = env.GetEnvOrError(chaincode.CoreChaincodeTlsClientCert)
			if err != nil {
				return err
			}
		}
	}
	connection.Address = address
	connection.TlsRequired = tlsRequiredBool
	connection.DialTimeout = timeout
	connection.ClientAuthRequired = clientAuthRequiredBool
	connection.ClientCert = clientCert
	connection.ClientKey = clientKey
	connection.RootCert = rootCert
	return nil
}
