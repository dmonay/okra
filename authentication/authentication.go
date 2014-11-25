package authentication

import (
	// "fmt"
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter) {
	expiration := time.Now().Local()
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
	// fmt.Println("cookie: ", cookie)
}
