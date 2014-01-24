package ressource

import (
	"fmt"
	"github.com/chrstphlbr/testHelpers"
	"testing"
)

func TestFileAdapter(t *testing.T) {
	fileName := fmt.Sprintf("%s/%s", filesDirectory, "greeting.json")
	// create file
	testHelpers.CreateFile(t, fileName, greetingJson)
	defer testHelpers.RemoveFile(t, fileName)
}
