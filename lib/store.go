package lib

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
	return s.Slot[idx].Search(key)
}

// delete data from store
func (s *Store) Delete(key string) bool {
	idx := Hash(key, StoreSize)
	return s.Slot[idx].Delete(key)
}

// get current load of store
func (s *Store) GetLoad() int {
	// return len(s.Slot)
	load := 0
	for i := range s.Slot {
		if s.Slot[i] != nil {
			load++
		}
	}
	return load
}
