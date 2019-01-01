package main

import (
	"testing"
	"msu-go-11/8/04_json/structs"
)

var data = structs.Json{
	Tasks: []structs.Task{
		{Title: "first", Important: true},
		{Title: "first number two", Important: true},
		{Title: "second"},
		{Title: "third"},
	},
	Auth: struct {
		SessionID string `json:"session_id"`
		CsrfToken string `json:"csrf_token"`
	}{
		SessionID: "423j345k34h5lh425h34k54343454353rfcwedf",
		CsrfToken: "jflkwnlhtl24hrt3kl4t3k4bt34",
	},
	User: structs.User{
		Name: "Vasya",
		Age: 10303,
		Children: []structs.User{
			{Name: "Petya", Age: 1, Children:[]structs.User{
				{Name: "Ivan", Age: 7},
				{Name: "Valeriy", Age: 89, Children:[]structs.User{
					{Name: "Ivan", Age: 7},
					{Name: "Valeriy", Age: 89, Children: []structs.User{
						{Name: "Petya", Age: 1, Children:[]structs.User{
							{Name: "Ivan", Age: 7},
							{Name: "Valeriy", Age: 89, Children:[]structs.User{
								{Name: "Ivan", Age: 7},
								{Name: "Valeriy", Age: 89},
							}},
						}},
						{Name: "Ivan", Age: 7},
						{Name: "Valeriy", Age: 89, Children: []structs.User{
							{Name: "Petya", Age: 1, Children:[]structs.User{
								{Name: "Ivan", Age: 7},
								{Name: "Valeriy", Age: 89, Children:[]structs.User{
									{Name: "Ivan", Age: 7},
									{Name: "Valeriy", Age: 89},
								}},
							}},
							{Name: "Ivan", Age: 7},
							{Name: "Valeriy", Age: 89},
						}},
					}},
				}},
			}},
			{Name: "Ivan", Age: 7},
			{Name: "Valeriy", Age: 89},
		},
	},
}

func BenchmarkGenerated(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generated(data)
		}
	})
}

func BenchmarkReflect(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Reflect(data)
		}
	})
}
