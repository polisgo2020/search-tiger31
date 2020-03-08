package main

import (
	"encoding/json"
	"fmt"
	"index2/search"
	"io/ioutil"
	"os"
)

func main() {
	var output string
	if len(os.Args) > 2 {
		output = os.Args[2]
	} else {
		output = "output.json"
	}
	if list, err := ioutil.ReadDir(os.Args[1]); err != nil {
		fmt.Println(err)
	} else {
		indexes := make([]search.Index, len(list))
		for i, file := range list {
			index, e := search.GetIndexFromFile(os.Args[1] + string(os.PathSeparator) + file.Name())
			if e != nil {
				err = e
				break
			}
			indexes[i] = *index
		}
		if err != nil {
			fmt.Println(err)
		} else {
			inv := search.InvertIndexes(indexes)
			if bytes, e := json.Marshal(inv.Json()); e == nil {
				if err = ioutil.WriteFile(output, bytes, 0644); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(e)
			}
		}
	}
}

