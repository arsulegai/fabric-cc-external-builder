package main

import (
	"errors"
	"external-builder/pkg/chaincode"
	"external-builder/pkg/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

const (
	BuilderName    = "Chaincode Server Builder"
	BuilderVersion = "1.0"
)

func init() {
	logger.InitLogger("release")
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Errorf(err.Error())
			os.Exit(chaincode.Failure)
		}
		log.Infof("successfully released %s:%s",
			BuilderName, BuilderVersion)
		os.Exit(chaincode.Success)
	}()

	log.Debugf("release called - step 3 of 4")
	// release BUILD_OUTPUT_DIR RELEASE_OUTPUT_DIR

	// Input validation
	if len(os.Args) != 3 {
		err = errors.New("invalid input, expected 2 parameters")
		return
	}
	buildOutputDir := os.Args[1]
	releaseOutputDir := os.Args[2]

	err = process(buildOutputDir, releaseOutputDir)
}

func process(buildOutputDir, releaseOutputDir string) error {
	// Copy the connection.json file to `chaincode/server/`
	// in the release directory
	connectionSrcFile := filepath.Join(buildOutputDir, chaincode.ConnectionFile)
	connectionDstFile := filepath.Join(releaseOutputDir, chaincode.ReleaseDir, chaincode.ConnectionFile)
	pathDir := filepath.Base(connectionDstFile)
	err := os.MkdirAll(pathDir, os.ModeDir)
	if err != nil {
		return err
	}
	// Create directories for release if not already present
	err = copy.Copy(connectionSrcFile, connectionDstFile)
	if err != nil {
		return err
	}

	// Copy the metadata inf directory, note that the release artefacts are
	// consumed by the peer node.
	metaInfSrcDir := filepath.Join(buildOutputDir, chaincode.MetaInf, chaincode.StateDB)
	metaInfDstDir := filepath.Join(releaseOutputDir, chaincode.StateDB)
	_, err = os.Stat(metaInfSrcDir)
	if !os.IsNotExist(err) {
		// Create directory for destination
		err = os.MkdirAll(metaInfDstDir, os.ModeDir)
		if err != nil {
			return err
		}
		err = copy.Copy(metaInfSrcDir, metaInfDstDir)
		if err != nil {
			return err
		}
	}
	return nil
}
