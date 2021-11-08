package storage

import (
	"strconv"
	"time"
)

type Store struct {
	m map[string]string
}

func NewStore() *Store {
	return &Store{map[string]string{}}
}

func (store *Store) Save(content string) string {
	ch := make(chan string)
	go func() {
		key := genKey()
		store.m[key] = content
		ch <- key
	}()
	return <-ch
}

func (store *Store) Get(key string) string {
	return store.m[key]
}

func genKey() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
