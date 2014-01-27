package resource

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type Repository interface {
	Update()
	Resources() []Adapter
}

type FileRepository struct {
	lock      sync.Mutex
	resources []Adapter
	path      string
}

func (repo *FileRepository) Update() {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	log.Println("start update")
	fileDescriptors, err := ioutil.ReadDir(repo.path)
	if err != nil {
		log.Printf("error retrieving file descriptors of directory (%s)\n", repo.path)
		log.Println(err)
		return
	}
	log.Printf("directories has %d files", len(fileDescriptors))
	// dirty implementation: remove all previous resources and add all again
	repo.resources = make([]Adapter, 0, len(fileDescriptors))
	var pathToFile string
	for _, fd := range fileDescriptors {
		if fd.IsDir() {
			// sub directories are not handled
			continue
		}
		// check for .json files
		if !strings.HasSuffix(fd.Name(), ".json") {
			// no .json file
			continue
		}
		pathToFile = fmt.Sprintf("%s/%s", repo.path, fd.Name())
		//log.Println(pathToFile)
		// add RessourceAsapter to resources
		repo.resources = append(repo.resources, NewFileAdapter(pathToFile))
	}
	log.Println("finished update")
}

func (repo FileRepository) Resources() []Adapter {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	return repo.resources
}

func NewFileRepository(path string) *FileRepository {
	log.Printf("created FileRepository at path: %s\n", path)
	return &FileRepository{path: path}
}
