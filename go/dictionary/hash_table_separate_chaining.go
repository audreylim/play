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
	fmt.Println(get(hashTable, "qwer"))
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
		if bucket == (Bucket{}) {
			break
		}
		if bucket.Key == key {
			return bucket.Value
		}
		if *bucket.Next != (Bucket{}) {
			bucket = *bucket.Next
		} else {
			break
		}
	}
	return ""
}

func add(hashTable [10]Bucket, key, value string) [10]Bucket {
	hash := createHash(key)
	bucket := hashTable[hash]

	newBucket := Bucket{}
	newBucket.Key = key
	newBucket.Value = value

	if bucket == (Bucket{}) {
		hashTable[hash] = newBucket
	} else {
		for {
			// Override value of existing key
			if bucket.Key == key {
				bucket.Value = value
				hashTable[hash] = bucket
				break
			}
			// Append to tail of linked list
			if bucket.Next == nil {
				bucket.Next = &newBucket
				hashTable[hash] = bucket
				break
			}
			bucket = *bucket.Next
		}
	}
	return hashTable
}

func remove(v string) {
}
