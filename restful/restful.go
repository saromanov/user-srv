package restful

import (
	"log"

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

	//var resp account.UpdateNameResponse
	resp, err := cl.UpdateName(context.TODO(), &account.UpdateNameRequest{
		Email: "123@mail.ru",
		Oldusername:"past",
		Username:"future",
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(resp)
}

func InitRestful() {
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
	ws.Route(ws.POST("/name").To(say.UpdateName))
	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
