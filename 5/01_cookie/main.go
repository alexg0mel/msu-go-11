package main

import (
	"fmt"
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

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			w.Write([]byte(loginFormTmpl))
			return
		} else if err != nil {
			PanicOnErr(err)
		}
		fmt.Fprint(w, "Welcome, "+sessionID.Value)
	})

	http.HandleFunc("/get_cookie", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		inputLogin := r.Form["login"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   inputLogin,
			Expires: expiration,
		}
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
