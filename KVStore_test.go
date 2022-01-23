package sbercloid_test_task

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKvStore(t *testing.T) {
	t.Run("Happy path. Post", func(t *testing.T) {
		expKey := "exp_key"
		expValue := "exp_value"
		kv := NewKVStore()

		err := kv.Post(expKey, expValue)
		assert.NoError(t, err)

		actValue, exist := kv.Get(expKey)
		assert.True(t, exist)
		assert.EqualValues(t, expValue, actValue)
	})

	t.Run("Happy path. Put", func(t *testing.T) {
		expKey := "exp_key"
		expValue := "exp_value"
		newExpValue := "new_exp_value"
		kv := NewKVStore()

		err := kv.Post(expKey, expValue)
		assert.NoError(t, err)

		err = kv.Put(expKey, newExpValue)
		assert.NoError(t, err)

		actValue, exist := kv.Get(expKey)
		assert.True(t, exist)
		assert.EqualValues(t, newExpValue, actValue)
	})

	t.Run("Happy path. Delete", func(t *testing.T) {
		expKey := "exp_key"
		expValue := "exp_value"
		kv := NewKVStore()

		err := kv.Post(expKey, expValue)
		assert.NoError(t, err)

		err = kv.Delete(expKey)
		assert.NoError(t, err)

		actValue, exist := kv.Get(expKey)
		assert.False(t, exist)
		assert.Nil(t, actValue)
	})
}

func TestKvStore_Get(t *testing.T) {
	t.Run("Key not exist", func(t *testing.T) {
		expKey := "exp_key"
		kv := NewKVStore()

		_, exist := kv.Get(expKey)
		assert.False(t, exist)
	})

	t.Run("Concurrency access to map", func(t *testing.T) {
		kv := NewKVStore()

		for i := 0; i < 100; i++ {
			err := kv.Post(i, i)
			assert.NoError(t, err)
		}

		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				v, err := kv.Get(i)
				assert.True(t, err)
				assert.Equal(t, i, v)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

func TestKvStore_Post(t *testing.T) {
	t.Run("Key already exist", func(t *testing.T) {
		expKey := "exp_key"
		expValue := "exp_value"
		kv := NewKVStore()

		err := kv.Post(expKey, expValue)
		assert.NoError(t, err)

		err = kv.Post(expKey, expValue)
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf(keyAlreadyExistErrorTemplate, expKey), err.Error())
	})

	t.Run("Concurrency access to map", func(t *testing.T) {
		kv := NewKVStore()

		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				err := kv.Post(i, i)
				assert.NoError(t, err)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

func TestKvStore_Put(t *testing.T) {
	t.Run("Key not exist", func(t *testing.T) {
		expKey := "exp_key"
		expValue := "exp_value"
		kv := NewKVStore()

		err := kv.Put(expKey, expValue)
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf(keyNotExistErrorTemplate, expKey), err.Error())
	})

	t.Run("Concurrency access to map", func(t *testing.T) {
		kv := NewKVStore()

		for i := 0; i < 100; i++ {
			err := kv.Post(i, i)
			assert.NoError(t, err)
		}

		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				err := kv.Put(i, i)
				assert.NoError(t, err)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

func TestKvStore_Del(t *testing.T) {
	t.Run("Key not exist", func(t *testing.T) {
		expKey := "exp_key"
		kv := NewKVStore()

		err := kv.Delete(expKey)
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf(keyNotExistErrorTemplate, expKey), err.Error())
	})

	t.Run("Concurrency access to map", func(t *testing.T) {
		kv := NewKVStore()

		for i := 0; i < 100; i++ {
			err := kv.Post(i, i)
			assert.NoError(t, err)
		}

		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				err := kv.Delete(i)
				assert.NoError(t, err)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}
