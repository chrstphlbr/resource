package ressource

import (
	"io/ioutil"
	"testing"
)

func TestFileAdapterOpenFile(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	// create file adapter to ./temp/greeting.json file
	fa := NewFileAdapter(greetingFileName)

	err := fa.openFile()
	if err != nil {
		t.Fatalf("returned error while opening: %v", err)
	}

	if fa.file == nil {
		t.Fatal("opened file correctly but variable file is not set")
	}
}

func TestFileAdapterName(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	// check if name is correct
	fa := NewFileAdapter(greetingFileName)

	if fa.Name() != greetingFileName {
		t.Fatalf("adapter name (%s) is not equal to file name (%s)", fa.Name(), greetingFileName)
	}
}

func TestFileAdapterGet(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	fa := NewFileAdapter(greetingFileName)

	reader, err := fa.Get()
	if err != nil {
		t.Fatalf("returned error while getting reader: %v", err)
	}

	if reader == nil {
		t.Fatal("no error and also no reader returned")
	}

	var content []byte
	content, err = ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal("could not read content from reader: %v", err)
	}
	t.Logf("content: %s", string(content))

}
