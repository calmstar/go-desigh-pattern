package singleton

// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

// 利用代码init完成单例
func init() {
	singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}
