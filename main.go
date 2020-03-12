package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"index2/search"
	"io/ioutil"
	"os"
)

func main() {
	var output string
	if len(os.Args) < 2 {
		fmt.Println("To few args")
		return
	}
	if len(os.Args) > 2 {
		output = os.Args[2]
	} else {
		output = "output.yaml"
	}
	if list, err := ioutil.ReadDir(os.Args[1]); err != nil {
		fmt.Println(err)
	} else {
		indexes := make([]search.Index, len(list))
		for i, file := range list {
			bytes, err := ioutil.ReadFile(os.Args[1] + string(os.PathSeparator) + file.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			index := search.GetIndex(file.Name(), string(bytes))
			indexes[i] = index
		}
		inv := search.InvertIndexes(indexes)
		if bytes, e := yaml.Marshal(inv); e == nil {
			if err = ioutil.WriteFile(output, bytes, 0644); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(e)
		}
	}
}

