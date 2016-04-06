package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	micro "github.com/micro/go-micro"
	proto "github.com/saromanov/user-srv/proto/account"
	auth  "github.com/saromanov/auth-srv/proto/account"
	oauth2 "github.com/saromanov/auth-srv/proto/oauth2"
	//"github/com/saromanov/user-srv/handler"
	"golang.org/x/net/context"
)

//Login with rpc, rest, native auth and oauth2

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func CreateUser() (string, string, string, error) {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.srv.user"))

	// Create new greeter client
	greeter := proto.NewAccountClient("go.micro.srv.user", service.Client())

	//Generate random username, email and password
	email := randStringBytes(5) + "@mail.com"
	username := randStringBytes(7)
	password := randStringBytes(10)
	// Call the greeter
	rsp, err := greeter.Create(context.TODO(), &proto.User{Email: email, Username: username,
		Password: password})
	if err != nil {
		return "", "", "", err
	}

	// Print response
	fmt.Println(rsp.Id)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Password: %s\n", password)

	return username, email, password, nil
}

func NativeLoginRPC(email, password string) error {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.srv.user"))

	// Create new greeter client
	greeter := proto.NewAccountClient("go.micro.srv.user", service.Client())

	rsp, err := greeter.LoginNative(context.TODO(), &proto.LoginRequest{Email: email, Password: password})
	if err != nil {
		return err
	}

	fmt.Println(rsp)
	return nil
}

func NativeLoginREST(email, password string) error {
	evreq := &proto.LoginRequest{Email: email, Password: password}
	res, _ := json.Marshal(evreq)
	resp, err := postRequest("http://localhost:8081/accounts/user/auth/native", res)
	if err != nil {
		return err
	}
	var postResp map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return err
	}
	fmt.Println(postResp)
	return nil
}

func CreateApp(id, secret string) error {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.srv.auth"))

	// Create new greeter client
	greeter := auth.NewAccountClient("go.micro.srv.auth", service.Client())

	_, err := greeter.Create(context.TODO(), &auth.CreateRequest{Account: &auth.Record{ClientId: id, ClientSecret: secret, Type:"user"}})
	if err != nil {
		return err
	}

	fmt.Println("New app was created")

	return nil
}

func Oauth2Authorize(id, secret string) (string, error) {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.srv.auth"))

	// Create new greeter client
	greeter := oauth2.NewOauth2Client("go.micro.srv.auth", service.Client())

	rsp, err := greeter.Authorize(context.TODO(), &oauth2.AuthorizeRequest{ClientId: id, State:"mystatetoken", ResponseType:"code", RedirectUri:"https://foo.bar.com"})
	if err != nil {
		return "", err
	}

	fmt.Println("Code: ", rsp.Code)
	return rsp.Code, nil
}

func Oauth2LoginRPC(id, secret, code string) error {
	return nil
}


func Oauth2LoginREST(id, secret, code string) error {
	return nil
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// postRequest is helpful method for build post requests for microservice
func postRequest(url string, params []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("1. Create User")
	username, email, pass, err := CreateUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n2. Native login via RPC")
	err = NativeLoginRPC(email, pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n3. Native login via REST")
	err = NativeLoginREST(email, pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n4. Create new app for RPC auth")
	err = CreateApp(username, pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n5. Oauth2 auth")
	code, err := Oauth2Authorize(username, pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n6. Oauth2 login")
	Oauth2LoginRPC(username, pass, code)
}
