package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xercoy/reaper"
)

func main() {
	gi, err := ghostimport.Import(os.Stdin)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	posts, _ := gi.GetPosts()

	json.NewEncoder(os.Stdout).Encode(posts)
}
