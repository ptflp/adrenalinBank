package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ptflp/adrenalinBank/api"
	"github.com/ptflp/adrenalinBank/core"
	"github.com/ptflp/adrenalinBank/db"
	"github.com/ptflp/adrenalinBank/models"
	"github.com/julienschmidt/httprouter"
	"sort"
	//"strconv"

	//"os"

	//"io/ioutil"
	"log"
	"net/http"
	"time"
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

func Api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	p := r.Header.Get("user")
	var ap uint8

	if p == "white" {
		ap = 0
	} else if p == "black" {
		ap = 1
	} else {
		fmt.Fprint(w, errors.New("Игрок не найден"))
		return
	}

	if ps.ByName("method") == "move" {
		decoderBody := json.NewDecoder(r.Body)
		var m models.ClientMoveObject
		err := decoderBody.Decode(&m)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		status, err := api.Move(ap, m)
		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			if status == "White checkmated Black" || status == "Black checkmated White" || status == "Draw" {
				time.Sleep(3 * time.Second)
				db.DeleteTestMoves()
			}
		}

	}
}

func Web(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch p.ByName("method") {
	case "operation":
		operations := core.chetoTam()
		operationsJson, err := json.Marshal(operations)
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}
		fmt.Fprintf(w, "%s", operationsJson)
	default:
		fmt.Fprintf(w, "404 error")
	}
	//if p.ByName("method") == "getmoves" {
	//	moves, err := api.GetMovesResponse("test")
	//	if err != nil {
	//		//fmt.Fprint(w, err.Error())
	//	} else {
	//		//fmt.Fprint(w, moves)
	//	}
	//} else if p.ByName("method") == "flushall" {
	//	db.DeleteTestMoves()
	//}
}

func main() {

	db.InitRedis()

	api.StartGame()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/web/:method", Web)
	router.POST("/api/:method", Api)

	log.Fatal(http.ListenAndServe(":8089", router))
}
