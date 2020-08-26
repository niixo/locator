package locator

import (
	"io/ioutil"
	"path/filepath"
)

// QueryLocator
type QueryLocator interface {
	Load(query string) (string, error)
}

// queryLocator
type queryLocator struct {
	fileDir string
}

// New returns a new instance of query locator
func New(fileDir string) QueryLocator {
	return &queryLocator{fileDir: fileDir}
}

// Load Reads SQL from a sql file
func (q *queryLocator) Load(query string) (string, error) {
	p := filepath.Join(q.fileDir, query+".sql")
	sql, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}

	return string(sql), nil
}
