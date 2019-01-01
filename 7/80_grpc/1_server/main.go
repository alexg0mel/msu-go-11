package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// не делайте так! не привыкайте к относительным путём
	// тут так только для того чтобы когда вы клонили себе репу - не надо было менять путь
	"../session"
)

// реализация grpc-сервиса
// сервис должен удовлетворять инфтерфейсу AuthCheckerClient (авто-сгенерённому в session.pb.go)
/*
type AuthCheckerClient interface {
	// создаёт сессию - принимает данные юзера, возвращает ID вессии
	CreateSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionID, error)
	// проверяет сессию - принимает ID сессии из куки - возвращает данные юзера
	CheckSession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error)
}
*/

type AuthChecker struct {
	sessionStorage map[string]*session.Session
}

func (ac *AuthChecker) CreateSession(ctx context.Context, sess *session.Session) (*session.SessionID, error) {
	fmt.Println("CreateSession", sess)
	// генерим ИД сессии
	id := RandStringRunes(10)
	// кладём в хранилище
	ac.sessionStorage[id] = sess
	// возвращаем новосозданную сессию клиенту
	return &session.SessionID{
		SessionID: id,
	}, nil
}

func (ac *AuthChecker) CheckSession(ctx context.Context, sid *session.SessionID) (*session.Session, error) {
	fmt.Println("CheckSession", sid)
	if sess, ok := ac.sessionStorage[sid.GetSessionID()]; ok {
		fmt.Println("session", sid.GetSessionID(), "found", sess)
		return sess, nil
	}
	fmt.Println("session", sid.GetSessionID(), "found")
	return nil, fmt.Errorf("Session %s not found", sid.GetSessionID())
}

// ----------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())

	// открываем порт
	lis, err := net.Listen("tcp", ":7001")
	PanicOnErr(err)
	// создаём grpc-сервер, к которому можно цеплять сервис
	server := grpc.NewServer()

	// регистрируем в сервере нашу реализацию сервиса, которая подходит под интерфейс
	session.RegisterAuthCheckerServer(server, &AuthChecker{
		sessionStorage: make(map[string]*session.Session),
	})

	// начинаем обрабатывать входящие соединения
	server.Serve(lis)
}

// ----------------------------------------------------------

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
