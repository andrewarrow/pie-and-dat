package pie

//import "fmt"
import "sort"

func FindKeys(jsonObj interface{}, keyId, keyName string) map[string]interface{} {
	return FindValue(jsonObj, keyId, keyName)
}
func FindArray(a []interface{}, keyId, keyName string) map[string]interface{} {
	v := map[string]interface{}{}
	if len(a) == 0 {
		return v
	}

	for _, v := range a {
		v = FindValue(v, keyId, keyName)
	}
	return v
}
func FindValue(val interface{}, keyId, keyName string) map[string]interface{} {
	switch v := val.(type) {
	case map[string]interface{}:
		return FindMap(v, keyId, keyName)
	case []interface{}:
		return FindArray(v, keyId, keyName)
	}
	return map[string]interface{}{}
}

func FindMap(m map[string]interface{}, keyId, keyName string) map[string]interface{} {
	v := map[string]interface{}{}
	if len(m) == 0 {
		return v
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key == keyId || key == keyName {
			v[key] = m[key]
		}
		FindValue(m[key], keyId, keyName)
	}
	return v
}
