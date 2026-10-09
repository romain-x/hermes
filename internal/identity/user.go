package identity

import (
	"net"
	"sync"
)


/* --------- User et ses fonctions --------- */


type User struct {
	name string
	conn net.Conn
	writeMutex sync.Mutex
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetConn() net.Conn {
	return u.conn
}


// La fonction écrit sur la connexion de l'utilisateur, de façon sécurisé
func (u *User) SafeWrite(b []byte) (int, error) {
	u.writeMutex.Lock()
	defer u.writeMutex.Unlock()
	return u.GetConn().Write(b)
}


/* --------- Fonctions manipulant la struct User --------- */

var userList = []*User{}


func addUser(u *User) {
	userList = append( userList, u )
}

func getUser(name string) *User {

	for u := range len(userList) {

		if name == userList[u].GetName() {
			return userList[u]
		}
	}

	return nil
}

// Retourne true si l'utilisateur est la.
func existsUserByName(name string) bool {

	for u := range len(userList) {

		if name == userList[u].GetName() {
			return true
		}
	}

	return false
}

func GetAllUserExcept(u *User) []*User {

	myList := []*User{}

	for i := range len(userList) {

		if userList[i].GetName() != u.GetName() {
			myList = append(myList, userList[i])
		}
	}

	return myList
}
