package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

type Person struct {
	Name        string
	ChineseName string
	Age         int
	Gender      string
	Email       string
}

type _Person struct {
	Name        string `faker:"name"`
	ChineseName string `faker:"chinese_name"`
	Age         int
	Gender      string `faker:"gender"`
	Email       string `faker:"email"`
}

type _Person_ struct {
	Name        string `faker:"name"`
	ChineseName string `faker:"chinese_name"`
	Age         int
	Gender      string `faker:"gender"`
	Email       string `faker:"email"`
}

func main() {

	var person Person
	var _person _Person
	var _Person_ _Person_

	if err := faker.FakeData(&person); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("使用 faker 生成数据：%+v\n", person)

	if err := faker.FakeData(&_person); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("使用 faker 生成数据：%+v\n", _person)

	if err := faker.FakeData(&_Person_); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("使用 faker 生成数据：%+v\n", _Person_)
}

// 结构体不加标签的情况下，生成的是随机数据
//使用 faker 生成数据：{Name:RaBpPTROBtlhQniRnZAOEiZCf ChineseName:kjMwmcaZXiEJujSWwgfCrbbJi Age:58 Gender:sFChOFrrNVBMLChsTZHVHrMnJ Email:pscrGGDfMxuRcAjWQgkMlirls}
// 结构体加标签的情况下，生成的是符合标签要求的数据
// 使用 faker 生成数据：{Name:Mrs. Caleigh Miller ChineseName:欧阳鸿朗 Age:11 Gender:Male Email:cEFZfXy@xNgWDUE.net}
