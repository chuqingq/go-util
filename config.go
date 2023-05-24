package util

import (
	"log"
	"os"
)

type Config struct {
	*Message
	filePath        string
	defaultFilePath string
}

func NewConfig(filePath, defaultFilePath string) (*Config, error) {
	m, err := MessageFromFile(filePath)
	if err != nil {
		log.Printf("WARN file %s not found", filePath)
		m, err = MessageFromFile(defaultFilePath)
		if err != nil {
			log.Printf("ERROR file %s not found", defaultFilePath)
			return nil, err
		}
		m.ToFile(filePath)
	}
	return &Config{
		Message:         m,
		filePath:        filePath,
		defaultFilePath: defaultFilePath,
	}, nil
}

// 恢复默认设置
func (c *Config) Reset() error {
	var err error
	c.Message, err = MessageFromFile(c.defaultFilePath)
	if err != nil {
		return err
	}
	c.save()
	return nil
}

// // configSave 模块级配置保存
// func configSave(path string, val interface{}) error {
// 	str := util.ToJson(val)
// 	util.D().Printf("configSave(%v) %v", path, str)
// 	msg, _ := util.MessageFromString(str)
// 	defaultConfig.Set(path, msg.Map("", nil)) // TODO
// 	return nil
// }

func (c *Config) SaveStruct(path string, v interface{}) {
	m := MessageFromStruct(v)
	c.Set(path, m)
}

// // configLoad 模块级配置加载
// func configLoad(path string, val interface{}) {
// 	log.Printf("defaultConfig: %v", defaultConfig)
// 	msg := defaultConfig.Get(path)
// 	msg.Unmarshal(val)
// }

func (c *Config) LoadStruct(path string, v interface{}) {
	m := c.Get(path)
	m.Unmarshal(v)
}

// Set 设置值。Message.Set + save。v支持string/int/bool等，如果是复合值，需要是Map/MustMap()
func (c *Config) Set(path string, v interface{}) {
	c.Message.Set(path, v)
	c.save()
}

// 说明：读取直接使用Message的String/Int/Bool等

func (c *Config) save() error {
	return c.ToFile(c.filePath)
}

func (c *Config) Remove() {
	os.Remove(c.filePath)
}
