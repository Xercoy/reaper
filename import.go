package ghostimport

import (
	"encoding/json"
	"io"
	"log"
)

func Import(r io.Reader) (GhostImport, error) {
	var gi GhostImport

	err := json.NewDecoder(r).Decode(&gi)
	if err != nil {
		log.Println("Error importing.")
		return GhostImport{}, err
	}

	return gi, nil
}

/*func (gi *GhostImport) GetPosts {

}*/
