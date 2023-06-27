package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const storeFile = "test.store"
const defaultStoreFile = "default-base.store"

func create() {
	os.WriteFile(storeFile, []byte(`{}`), 0644)
	os.WriteFile(defaultStoreFile, []byte(`{"col1":{"key1":"value1"}}`), 0644)
}

func clean() {
	os.Remove(storeFile)
	os.Remove(defaultStoreFile)
}

func TestConfigDefault(t *testing.T) {
	defer clean()
	os.WriteFile(defaultStoreFile, []byte(`{"col1":{"key1":"value1"}}`), 0644)

	// defaultStoreFile不存在，也可以创建config，会从defaultStoreFile创建
	conf, err := NewConfig(storeFile, defaultStoreFile)
	assert.Nil(t, err)

	value := conf.Get("col1.key1").MustString()
	assert.Equal(t, "value1", value)
}

func TestConfigReset(t *testing.T) {
	defer clean()
	os.WriteFile(storeFile, []byte(`{}`), 0644)
	os.WriteFile(defaultStoreFile, []byte(`{"col1":{"key1":"value1"}}`), 0644)

	conf, err := NewConfig(storeFile, defaultStoreFile)
	assert.Nil(t, err)

	value := conf.Get("col1.key1").MustString()
	assert.Equal(t, "", value)

	err = conf.Reset()
	assert.Nil(t, err)

	value = conf.Get("col1.key1").MustString("not exist")
	assert.Equal(t, "value1", value)
}

func TestConfigSetGet(t *testing.T) {
	defer clean()
	create()

	conf, err := NewConfig(storeFile, defaultStoreFile)
	assert.Nil(t, err)

	value := conf.Get("col1.key1").MustString()
	assert.Equal(t, "", value)

	conf.Set("col1.key1", "value1")
	value = conf.Get("col1.key1").MustString("not exist")
	assert.Equal(t, "value1", value)

	// update key
	conf.Set("col1.key1", "value2")
	value = conf.Get("col1.key1").MustString("not exist")
	assert.Equal(t, "value2", value)
}
