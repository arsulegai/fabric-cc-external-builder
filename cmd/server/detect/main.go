package main

import (
	"errors"
	"external-builder/pkg/chaincode"
	"external-builder/pkg/logger"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	BuilderName    = "Chaincode Server Builder"
	BuilderVersion = "1.0"
)

func init() {
	logger.InitLogger("detect")
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(chaincode.Failure)
		}
		log.Infof("successfully detected %s:%s",
			BuilderName, BuilderVersion)
		os.Exit(chaincode.Success)
	}()

	log.Debugf("detect called - step 1 of 4")
	// detect CHAINCODE_SOURCE_DIR CHAINCODE_METADATA_DIR

	// Input validation
	if len(os.Args) != 3 {
		err = errors.New("invalid input, expected 2 parameters")
		return
	}
	chaincodeSourceDir := os.Args[1]
	chaincodeMetadataDir := os.Args[2]

	err = process(chaincodeSourceDir, chaincodeMetadataDir)
}

func process(chaincodeSourceDir, chaincodeMetadataDir string) error {
	// Process the chaincode, chaincode as a service
	chaincodeServer := chaincode.GetServer(chaincode.Detect, chaincodeSourceDir, chaincodeMetadataDir, "")
	metadata, err := chaincodeServer.GetMetadata()
	if err != nil {
		return err
	}

	if metadata.Type != chaincode.External {
		return fmt.Errorf("skipping the chaincode builder, "+
			"type is %s but expected 'external'", metadata.Type)
	}
	return nil
}
