package iface

type IGroup interface {
	GET(string, interface{})
	POST(string, interface{})
	PUT(string, interface{})
	PATCH(string, interface{})
	DELETE(string, interface{})
}
