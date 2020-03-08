package search

import (
	"io/ioutil"
	"os"
	"strings"
)

type SortedStringSet struct {
	i map[string]interface{}
}

type Index struct {
	filename string
	*SortedStringSet
}

type InvertedIndex struct {
	Index map[string]*SortedStringSet
}

func (i *InvertedIndex) Add(key, filename string) {
	if _, exists := i.Index[key]; !exists {
		i.Index[key] = &SortedStringSet{i: map[string]interface{}{}}
	}
	i.Index[key].Add(filename)
}

func (i *InvertedIndex) Json() map[string]interface{} {
	data := make(map[string]interface{})
	for key, value := range i.Index {
		files := make([]string, len(value.i))
		i := 0
		for file, _ := range value.i {
			files[i] = file
			i++
		}
		data[key] = files
	}
	return data
}

func InvertIndexes(indexes []Index) InvertedIndex {
	idx := InvertedIndex{make(map[string]*SortedStringSet)}
	for _, index := range indexes {
		for field, _ := range index.i {
			idx.Add(field, index.filename)
		}
	}
	return idx
}

func (s *SortedStringSet) Add(str string) {
	if _, e := s.i[str]; !e {
		s.i[str] = nil
	}
}

func GetIndexFromFile(path string) (*Index, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	index := GetIndex(path[strings.LastIndex(path, string(os.PathSeparator)) + 1:], string(bytes))
	return &index, nil
}

func GetIndex(filename, str string) Index {
	index := SortedStringSet{i: make(map[string]interface{}, 2)}
	for _, idx := range strings.Fields(str) {
		index.Add(idx)
	}
	return Index{
		filename:        filename,
		SortedStringSet: &index,
	}
}