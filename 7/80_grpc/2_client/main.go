package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"../session"
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

var Auth session.AuthCheckerClient

// основная работа тут
func checkSession(r *http.Request) *session.Session {
	// обработка сессии. код взять из 5/02_session
	cookieSessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil
	} else if err != nil {
		PanicOnErr(err)
	}
	// идём в grpc-серсис авторизации, вызывая метод, который на той стороне вызовет соответствующую реализацию
	sess, err := Auth.CheckSession(context.Background(), &session.SessionID{SessionID: cookieSessionID.Value})
	if err != nil {
		return nil
	}
	return sess
}

func main() {

	// соединяемся с нашим grpc-сервисом авторизации
	// grpc.WithInsecure() означает что мы не будет шифровать обмен данными с сервисом
	grcpConn, err := grpc.Dial("127.0.0.1:7001", grpc.WithInsecure())
	PanicOnErr(err)
	defer grcpConn.Close()

	// создаём клиент до нашего микросервиса
	Auth = session.NewAuthCheckerClient(grcpConn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sess := checkSession(r)
		if sess == nil {
			w.Write([]byte(loginFormTmpl))
			return
		}
		fmt.Fprint(w, "Welcome, "+sess.GetLogin()+", ua: "+sess.GetUseragent())
	})

	http.HandleFunc("/get_cookie", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		inputLogin := r.Form["login"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)

		// отправляем комнаду на создание сессии в сервис авторизации
		sess, err := Auth.CreateSession(context.Background(), &session.Session{
			// сохраняем в сессии логин
			Login: inputLogin,
			// ... и юзерагент
			Useragent: r.UserAgent(),
		})
		PanicOnErr(err)

		// теперь мы ничего не знаем про то как генерируется значение куки, его нам присылает сервис авторизации
		cookie := http.Cookie{Name: "session_id", Value: sess.GetSessionID(), Expires: expiration}
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
