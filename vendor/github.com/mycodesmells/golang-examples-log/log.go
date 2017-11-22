package log

import (
	"fmt"
	"time"
)

func Println(v ...interface{}) {
	v = append(v, 0)
	copy(v[1:], v[:])
	v[0] = time.Now().Format(time.RFC3339)
	fmt.Println(v...)
}
