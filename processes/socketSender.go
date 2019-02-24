package processes

import (
	"github.com/gorilla/websocket"
	"github.com/ptflp/adrenalinBank/db"
	"time"
)
var Conn websocket.Conn
var Connected = 0
func SocketSend() {
	for {
		operationsJson, _ := db.GetData("operations")
		if operationsJson == "" {
			operationsJson = "[]"
		}
		raw := []byte(operationsJson)
		msgType := 1
		//bytes := []byte(operationsJson)
		if Connected != 0 {
			if err := Conn.WriteMessage(msgType, raw); err != nil {

			}
		}
		//var operationsRaw []models.Operation
		//json.Unmarshal(bytes, &operationsRaw)
		//for _, operation := range operationsRaw {
		//	checkCache, _ := db.GetData(operation.Id)
		//	if checkCache == "" {
		//		db.SaveOperation(operation)
		//		operationJson, err := json.Marshal(operation)
		//		if err != nil {
		//			log.Fatal("Cannot encode to JSON ", err)
		//		}
		//	}
		//}

		time.Sleep(3500 * time.Millisecond)
	}
}