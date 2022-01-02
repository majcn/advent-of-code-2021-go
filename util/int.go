package util

import "strconv"

const MaxInt = int(^uint(0) >> 1)

func ParseInt(s interface{}) (rc int) {
	switch s.(type) {
	case byte:
		rc = int(s.(byte) - '0')
	case rune:
		rc = int(s.(rune) - '0')
	case string:
		rc, _ = strconv.Atoi(s.(string))
	case []int:
		tmp := make([]byte, len(s.([]int)))
		for i, v := range s.([]int) {
			tmp[i] = byte('0' + v)
		}
		rc, _ = strconv.Atoi(string(tmp))
	case []byte:
		rc, _ = strconv.Atoi(string(s.([]byte)))
	default:
		rc = -1
	}

	return
}
