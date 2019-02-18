package pie

import "fmt"
import "sort"

func FindKeys(jsonObj interface{}, keyId, keyName string) []string {
	FindValue(jsonObj, keyId, keyName)
	pair := []string{"", ""}
	return pair
}
func FindArray(a []interface{}, keyId, keyName string) {
	if len(a) == 0 {
		return
	}

	for _, v := range a {
		FindValue(v, keyId, keyName)
	}
}
func FindValue(val interface{}, keyId, keyName string) {
	switch v := val.(type) {
	case map[string]interface{}:
		FindMap(v, keyId, keyName)
	case []interface{}:
		FindArray(v, keyId, keyName)
	}
}

func FindMap(m map[string]interface{}, keyId, keyName string) {
	if len(m) == 0 {
		return
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key == keyId || key == keyName {
			fmt.Println(key, m[key])
		}
		//buf.WriteString(fmt.Sprintf("\"%s\": ", key))
		FindValue(m[key], keyId, keyName)
	}
}
