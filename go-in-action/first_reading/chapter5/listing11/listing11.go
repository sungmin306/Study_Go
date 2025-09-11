package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Bill", "bill@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

}
