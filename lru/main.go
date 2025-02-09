package main

import (
	"fmt"
	"lru/pkg/model"
)

func main() {
	cache := model.NewLRUCache(2)

	cache.Set(1, 10)
	cache.Set(2, 20)

	if val, ok := cache.Get(1); ok {
		fmt.Println("Key 1:", val)
	}

	cache.Set(3, 30)

	if _, ok := cache.Get(2); !ok {
		fmt.Println("Key 2 not found")
	}

	cache.Set(4, 40)

	if _, ok := cache.Get(1); !ok {
		fmt.Println("Key 1 not found")
	}

	if val, ok := cache.Get(3); ok {
		fmt.Println("Key 3:", val)
	}

	if val, ok := cache.Get(4); ok {
		fmt.Println("Key 4:", val)
	}
}
