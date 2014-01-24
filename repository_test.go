package ressource

import (
	"fmt"
	"github.com/chrstphlbr/testHelpers"
	"testing"
)

const (
	filesDirectory = "../files"
	greetingJson   = `{
		"hello": {
			"en": "hello",
			"de": "hallo",
			"ru": "Привет"
		}
	}`
)

func TestFileRessourceRepositoryUpdate(t *testing.T) {
	fileName := fmt.Sprintf("%s/%s", filesDirectory, "greeting.json")
	// create file
	testHelpers.CreateFile(t, fileName, greetingJson)
	// delete file
	defer testHelpers.RemoveFile(t, fileName)

	repo := NewFileRepository(filesDirectory)
	// should update ressources to one RessourceAdapter
	repo.Update()
	size := len(repo.ressources)
	if size != 1 {
		// not correct size, should be 1
		t.Fatalf("ressource size should be 1 but was actually %d\n", size)
	}
}
