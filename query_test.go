package locator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var query QueryLocator

func TestMain(m *testing.M) {
	query = NewQuery("testdata")

	code := m.Run()

	os.Exit(code)
}

func TestStruct_Load(t *testing.T) {
	want := "SELECT * FROM posts WHERE id = ? ORDER BY id DESC"
	got, err := query.Load("select_posts_by_id")
	assert.Equal(t, want, got, "Could not get the SQL for the specified file")
	assert.Nil(t, err, "Couldn't read the SQL")
}

func TestStruct_NotExistsFileLoad(t *testing.T) {
	got, err := query.Load("select_posts_by_name")
	assert.Empty(t, got, "blank letter didn't come back")
	assert.Error(t, err, "did not receive an error")
}
