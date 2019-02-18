package pie

import "bytes"
import "fmt"
import "encoding/json"
import "strings"
import "sort"
import "strconv"

type Pie struct {
	Data interface{}
}

func MarshalString(str string, buf *bytes.Buffer) {
	strBytes, _ := json.Marshal(str)
	str = string(strBytes)

	StringMaxLength := 255
	if len(str) >= StringMaxLength {
		str = fmt.Sprintf("%s...", str[0:StringMaxLength])
	}

	buf.WriteString(str)
}

func Marshal(jsonObj interface{}) ([]byte, error) {
	buffer := bytes.Buffer{}
	initialDepth := 0
	MarshalValue(jsonObj, &buffer, initialDepth)
	return buffer.Bytes(), nil
}

func WriteIndent(buf *bytes.Buffer, depth int) {
	buf.WriteString(strings.Repeat(" ", 2*depth))
}

func MarshalMap(m map[string]interface{}, buf *bytes.Buffer, depth int) {
	remaining := len(m)

	if remaining == 0 {
		buf.WriteString("{}")
		return
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	buf.WriteString("{")
	buf.WriteByte('\n')

	for _, key := range keys {
		WriteIndent(buf, depth+1)
		buf.WriteString(fmt.Sprintf("\"%s\": ", key))
		MarshalValue(m[key], buf, depth+1)
		remaining--
		if remaining != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte('\n')
	}
	WriteIndent(buf, depth)
	buf.WriteString("}")
}

func MarshalArray(a []interface{}, buf *bytes.Buffer, depth int) {
	if len(a) == 0 {
		buf.WriteString("[]")
		return
	}

	buf.WriteString("[")
	buf.WriteByte('\n')

	for i, v := range a {
		WriteIndent(buf, depth+1)
		MarshalValue(v, buf, depth+1)
		if i < len(a)-1 {
			buf.WriteString(",")
		}
		buf.WriteByte('\n')
	}
	WriteIndent(buf, depth)
	buf.WriteString("]")
}

func MarshalValue(val interface{}, buf *bytes.Buffer, depth int) {
	switch v := val.(type) {
	case map[string]interface{}:
		MarshalMap(v, buf, depth)
	case []interface{}:
		MarshalArray(v, buf, depth)
	case string:
		MarshalString(v, buf)
	case float64:
		buf.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		buf.WriteString(strconv.FormatBool(v))
	case nil:
		buf.WriteString("null")
	}
}

// credit to https://github.com/TylerBrock/
