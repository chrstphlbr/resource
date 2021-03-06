package resource

import (
	"github.com/chrstphlbr/testHelpers"
	"os"
	"testing"
)

const (
	filesDirectory = "./temp"
	greetingJson   = `{
		"hello": {
			"en": "hello",
			"de": "hallo",
			"ru": "Привет"
		}
	}`
	greetingFileName = "./temp/greeting.json"
)

func setUp(t *testing.T) {
	// create repository
	os.Mkdir(filesDirectory, 0700)
	// create file
	testHelpers.CreateFile(t, greetingFileName, greetingJson)
}

func tearDown(t *testing.T) {
	// delete file
	testHelpers.RemoveFile(t, filesDirectory)
}

func TestFileResourceRepositoryUpdate(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	repo := NewFileRepository(filesDirectory)
	// should update resources to one ResourceAdapter
	repo.Update()
	size := len(repo.Resources())
	if size != 1 {
		// not correct size, should be 1
		t.Fatalf("resource size should be 1 but was actually %d\n", size)
	}
}
