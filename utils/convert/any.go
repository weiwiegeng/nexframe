package convert

func Any(value interface{}) *any {
	var a any = value
	// 返回 any 类型变量的地址
	return &a
}