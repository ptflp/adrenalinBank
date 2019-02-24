package processes
import (
	"github.com/gorilla/websocket"
	"time"
)
var Conn websocket.Conn
func SocketSend() {
	for {
		msgType := 1
		text := "{\"id\":34,\"value\":100000,\"scaledValue\":9.007333185232472}"
		msg := []byte(text)
		if err := Conn.WriteMessage(msgType, msg); err != nil {
			return
		}

		time.Sleep(500 * time.Millisecond)
	}
}