package dfs

import (
	"github.com/fairdatasociety/fairOS-dfs/pkg/blockstore/bee/mock"
	"github.com/fairdatasociety/fairOS-dfs/pkg/logging"
	"github.com/fairdatasociety/fairOS-dfs/pkg/user"
	"github.com/spf13/afero"
)

// NewMockDfsAPI is the main entry point for the df controller with a mock bee client.
func NewMockDfsAPI(dataDir, apiUrl, debugApiUrl, postageBlockId string, logger logging.Logger) (*DfsAPI, error) {
	c := mock.NewMockBeeClient()
	if !c.CheckConnection() {
		return nil, ErrBeeClient
	}
	fs := afero.NewMemMapFs()
	users := user.NewUsers(dataDir, c, logger, fs)
	return &DfsAPI{
		dataDir: dataDir,
		client:  c,
		users:   users,
		logger:  logger,
		os:      fs,
	}, nil
}
