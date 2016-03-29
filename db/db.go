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
	/*sess := &user.Session{}

	r := st["readSession"].QueryRow(id)
	if err := r.Scan(&sess.Id, &sess.Username, &sess.Created, &sess.Expires); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return sess, nil*/
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

	err := coll.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
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

func Read(id string) (*user.User, error) {
	var u user.User
	err := coll.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func Search(username, email string, limit, offset int64) ([]*user.User, error) {
	/*var r *sql.Rows
	var err error

	if len(username) > 0 && len(email) > 0 {
		r, err = st["searchUsernameAndEmail"].Query(username, email, limit, offset)
	} else if len(username) > 0 {
		r, err = st["searchUsername"].Query(username, limit, offset)
	} else if len(email) > 0 {
		r, err = st["searchEmail"].Query(email, limit, offset)
	} else {
		r, err = st["list"].Query(limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer r.Close()*/

	var users []*user.User

	/*for r.Next() {
		user := &user.User{}
		var s, p string
		if err := r.Scan(&user.Id, &user.Username, &user.Email, &s, &p, &user.Created, &user.Updated); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("not found")
			}
			return nil, err
		}
		users = append(users, user)

	}
	if r.Err() != nil {
		return nil, err
	}*/

	return users, nil
}

func UpdatePassword(id string, salt string, password string) error {
	//_, err := st["updatePassword"].Exec(salt, password, time.Now().Unix(), id)
	return nil
}

func SaltAndPassword(username, email string) (string, string, error) {

	var salt, pass string

	return salt, pass, nil
}
