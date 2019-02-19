package pie

//import "fmt"
import "sort"

var Pair = map[string]string{}

func FindKeys(jsonObj interface{}, keyId, keyName string) {
	FindValue(jsonObj, keyId, keyName)
}
func FindArray(a []interface{}, keyId, keyName string) {
	if len(a) == 0 {
		return
	}

	for _, b := range a {
		FindValue(b, keyId, keyName)
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
			Pair[key] = m[key].(string)
		}
		FindValue(m[key], keyId, keyName)
	}
}
