package lib

func InitStore() *Store {
	store := &Store{}

	// for i := range store.Slot {
	// 	store.Slot[i] = &Bucket{Head: nil, Length: 0}
	// }
	return store
}

func GetStore() *Store {
	return &Store{}
}

func Hash(key string, storeLen int) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % storeLen
}
