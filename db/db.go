package db

import (
	//"database/sql"
	//"errors"
	/*"fmt"
	"log"
	"strings"*/
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	user "github.com/saromanov/user-srv/proto/account"
)

var (
	database string
	collection string
	session *mgo.Session
	coll *mgo.Collection
)

func Init() {
	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	coll = session.DB("users").C("users")
}

func CreateSession(sess *user.Session) error {
	if sess.Created == 0 {
		sess.Created = time.Now().Unix()
	}

	if sess.Expires == 0 {
		sess.Expires = time.Now().Add(time.Hour * 24 * 7).Unix()
	}

	//_, err := st["createSession"].Exec(sess.Id, sess.Username, sess.Created, sess.Expires)
	return nil
}

func DeleteSession(id string) error {
	//_, err := st["deleteSession"].Exec(id)
	return nil
}

func ReadSession(id string) (*user.Session, error) {
	return nil, nil
}

func Create(user *user.User, salt string, password string) error {
	user.Created = time.Now().Unix()
	user.Updated = time.Now().Unix()
	err := coll.Insert(user)
	//_, err := st["create"].Exec(user.Id, user.Username, user.Email, salt, password, user.Created, user.Updated)
	return err
}

func Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("%s is not a ObjectId value", id)
	}

	err := coll.Remove(bson.M{"id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}
	return nil
}

func Update(user *user.User) error {
	user.Updated = time.Now().Unix()
	//_, err := st["update"].Exec(user.Username, user.Email, user.Updated, user.Id)
	return nil
}

func Search()([]*user.User, error) {
	var acc []*user.User
	err := coll.Find(bson.M{}).All(&acc)
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func UpdateName(email, oldname, name string) error {
	var acc user.User
	err := coll.Find(bson.M{"email": email, "username": oldname}).One(&acc)
	if err != nil {
		return err
	}
	return coll.Update(bson.M{"email": email, "username": oldname}, bson.M{"$set":bson.M{"username": name}})
}

func GetPassword(email,username string) (string, string, error) {
	var acc user.User
	err := coll.Find(bson.M{"email": email}).One(&acc)
	if err != nil {
		return "", "", err
	}

	return acc.Password, acc.Username, nil
}

func Read(id string) (*user.User, error) {
	var u user.User
	err := coll.Find(bson.M{"id": id}).One(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetAccounts(limit,skip int)([]user.User, error) {
	var users []user.User
	err := coll.Find(bson.M{}).Skip(skip).Limit(limit).All(&users)
	return users, err
}

func FindAccount(email, pass string) {
	
}

func UpdatePassword(id string, salt string, password string) error {
	//_, err := st["updatePassword"].Exec(salt, password, time.Now().Unix(), id)
	return nil
}

func SaltAndPassword(username, email string) (string, string, error) {

	var salt, pass string

	return salt, pass, nil
}
