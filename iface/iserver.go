package iface

// IServer a server interface
type IServer interface {
	Start()
	Stop()
	Serve()
}
