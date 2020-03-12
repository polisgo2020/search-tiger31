package search

import (
	"strings"
)

type SortedStringSet struct {
	i map[string]interface{}
}

type Index struct {
	key string
	*SortedStringSet
}

type InvertedIndex map[string]*SortedStringSet

func (i InvertedIndex) Add(key, filename string) {
	if _, exists := i[key]; !exists {
		i[key] = &SortedStringSet{i: map[string]interface{}{}}
	}
	i[key].Add(filename)
}

func (i InvertedIndex) MarshalYAML() (interface{}, error)  {
	data := make(map[string]interface{})
	for key, value := range i {
		files := make([]string, len(value.i))
		i := 0
		for file := range value.i {
			files[i] = file
			i++
		}
		data[key] = files
	}
	return data, nil
}

func InvertIndexes(indexes []Index) InvertedIndex {
	idx := InvertedIndex{}
	for _, index := range indexes {
		for field, _ := range index.i {
			idx.Add(field, index.key)
		}
	}
	return idx
}

func (s *SortedStringSet) Add(str string) {
	if _, e := s.i[str]; !e {
		s.i[str] = nil
	}
}

func GetIndex(key, str string) Index {
	index := SortedStringSet{i: make(map[string]interface{})}
	for _, idx := range strings.Fields(str) {
		index.Add(idx)
	}
	return Index{
		key:             key,
		SortedStringSet: &index,
	}
}