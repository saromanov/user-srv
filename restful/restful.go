package restful

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"github.com/saromanov/user-srv/proto/account"

	"golang.org/x/net/context"
)

type User struct{}

var (
	cl account.AccountClient
)

func (s *User) UpdateName(req *restful.Request, rsp *restful.Response) {
	log.Print("Received User.UpdateName API request")

	//name := req.PathParameter("name")
	email, err := req.BodyParameter("email")
	if err != nil {
		rsp.WriteError(500, err)
	}

	old, err := req.BodyParameter("oldUserName")
	if err != nil {
		rsp.WriteError(500, err)
	}

	newName, err := req.BodyParameter("newUserName")
	if err != nil {
		rsp.WriteError(500, err)
	}

	//var resp account.UpdateNameResponse
	resp, err := cl.UpdateName(context.TODO(), &account.UpdateNameRequest{
		Email: email,
		Oldusername:old,
		Username:newName,
	})

	if err != nil {
		rsp.WriteError(500,  err)
	}

	rsp.WriteEntity(resp)
}

func InitRestful2() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.srv.user"),
	)

	service.Init()

	// setup Greeter Server Client
	cl = account.NewAccountClient("go.micro.srv.user", client.DefaultClient)

	// Create RESTful handler
	say := new(User)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/accounts")
	ws.Route(ws.POST("/name").To(say.UpdateName)).Doc("Update name on user").
	Consumes("application/json")
	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

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
	var req account.User
	var err error

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := cl.Create(context.TODO(), &req)
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

func LoginNative(w http.ResponseWriter, r *http.Request) {

}

func InitRestful() {
	service := web.NewService(
		web.Name("go.micro.srv.user"),
	)

	service.Init()

	// setup Greeter Server Client
	cl = account.NewAccountClient("go.micro.srv.user", client.DefaultClient)

	http.HandleFunc("/accounts/name", Update)
	http.HandleFunc("/accounts/user/create", Create)
	http.ListenAndServe(":8081", nil)
}