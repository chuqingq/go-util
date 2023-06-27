package util

import (
	"log"

	sjson "github.com/chuqingq/simple-json"
)

// Config 配置
type Config struct {
	*sjson.Json
	ConfigFile        string
	DefaultConfigFile string
}

// NewConfig 从配置文件创建配置
func NewConfig(filePath, defaultFilePath string) (*Config, error) {
	m, err := sjson.FromFile(filePath)
	if err != nil {
		log.Printf("WARN file %s not found", filePath)
		m, err = sjson.FromFile(defaultFilePath)
		if err != nil {
			log.Printf("ERROR file %s not found", defaultFilePath)
			return nil, err
		}
		m.ToFile(filePath)
	}
	return &Config{
		Json:              m,
		ConfigFile:        filePath,
		DefaultConfigFile: defaultFilePath,
	}, nil
}

// Reset 恢复默认设置
func (c *Config) Reset() error {
	var err error
	c.Json, err = sjson.FromFile(c.DefaultConfigFile)
	if err != nil {
		return err
	}
	c.save()
	return nil
}

// LoadStruct 从配置文件加载到结构体
func (c *Config) LoadStruct(path string, v interface{}) {
	m := c.Get(path)
	m.ToStruct(v)
}

// Set 设置值。相当于sjson.Set()+c.save()
func (c *Config) Set(path string, v interface{}) {
	c.Json.Set(path, v)
	c.save()
}

func (c *Config) save() error {
	return c.ToFile(c.ConfigFile)
}
