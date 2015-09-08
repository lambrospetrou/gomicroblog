package auth

import (
	"fmt"
	"github.com/lambrospetrou/lpgoauth"
	"io/ioutil"
	"net/http"
	"strings"
)

// BasicHandler creates a handler that encapsulates the given
// handler into another handler that performs Basic Authentication
func BasicHandler(fn http.HandlerFunc) http.HandlerFunc {
	return lpgoauth.BasicAuthHandler(IsBasicCredValid, fn)
}

// isBasicCredValid checks if the username:password are correct and valid
// against our database (file) of authorised users
func IsBasicCredValid(user string, pass string) bool {
	body, err := ioutil.ReadFile("sec/users.txt")
	if err != nil {
		fmt.Println("Cannot read users.txt!!!")
		return false
	}
	users := strings.Split(string(body), "\n")
	for _, u := range users {
		utokens := strings.SplitN(u, ":", 2)
		if lpgoauth.SecureCompare(user, utokens[0]) {
			return lpgoauth.SecureCompare(pass, utokens[1])
		}
	}
	return false
}
