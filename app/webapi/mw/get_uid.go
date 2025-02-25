package mw

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func GetUid(r *app.RequestContext) uint32 {
	value, _ := r.Get("identity")
	// 定义一个变量来存储最终的 uint32 值
	var uint32Value uint32

	// 根据 value 的实际类型进行转换
	switch v := value.(type) {
	case int:
		// 如果 value 是 int 类型
		uint32Value = uint32(v)
	case int64:
		// 如果 value 是 int64 类型
		uint32Value = uint32(v)
	case string:
		// 如果 value 是 string 类型
		parsedValue, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			fmt.Println("Error parsing string to uint32:", err)
		}
		uint32Value = uint32(parsedValue)
	default:
		// 如果 value 是其他类型
		fmt.Println("Unsupported type:", v)

	}

	return uint32Value

}
