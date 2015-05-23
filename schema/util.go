package schema

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func String(target interface{}, key string) string {
	d, ok := target.(map[string]interface{})
	if !ok {
		return ""
	}

	v, ok := d[key]
	if !ok {
		return ""
	}

	var res string

	switch v.(type) {
	case string:
		res = v.(string)
	case float64:
		res = fmt.Sprintf("%v", v)
	case []interface{}:
		rs := v.([]interface{})
		for _, r := range rs {
			if vv, ok := r.(string); ok {
				res = res + vv
			}
		}
	}
	return res
}

func Interface(target interface{}, key string, typeStr string) interface{} {
	if typeStr == "string" {
		return String(target, key)
	}

	var i interface{}
	d, ok := target.(map[string]interface{})
	if !ok {
		return i
	}

	v, ok := d[key]
	if !ok {
		return String(target, key)
	}
	return v
}

func isSupportExt(s string) bool {
	return isJSONExt(s) || isYAMLExt(s)
}

func isJSONExt(s string) bool {
	ext := path.Ext(s)
	return ext == ".json"
}

func isYAMLExt(s string) bool {
	ext := path.Ext(s)
	return ext == ".yaml" || ext == ".yml"
}

func baseResourceName(s string) string {
	ss := strings.Split(s, "/")
	name := ss[len(ss)-1]

	ext := path.Ext(name)
	return name[0 : len(name)-len(ext)]
}

func isExtJSONFile(info os.FileInfo) bool {
	return isFile(info) && isMatchExt(info, ".json")
}

func isExtYaml(info os.FileInfo) bool {
	return isFile(info) && isMatchExt(info, ".yml", ".yaml")
}

func isMatchExt(info os.FileInfo, exts ...string) bool {
	e := filepath.Ext(info.Name())
	for _, ext := range exts {
		if e == ext {
			return true
		}
	}
	return false
}

func isFile(info os.FileInfo) bool {
	if info == nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}

func resourceName(name string) string {
	base := path.Base(name)
	e := filepath.Ext(base)
	return base[0 : len(base)-len(e)]
}
