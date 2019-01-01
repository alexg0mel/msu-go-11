package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var loginFormTmpl = `
<html>
	<body>
	<form action="/get_cookie" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`

var sessions = map[string]string{}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sessionID, err := r.Cookie("session_id")

		if err == http.ErrNoCookie {
			w.Write([]byte(loginFormTmpl))
			return
		} else if err != nil {
			PanicOnErr(err)
		}

		username, ok := sessions[sessionID.Value]

		if !ok {
			fmt.Fprint(w, "Session not found")
		} else {
			fmt.Fprint(w, "Welcome, "+username)
		}
	})

	http.HandleFunc("/get_cookie", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		inputLogin := r.Form["login"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)

		sessionID := RandStringRunes(32)
		sessions[sessionID] = inputLogin

		cookie := http.Cookie{Name: "session_id", Value: sessionID, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.ListenAndServe(":8081", nil)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
