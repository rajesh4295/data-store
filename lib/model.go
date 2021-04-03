package lib

const StoreSize = 13

// hash table
type Store struct {
	Slot [StoreSize]*Bucket
}

// linked list
type Bucket struct {
	Head   *BucketNode
	Length int
}

type StoreData struct {
	Key   string
	Value string
}

// linked list node
type BucketNode struct {
	Data StoreData
	Next *BucketNode
}
