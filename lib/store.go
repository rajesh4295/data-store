package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
)

// Insert data in store
// Accepts arguement of type StoreData.
// Returns void
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

// Search data in store.
// Accepts key od type string as an argument.
// Returns boolean flag or error.
func (s *Store) Search(key string) (flag bool, err error) {
	idx := Hash(key, StoreSize)
	if s.Slot[idx] != nil {
		return s.Slot[idx].Search(key), nil
	}
	errMsg := fmt.Sprintf("%s doesn't exists", key)
	return false, errors.New(errMsg)
}

// Get value from store using key.
// Accepts key of type string as an argument.
// Returns data of type Storedata or err.
func (s *Store) Get(key string) (data StoreData, err error) {
	idx := Hash(key, StoreSize)
	if s.Slot[idx] != nil {
		return s.Slot[idx].Get(key), nil
	}
	errMsg := fmt.Sprintf("%s doesn't exists", key)
	return StoreData{}, errors.New(errMsg)
}

// Delete data from store.
// Accepts key of type string.
// Returns boolean flag or error
func (s *Store) Delete(key string) (flag bool, err error) {
	idx := Hash(key, StoreSize)
	if s.Slot[idx] != nil {
		del := s.Slot[idx].Delete(key)
		// If bucket at slot gets empty, then empty that slot to reduce the load factor
		if s.Slot[idx].Length == 0 {
			s.Slot[idx] = nil
		}
		return del, nil
	}
	errMsg := fmt.Sprintf("%s doesn't exists", key)
	return false, errors.New(errMsg)
}

// Get current load of store
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

// Export all data in the store to json file.
func (s *Store) Export() error {
	data := make(map[string][]StoreData)
	for i := range s.Slot {
		if s.Slot[i] != nil {
			key := strconv.Itoa(i)
			data[key] = s.Slot[i].GetAll()
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./data.json", j, fs.FileMode(0777))
	if err != nil {
		return err
	}
	return nil
}
