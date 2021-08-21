package chaincode

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// BuilderI is the interface for chaincode to be built
type BuilderI interface {
	// GetMetadata returns the metadata.json file
	GetMetadata() (Metadata, error)

	// GetConnection returns the connection.json file
	GetConnection() (Connection, error)

	// SaveConnection saves the connection json file into the output directory
	SaveConnection(connection Connection) error
}

func (s Server) GetMetadata() (Metadata, error) {
	metadataFile := filepath.Join(s.ChaincodeMetadataDir, MetadataFile)
	var metadata Metadata
	bytes, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		return metadata, err
	}
	err = json.Unmarshal(bytes, &metadata)
	return metadata, err
}

func (s Server) GetConnection() (Connection, error) {
	connectionFile := filepath.Join(s.ChaincodeSourceDir, ConnectionFile)
	var connection Connection
	bytes, err := ioutil.ReadFile(connectionFile)
	if err != nil {
		return connection, err
	}
	err = json.Unmarshal(bytes, &connection)
	return connection, err
}

func (s Server) SaveConnection(connection Connection) error {
	connectionFile := filepath.Join(s.BuildOutputDir, ConnectionFile)
	bytes, err := json.Marshal(connection)
	if err != nil {
		return err
	}
	pathDir := filepath.Base(connectionFile)
	err = os.MkdirAll(pathDir, os.ModeDir)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(connectionFile, bytes, fs.ModeAppend)
}
