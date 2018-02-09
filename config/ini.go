package config

import (
	"strconv"
	"strings"
	"bufio"
	"io"
	"os"
)

type IniConfig struct {
	configMap map[string]string
	strcet    string
}

func (c *IniConfig) Load(filePtah string) {
	c.configMap = make(map[string]string)

	f, err := os.Open(filePtah)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1: n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + "&&" + frist
		c.configMap[key] = strings.TrimSpace(second)
	}
}

func (c IniConfig) Read(node, key string) string {
	key = node + "&&" + key
	v, found := c.configMap[key]
	if !found {
		return ""
	}
	return v
}

func (c IniConfig) ReadString(node, key, defaultValue string) string {
	key = node + "&&" + key
	v, found := c.configMap[key]
	if !found {
		return defaultValue
	}
	return v
}

func (c IniConfig) ReadInt(node, key string, defaultValue int) int {
	key = node + "&&" + key
	v, found := c.configMap[key]
	if !found {
		return defaultValue
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	} else {
		return value
	}
}

func (c IniConfig) ReadInt64(node, key string, defaultValue int64) int64 {
	key = node + "&&" + key
	v, found := c.configMap[key]
	if !found {
		return defaultValue
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	} else {
		return int64(value)
	}
}

func (c IniConfig) ReadBool(node, key string, defaultValue bool) bool {
	key = node + "&&" + key
	v, found := c.configMap[key]
	if !found {
		return defaultValue
	}
	if strings.ToLower(v) == "true" {
		return true
	}
	if strings.ToLower(v) == "false" {
		return false
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	} else {
		return value != 0
	}
}
