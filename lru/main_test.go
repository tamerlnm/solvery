package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	// Test 1: Добавление первых двух элементов
	cache.Set(1, 10)
	cache.Set(2, 20)

	val, ok := cache.Get(1)
	if !ok || val != 10 {
		t.Errorf("Expected Get(1) to return 10, got %v (exists: %v)", val, ok)
	}

	val, ok = cache.Get(2)
	if !ok || val != 20 {
		t.Errorf("Expected Get(2) to return 20, got %v (exists: %v)", val, ok)
	}

	// Test 2: Добавление нового элемента (удаление самого старого (это key 1))
	cache.Set(3, 30)

	_, ok = cache.Get(1)
	if ok {
		t.Error("Expected Get(1) to return false, but it exists")
	}

	val, ok = cache.Get(3)
	if !ok || val != 30 {
		t.Errorf("Expected Get(3) to return 30, got %v (exists: %v)", val, ok)
	}

	// Test 3: Добавление еще одного элемента (удаление второго старого (это key 2))
	cache.Set(4, 40)

	_, ok = cache.Get(2)
	if ok {
		t.Error("Expected Get(2) to return false, but it exists")
	}

	val, ok = cache.Get(4)
	if !ok || val != 40 {
		t.Errorf("Expected Get(4) to return 40, got %v (exists: %v)", val, ok)
	}

	// Test 4: Перемещение элементов при использовании, key 3 будет удален
	cache.Set(5, 50)

	val, ok = cache.Get(4)
	if !ok || val != 40 {
		t.Errorf("Expected Get(4) to return 40, got %v (exists: %v)", val, ok)
	}

	_, ok = cache.Get(3)
	if ok {
		t.Error("Expected Get(3) to return false, but it exists")
	}

	// Test 5: Обновление значения существующего ключа
	cache.Set(5, 55)
	val, ok = cache.Get(5)
	if !ok || val != 55 {
		t.Errorf("Expected Get(5) to return 55 after update, got %v (exists: %v)", val, ok)
	}
}
