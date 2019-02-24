package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ptflp/adrenalinBank/api"
	"github.com/ptflp/adrenalinBank/db"
	"github.com/ptflp/adrenalinBank/processes"
	"log"
	"net/http"
)

type ErrStruct struct {
	Error string `json:"error"`
}

type operation struct {
	Id                  string
	Date                string
	Timestamp           int64
	Amount              float64
	Currency            string
	Details             string
	CounterpartName     string
	CounterpartBankName string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {

	db.InitRedis()

	api.StartGame()
	go processes.Socket()
	go processes.SocketSend()
	go processes.SyncData()
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8089", router))

}
