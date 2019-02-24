package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ptflp/adrenalinBank/api"
	"github.com/ptflp/adrenalinBank/core"
	"github.com/ptflp/adrenalinBank/db"
	"github.com/ptflp/adrenalinBank/models"
	"log"
	"net/http"
	"sort"
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
		var operations []models.Operation
		var users []models.User
		var userModel models.User
		token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiIzIiwibmFtZSI6ImxvZ2luMSIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiMWEyYjE0YjMtOTViMC00MjQwLTg4OTYtMGE2NGE4OWRkZjVjIiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDg5OTc5NiwiZXhwIjoxNTUxNTA0NTk2LCJpYXQiOjE1NTA4OTk3OTYsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.nzPux_UlisWbxxtalV257Y35zLAOnheHokQxJprg0dUkvQzLQTzrGUsjAk3BpoNuwJuL6aJcD5e3rq9Zbi8Khn2o_ZYjrz2jyWehJJ3tPo7IOzWdVhCp5mGz6ducPjRRUTIcGcv3jCeP3T8KTJAv_W4I1xUkS3PcJ1DXWVrbJwRWbEsMpnZS6MiSOluWZ-AV_SBDX0dXZr5y0OnexHNxg9DnHTczHug5xaUlJNQKp_H_S0usGbroHc8bwnSb0NnS_K03WaisSILPw8fUDBqDr8S60tmZodRGqmR6OarbswpgjGmUwmYWP5BsELJFDQI-uWEy0Mg8pHwarcPRLIJJrg"
		productId := "2268927"
		userModel.Token = token
		userModel.ProductId = productId
		users = append(users,userModel)
		userModel.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiI1IiwibmFtZSI6ImxvZ2luMiIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiMTdmMGRlMGUtOThjOC00YWQ5LTk5NzYtYTlkOTE4N2FlZDI4IiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDk1MDkzNCwiZXhwIjoxNTUxNTU1NzM0LCJpYXQiOjE1NTA5NTA5MzQsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.WEEWYw_8jyNtpVpMrN21rdnYnCSKY98C6DElMApcS1-5WfokjTR-1RMSUJLwGLd0-E31k3HqrLbAs6OlH7Q3erWLj5v5m3M7wvvikmt0kd4fsxktiCANrcFiPRSbZ0TGfVn-pmqamO8NtMeAnIB8IgP0W6xXiJ8tAyRbinoYSCSOtsxMPsgPChrqH3p_M9wyiqMO-dJXpjrdqIXzlf_9o9qLQc_Hyq33AwymiRVolqL_GnpwRLhg5gtdrYARyIcEom23pRTrqTDOdHaAbqeKtJIlpR1mE-HxZ2gP3dhgvBK_BYVqbJz0HE1bxRZhKaVsawknjJzE1AJrE0cBe0DbIQ"
		userModel.ProductId = "2268946"
		users = append(users,userModel)
		userModel.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiI2IiwibmFtZSI6ImxvZ2luMyIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiMDJmYjhiMzYtNTBkNC00MDg5LThiYmItNmY3OTdjYzdkOWRlIiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDk1MDk5MiwiZXhwIjoxNTUxNTU1NzkyLCJpYXQiOjE1NTA5NTA5OTIsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.ASG8EGv8y8sIAslh3l7_OaC7_zNV9s51elrSNuBjCfQzLxgGTgGVun_shkqq5ZinlqXI-gwbGWRWdy3sZNtZ0egoqT-vH8MD_TkG1LoSA7-Sz119aDSftZEQWRioxr7QYit3DciZggKIuq9v9pdlVhq1s0RIs3C5uMTx4wN352yHJHEmb779BYxqxqMYRhmgxwNQfFulBgPTOYJyuARyntbMBSYFjDr2tUEDIqWvwR68iUH4R-3PmqQtFz1-vKDLIm2NREA7_UAJTYXl_pYdHtUOeWR4sZjDEYoh5VKfwzVU4atgndXGx1OGEl3n9pQwEFTC3jnt_enQGl44Don1JA"
		userModel.ProductId = "2269093"
		users = append(users,userModel)
		userModel.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiI3IiwibmFtZSI6ImxvZ2luNCIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiNGQzZGM0YTUtMWI1NS00ZDVlLWJmZmYtMDFkNmE2NjE1NjFjIiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDk1MTM0NywiZXhwIjoxNTUxNTU2MTQ3LCJpYXQiOjE1NTA5NTEzNDcsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.RQ_oNuosLw-leWLpm5uWTV9kHXMYDlff8EdCR_Fdx4oLrwN77tXxs-J9d6T6w4V7P_ap6-KijQnzitWFfm5DWytL_sOMXVW9zqvITGr-C-kUYMkiOaTDanvVy7ERaCMidt4nYaT0y_6W5Tbum7Mj-PeMurqZaRJ4yMtw8s6Cy4PPdmv9YzwX_kJbKUnhtlmh0M3HlGoXXrtZAcJGEu3TnRm5ffusJigdPl66duOjpoUDJeAols5MSWw_6E6WjxtwfBhaqgYgislLlLJIpP3F71zW6VBVAOisVmEYRvQaK1YCV8emEfnLCE4cnj4_WlqmPLinyoFE_zec0pbWKLsxWw"
		userModel.ProductId = "2269122"
		users = append(users,userModel)
		userModel.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiI4IiwibmFtZSI6ImxvZ2luNSIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiOTcyZTJkODMtMmUwYy00MjEzLWFmNGQtODFjYjhiYTIwNzAxIiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDk1MTU1MywiZXhwIjoxNTUxNTU2MzUzLCJpYXQiOjE1NTA5NTE1NTMsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.CHJCVDdp5EV7lZ5O0yMLhq1HRUU8r6bUHcfffdMHkqpeSrlHMUF6ob63WSTxrJFgK5Kj6LMhwkhJvHqip039vwyIZPaDUqZJ1V36i7bnPBDrxcbiUA50qjML96G-sVIs7wQZ97eRAcP4IXSA58nItbzYM7eKExXOjl2pHLOG4EvHfPczzXR4_eVrOnWuz2eSe26Xwcthg8qNQieYC6D_63hWn4FrLhiECCDBEIfUbJfZxr0f-rDmBdy_FGLmIUTXfaGZFdDWxZEnTDtl2CKdO6l7g73dR5Y70RmQyJUc47vXbmpxgMynuBUml2NUFeh3B5uIKyOqzAbyru8_c0cYiw"
		userModel.ProductId = "2269141"
		users = append(users,userModel)
		userModel.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiI5IiwibmFtZSI6ImxvZ2luNiIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiZjM4ODAyMDgtNjhhNC00NTA5LWJlOTgtZGZjOWZjOTg1MWY5IiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDk1MTYxMywiZXhwIjoxNTUxNTU2NDEzLCJpYXQiOjE1NTA5NTE2MTMsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.pinhYCxoMi29gu4mU1rVcq_3H7rnAo38Ws0_QbgdY_gEX6TMO8H17xdcuLChjazx_1ScPwkddqHcALhCF9r_1HScWFqwLYe1Ix0rU2DmXIVNIbdaOS7q_WhcrgClZBn20xGJwDNeVzWhew7POwf32jMDL2wZdkLK3tHniZE1RCzYQDMOFdM7Ek4xaOLAMZtuQX0xLFkgNO0pE3z5t9Plip983kG4Uo6cq1LLVIUCRSRP42hXF4PMGchZZXoEJEChVrnQ56F1bhr77eBg1aeAZ2iZJC9-rkSiQoJsdR_9QQYx72H1FOSWXUiWklvx_cvkTDplimn5lXogRAd99KJN1A"
		userModel.ProductId = "2269141"
		users = append(users,userModel)
		userModel.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiI0IiwibmFtZSI6ImxvZ2luNyIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiNzBhNjZkZGQtY2VjOS00MzI1LWIyYjItYzAyZTc4ZWUxM2MyIiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDk1MTgwMSwiZXhwIjoxNTUxNTU2NjAxLCJpYXQiOjE1NTA5NTE4MDEsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.Ze-xEpQR4RjOj2jydPvnqouNAP-PYpMt-Oj5wlIYj7S9crqn99apnobUBHnbSOw7tbJuesMCmdcT5h86X0Mfv9UTKl_PSWCuSIdE9iesUh35rG4M5PTe4hIr6_N4uKp265v8NIH17ItsxlwdBVYqynEwAHDknFIvSxa5QbEH3GHJV-Yu4K71JpLyt4k1VwVyhpPgvLOAk1ftf5ZEZTQnuKMyBZwscIS45bM5aGXtTnHZtTYkPJjEGJIRh5l6Lzd1bo_KxeZneM7UJU6Vr8fozlaFCX3Z2mgAI-AZVKs4tToRxQfvk4c3kYhAB-7Xg94q7h0RjoaAZ1VHEPhTUod1FQ"
		userModel.ProductId = "2269212"
		users = append(users,userModel)
		for _, user := range users {
			operationClients := core.Operations(user)
			for _, operation := range operationClients {
				operations = append(operations, operation)
			}
		}
		//operations = core.SortOperationsAsc(operations)
		sort.Slice(operations, func(i, j int) bool {
			return operations[i].Timestamp < operations[j].Timestamp
		})
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
    go core.Socket()
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/web/:method", Web)
	router.POST("/api/:method", Api)

	log.Fatal(http.ListenAndServe(":8089", router))

}