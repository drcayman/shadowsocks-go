package ss

import "sync"

var services sync.Map

func bultinServiceHandler(conn Conn, lis *listener) (c Conn) {
	var dst string
	if conn.GetDst() != nil {
		dst = conn.GetDst().String()
	}
	if len(dst) == 0 {
		c = conn
		return
	}
	v, ok := services.Load(dst)
	if !ok {
		c = conn
		return
	}
	handler := v.(listenHandler)
	c = handler(conn, lis)
	return
}

// StoreServiceHandler stores the handler to services map with key addr
func StoreServiceHandler(addr string, handler listenHandler) {
	services.Store(addr, handler)
}

func init() {
	StoreServiceHandler(echoaddr, echoHandler)
}

func echoHandler(conn Conn, lis *listener) (c Conn) {
	go Pipe(conn, conn, lis.c)
	c = nilConn
	return
}
