package main

import "fmt"

// 通知类行为接口
type notifier interface {
	notify()
}

// 定义用户类型
type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 定义管理员类型
type admin struct {
	name  string
	email string
}

func (a admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

// 发通知
func sendNotification(n notifier) {
	n.notify()
}

func main() {

	u := user{name: "bamaw", email: "123321@gmail.com"}
	sendNotification(u)

}
