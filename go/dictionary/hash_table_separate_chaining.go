package main

import (
	"fmt"
)

type Bucket struct {
	Key   string
	Value string
	Next  *Bucket
}

func main() {
	hashTable := [10]Bucket{}

	hashTable = add(hashTable, "a key", "a value")
	newBucket := Bucket{}
	newBucket.Key = "qwer"
	newBucket.Value = "qsdfwer"
	hashTable[8] = newBucket
	hashTable = add(hashTable, "a key", "a valueasdf")

	fmt.Println(hashTable)
}

func createHash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = hash<<5 - hash
		hash = hash + int(key[i])
		hash = hash & hash
	}

	return hash % 10
}

func get(hashTable [10]Bucket, key string) string {
	hash := createHash(key)
	bucket := hashTable[hash]
	for {
		if bucket.Key == key {
			return bucket.Value
		}
		bucket = *bucket.Next
	}
	return ""
}

func add(hashTable [10]Bucket, key, value string) [10]Bucket {
	hash := createHash(key)
	bucket := hashTable[hash]

	newBucket := Bucket{}
	newBucket.Key = key
	newBucket.Value = value

	var b *Bucket
	b = &bucket

	if bucket == (Bucket{}) {
		hashTable[hash] = newBucket
	} else {
		for {
			if b.Key == key {
				b.setBucket(key, value)
				hashTable[hash] = bucket
				break
			}
			if b.Next == nil {
				b.Next = &newBucket
				hashTable[hash] = bucket
				break
			}
			b = b.Next
		}
	}
	return hashTable
}

func (b *Bucket) setBucket(key, value string) {
	b.Value = value
}

func remove(v string) {
}
