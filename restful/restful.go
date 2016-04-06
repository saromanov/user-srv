package restful

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"github.com/saromanov/user-srv/proto/account"

	"golang.org/x/net/context"
)

type User struct{}

var (
	cl account.AccountClient
)

func Update(w http.ResponseWriter, r *http.Request) {
	var req account.UpdateNameRequest
	var err error
	//Trying to decode request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//var resp account.UpdateNameResponse
	_, err = cl.UpdateName(context.TODO(), &account.UpdateNameRequest{
		Email: req.Email,
		Oldusername:req.Oldusername,
		Username:req.Username,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var usr account.User
	var err error

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	req := client.NewRequest("go.micro.srv.user", "Account.Create", &usr)

	rsp := &account.CreateResponse{}
	err = client.Call(context.Background(), req, rsp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	result, err := json.Marshal(rsp)
	if err != nil {
		log.Print(fmt.Sprintf("%v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(result))
}

func LoginNative(w http.ResponseWriter, r *http.Request) {
	var req account.LoginRequest
	var err error

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := cl.LoginNative(context.TODO(), &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		log.Print(fmt.Sprintf("%v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		log.Print(fmt.Sprintf("%v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(result))
}

func InitRestful() {
	service := web.NewService(
		web.Name("go.micro.web.user"),
	)

	service.Init()
	go service.Run()

	router := mux.NewRouter()
	// setup Greeter Server Client
	cl = account.NewAccountClient("go.micro.srv.user", client.DefaultClient)

	router.HandleFunc("/accounts/name", Update).Methods("POST")
	//router.HandleFunc("/accounts/user/create", Create).Methods("POST")
	service.HandleFunc("/accounts/user/create", Create)
	router.HandleFunc("/accounts/user/auth/native", LoginNative).Methods("POST")
	http.ListenAndServe(":8081", nil)
}