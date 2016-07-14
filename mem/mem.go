package mem

import (
	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheClient struct {
	Servers []string `json:"servers"`
	client  *memcache.Client
}

func New(config string) (m *MemcacheClient, err error) {
	m = &MemcacheClient{}
	err = json.Unmarshal([]byte(config), &m)
	if err != nil {
		return
	}
	m.client = memcache.New(m.Servers...)
	return
}

func (c *MemcacheClient) Get(key string) string {
	data, err := c.client.Get(key)
	if err != nil {
		return ""
	}
	return string(data.Value)
}

func (c *MemcacheClient) Add(key string, value string, expiresAt int32) error {
	data := &memcache.Item{Key: key, Value: []byte(value), Expiration: expiresAt}
	return c.client.Add(data)
}

func (c *MemcacheClient) Set(key string, value string, expiresAt int32) error {
	data := &memcache.Item{Key: key, Value: []byte(value), Expiration: expiresAt}
	return c.client.Set(data)
}

func (c *MemcacheClient) Delete(key string) error {
	return c.client.Delete(key)
}

func (c *MemcacheClient) Delay(key string, expiresAt int32) error {
	v := c.Get(key)
	data := &memcache.Item{Key: key, Value: []byte(v), Expiration: expiresAt}
	return c.client.Set(data)
}
