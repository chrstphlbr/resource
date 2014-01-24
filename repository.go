package ressource

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type Repository interface {
	Update()
	Ressources() []Adapter
}

type FileRepository struct {
	sync.Mutex
	ressources []Adapter
	path       string
}

func (repo *FileRepository) Update() {
	repo.Lock()
	defer repo.Unlock()
	log.Println("start update")
	fileDescriptors, err := ioutil.ReadDir(repo.path)
	if err != nil {
		log.Printf("error retrieving file descriptors of directory (%s)\n", repo.path)
		log.Println(err)
		return
	}
	log.Printf("directories has %d files", len(fileDescriptors))
	// dirty implementation: remove all previous ressources and add all again
	repo.ressources = make([]Adapter, 0, len(fileDescriptors))
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
		// add RessourceAsapter to ressources
		repo.ressources = append(repo.ressources, NewFileAdapter(pathToFile))
	}
	log.Println("finished update")
}

func (repo FileRepository) Ressources() []Adapter {
	repo.Lock()
	defer repo.Unlock()
	return repo.ressources
}

func NewFileRepository(path string) *FileRepository {
	log.Printf("created FileRepository at path: %s\n", path)
	return &FileRepository{path: path}
}
