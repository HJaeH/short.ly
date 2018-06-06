package base62

import (
	"fmt"
	"math"
	"strings"
)

const (
	base62elements = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Base           = int64(len(base62elements))
)

func EncodeBase62(n int64) string {
	if n == 0 {
		return string(base62elements[0])
	}

	b := make([]byte, 0)
	for n > 0 {
		r := math.Mod(float64(n), float64(Base))
		n /= Base
		b = append([]byte{base62elements[int(r)]}, b...)
	}
	return string(b)
}

func DecodeBase62(s string) (int64, error) {
	var r int64
	for _, c := range []byte(s) {
		i := strings.IndexByte(base62elements, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		r = Base*r + int64(i)
	}
	return r, nil
}
