package util

import "strings"
import "strconv"
import "io"
import "os"

func buildStr(str interface{}) string {
	switch str.(type) {
	case nil:
		return "nil"
	case int:
		return strconv.Itoa(int(str.(int)))
	case int32:
		return strconv.FormatInt(int64(str.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(str.(int64)), 10)
	case string:
		return string(str.(string))
	case float64:
		return strconv.FormatFloat(str.(float64), 'f', 6, 64)
	case float32:
		return strconv.FormatFloat(float64(str.(float32)), 'f', 6, 64)
	default:
		break
	}
	return ""
}

func Concat(strs ...interface{}) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(buildStr(str))
	}
	return sb.String()
}

func Println(strs ...interface{}) {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(buildStr(str))
	}
	sb.WriteString("\n")
	io.Copy(os.Stdout, strings.NewReader(sb.String()))
}
