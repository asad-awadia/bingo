package codecs

import (
	"time"
)

func EncodeValue(b []byte, v interface{}, inverse bool) (int, error) {
	if v == nil {
		return EncodeNil(b, inverse), nil
	}
	switch c := v.(type) {
	case int8:
		return EncodeInt8(b, c, inverse)
	case int16:
		return EncodeInt16(b, c, inverse), nil
	case int:
		return EncodeInt32(b, int32(c), inverse), nil
	case int32:
		return EncodeInt32(b, c, inverse), nil
	case int64:
		return EncodeInt64(b, c, inverse), nil
	case float32:
		return EncodeFloat32(b, c, inverse), nil
	case float64:
		return EncodeFloat64(b, c, inverse), nil
	case string:
		return EncodeString(b, c, inverse), nil
	case time.Time:
		return EncodeTime(b, c, inverse), nil
	default:
		return 0, ErrUnknownType
	}
}
