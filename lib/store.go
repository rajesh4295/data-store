package lib

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
)

// insert data in store
func (s *Store) Insert(data StoreData) {
	idx := Hash(data.Key, StoreSize)
	// if slot is empty, then first init the bucket and then inster data
	if s.Slot[idx] == nil {
		s.Slot[idx] = &Bucket{Head: nil, Length: 0}
		s.Slot[idx].Insert(data)
	} else {
		s.Slot[idx].Insert(data)
	}
}

// search data in store
func (s *Store) Search(key string) bool {
	idx := Hash(key, StoreSize)
	if s.Slot[idx] != nil {
		return s.Slot[idx].Search(key)
	}
	fmt.Println(key, "doesn't exists")
	return false
}

func (s *Store) Get(key string) StoreData {
	idx := Hash(key, StoreSize)
	if s.Slot[idx] != nil {
		return s.Slot[idx].Get(key)
	}
	fmt.Println(key, "doesn't exists")
	return StoreData{}
}

// delete data from store
func (s *Store) Delete(key string) bool {
	idx := Hash(key, StoreSize)
	if s.Slot[idx] != nil {
		del := s.Slot[idx].Delete(key)
		// If bucket at slot gets empty, then empty that slot to reduce the load factor
		if s.Slot[idx].Length == 0 {
			s.Slot[idx] = nil
		}
		return del
	}
	fmt.Println(key, "doesn't exists")
	return false
}

// get current load of store
// TODO: If load goes above 80% then double the size of store
func (s *Store) GetLoad() int {
	// return len(s.Slot)
	load := 0.0
	for i := range s.Slot {
		if s.Slot[i] != nil {
			load++
		}
	}
	load = float64(load) / float64(StoreSize) * 100
	return int(load)
}

func (s *Store) Export() {
	data := make(map[string][]StoreData)
	for i := range s.Slot {
		if s.Slot[i] != nil {
			key := strconv.Itoa(i)
			data[key] = s.Slot[i].GetAll()
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("could not marshal data", err)
	}
	err = ioutil.WriteFile("./data.json", j, fs.FileMode(0777))
	if err != nil {
		fmt.Println("could not write data to file", err)
	}
}
