package pie

//import "fmt"
import "sort"

var Pair = []map[string]string{}

func FindKeys(jsonObj interface{}, keyId, keyName string) {
	Pair = []map[string]string{}
	FindValue(0, jsonObj, keyId, keyName)
}

func FindArray(a []interface{}, keyId, keyName string) {
	if len(a) == 0 {
		return
	}

	for i, b := range a {
		FindValue(i, b, keyId, keyName)
	}
}

func FindValue(i int, val interface{}, keyId, keyName string) {
	switch v := val.(type) {
	case map[string]interface{}:
		FindMap(i, v, keyId, keyName)
	case []interface{}:
		FindArray(v, keyId, keyName)
	}
}

func FindMap(i int, m map[string]interface{}, keyId, keyName string) {
	if len(m) == 0 {
		return
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	t := map[string]string{}
	for _, key := range keys {
		if key == keyId {
			t["id"] = m[key].(string)
		}
		if key == keyName {
			t["name"] = m[key].(string)
		}
		FindValue(i, m[key], keyId, keyName)
	}
	if len(t) > 0 {
		Pair = append(Pair, t)
	}
}
