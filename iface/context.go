package iface

// IContext 重写gin.context
type IContext interface {
	Token(...string) string
}
