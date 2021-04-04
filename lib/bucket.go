package lib

import (
	"fmt"
)

// insert new node in bucket
func (b *Bucket) Insert(data StoreData) {
	if !b.Search(data.Key) {
		newNode := &BucketNode{Data: data, Next: b.Head}
		b.Head = newNode
		b.Length++
	} else {
		fmt.Println(data.Key, "already exists")
	}
}

//search node in bucket
func (b *Bucket) Search(key string) bool {
	current := b.Head
	for current != nil {
		if current.Data.Key == key {
			return true
		}
		current = current.Next
	}
	return false
}

func (b *Bucket) Get(key string) StoreData {
	if b.Search(key) {
		current := b.Head
		for current != nil {
			if current.Data.Key == key {
				return current.Data
			}
			current = current.Next
		}
	} else {
		fmt.Println(key, "doesn't exists")
	}
	return StoreData{}
}

func (b *Bucket) GetAll() []StoreData {
	data := make([]StoreData, 0)
	current := b.Head
	for current != nil {
		data = append(data, current.Data)
		current = current.Next
	}
	return data
}

//delete
func (b *Bucket) Delete(key string) bool {
	if b.Search(key) {
		if b.Head.Data.Key == key {
			b.Head = b.Head.Next
			b.Length--
			return true
		}

		prev := b.Head

		for prev.Next != nil {
			if prev.Next.Data.Key == key {
				prev.Next = prev.Next.Next
				b.Length--
				return true
			}
			prev = prev.Next
		}
		return false
	} else {
		fmt.Println(key, "doesn't exists")
	}
	return false
}
