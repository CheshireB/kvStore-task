package sbercloid_test_task

import (
	"sync"
)

type (
	// KVStoreI Be careful if you try to use not comparable object as a key
	// Then panic will acquire
	KVStoreI interface {
		Get(k interface{}) (interface{}, bool)
		Post(k, v interface{}) error
		Put(k, v interface{}) error
		Delete(k interface{}) error
	}

	kvStore struct {
		RWMutex *sync.RWMutex
		store   map[interface{}]interface{}
	}
)

func NewKVStore() KVStoreI {
	return &kvStore{
		RWMutex: &sync.RWMutex{},
		store:   map[interface{}]interface{}{},
	}
}

func (kv *kvStore) Get(k interface{}) (interface{}, bool) {
	kv.RWMutex.RLock()
	defer kv.RWMutex.RUnlock()

	return kv.get(k)
}

func (kv *kvStore) Post(k, v interface{}) error {
	if _, exist := kv.Get(k); exist {
		return NewAlreadyExistError(k)
	}

	kv.RWMutex.Lock()
	kv.storeKV(k, v)
	kv.RWMutex.Unlock()

	return nil
}

func (kv *kvStore) Put(k, v interface{}) error {
	if _, exist := kv.Get(k); !exist {
		return NewKeyNotExistError(k)
	}

	kv.RWMutex.Lock()
	kv.storeKV(k, v)
	kv.RWMutex.Unlock()

	return nil
}

func (kv *kvStore) Delete(k interface{}) error {
	if _, exist := kv.Get(k); !exist {
		return NewKeyNotExistError(k)
	}

	kv.RWMutex.Lock()
	kv.delete(k)
	kv.RWMutex.Unlock()
	return nil
}

func (kv *kvStore) get(k interface{}) (interface{}, bool) {
	v, ok := kv.store[k]
	return v, ok
}

func (kv *kvStore) storeKV(k, v interface{}) {
	kv.store[k] = v
}

func (kv *kvStore) delete(k interface{}) {
	delete(kv.store, k)
}
