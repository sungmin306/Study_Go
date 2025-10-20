package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
