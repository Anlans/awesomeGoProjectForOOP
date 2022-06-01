package model

type userInfo struct {
	Name string
}

// >> TODO 工厂模式
func NewUserInfo(name string) *userInfo {
	return &userInfo{
		Name: name,
	}
}
