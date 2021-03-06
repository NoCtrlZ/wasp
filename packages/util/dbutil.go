package util

import (
	"fmt"
	"github.com/iotaledger/hive.go/kvstore"
)

func DbSetMulti(store kvstore.KVStore, keys [][]byte, values [][]byte) error {
	if len(keys) != len(values) {
		return fmt.Errorf("number of keys muts be equal to number of values")
	}
	atomic := store.Batched()
	for i := range keys {
		if err := atomic.Set(keys[i], values[i]); err != nil {
			return err
		}
	}
	return atomic.Commit()
}

func DbGetMulti(store kvstore.KVStore, keys [][]byte) ([][]byte, error) {
	ret := make([][]byte, len(keys))
	var err error
	for i, k := range keys {
		if ret[i], err = store.Get(k); err != nil {
			return nil, err
		}
	}
	return ret, nil
}
