package main

import "fmt"

type Monkey struct {
	Name string
}

// 爬山
func (m *Monkey) climbing() {
	fmt.Println(m.Name, "生来会爬树....")
}

// 声明接口
type Birder interface {
	Flying()
}

type Fisher interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}

func (lm LittleMonkey) Flying() {
	fmt.Println(lm.Name, "学会飞翔....")
}

func (lm LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "学会游泳....")
}

func main() {

	littleMonkey := LittleMonkey{
		Monkey{
			Name: "小猴子",
		},
	}

	littleMonkey.climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()

}
