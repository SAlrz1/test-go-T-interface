package main

import (
	"errors"
	"fmt"
	"strconv" // 引入strconv包，用于正确的数字/字符串转换
)

// 泛型转换函数：正确实现类型转换+泛型返回
func ConvertTo[T any](data interface{}) (T, error) {
	// 声明T类型的返回值（初始为零值）
	var result T

	switch tp := data.(type) {
	case int:
		// 正确：int转string（123→"123"）
		strVal := strconv.Itoa(tp)
		// 类型断言：将string转成泛型T（需确保T是string，否则断言失败）
		if v, ok := any(strVal).(T); ok {
			result = v
		} else {
			return result, errors.New("目标类型不是string，无法转换")
		}
	case string:
		return result, errors.New("不支持string类型转换")
	case float64:
		// float64转int
		intVal := int(tp)
		// 类型断言：将int转成泛型T（需确保T是int，否则断言失败）
		if v, ok := any(intVal).(T); ok {
			result = v
		} else {
			return result, errors.New("目标类型不是int，无法转换")
		}
	default:
		return result, errors.New("不支持的原始类型")
	}
	return result, nil
}

func main() {
	// 测试1：int(123) → string
	m, err := ConvertTo[string](123)
	fmt.Println("m:", m, "err:", err) // 输出：m: 123 err: <nil>

	// 测试2：string("456") → int（返回错误）
	m1, err1 := ConvertTo[int]("456")
	fmt.Println("m1:", m1, "err1:", err1) // 输出：m1: 0 err1: 不支持string类型转换

	// 测试3：float64(3.14) → int
	m2, err2 := ConvertTo[int](3.14)
	fmt.Println("m2:", m2, "err2:", err2) // 输出：m2: 3 err2: <nil>

	// 测试4：错误场景（int→int，原始类型不匹配）
	m3, err3 := ConvertTo[int](123)
	fmt.Println("m3:", m3, "err3:", err3) // 输出：m3: 0 err3: 不支持的原始类型
}
