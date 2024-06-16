package pokecache

import (
	"testing"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Errorf("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s not found", c.inputKey)
		}

		if string(actual) != string(c.inputVal) {
			t.Errorf("value doesn't match. actual: %s, expected: %s", string(actual), string(c.inputVal))
		}
	}
}
