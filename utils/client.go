package utils

// 设置客户端内容
// TODO client包
// 设置语言 视距 接受的Chat类型 Connection对象
// 构造方式：传入：语言 视距 接受的Chat类型

// Connection对象为nil
// 后期使用Bot提供的服务器ip以及Account创建Connection对象

type Settings struct {
	Locale       string // 地区
	ViewDistance int    // 视距
	ChatMode     int    // 聊天模式
	ChatColors   bool   // 聊天颜色
	// TODO Habit
}
