package dto

// RegisterForm 注册表单验证
type RegisterForm struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=18"`
}
