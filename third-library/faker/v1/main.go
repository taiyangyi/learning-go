package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

func main() {

	chineseName := faker.ChineseName()
	fmt.Println("生成的中文姓名：", chineseName)

	name := faker.Name()
	fmt.Println("生成的姓名：", name)

	email := faker.Email()
	fmt.Println("生成的邮箱地址：", email)

	pwd := faker.Password()
	fmt.Println("生成的密码：", pwd)

	phoneNumber := faker.Phonenumber()
	fmt.Println("生成的电话号码：", phoneNumber)

	gender := faker.Gender()
	fmt.Println("生成的性别：", gender)

	year := faker.YearString()
	fmt.Println("生成的年份：", year)

	ipv4 := faker.IPv4()
	fmt.Println("生成IPv4：", ipv4)

	ipv6 := faker.IPv6()
	fmt.Println("生成IPv6：", ipv6)

	macAddress := faker.MacAddress()
	fmt.Println("生成mac地址：", macAddress)

	domainName := faker.DomainName()
	fmt.Println("生成域名：", domainName)

	url := faker.URL()
	fmt.Println("生成URL：", url)

	uuiddigit := faker.UUIDDigit()
	fmt.Println("生成UUID：", uuiddigit)

}

// OUTPUT:
// 生成的中文姓名： 公鸿德
// 生成的姓名： Lady Audie Reynolds
// 生成的邮箱地址： KpmPJgl@LGdKqxc.com
// 生成的密码： lpvaGDZQBAxZUgkPsSQuajHnskjHFsceadUNtsscEpvsbUgASe
// 生成的电话号码： 723-956-1041
// 生成的性别： Prefer to skip
// 生成的年份： 1986
// 生成IPv4： 53.220.70.183
// 生成IPv6： 64c0:1ca7:5a29:6207:6ab1:85b0:5b12:9c0d
// 生成mac地址： 7a:64:29:77:11:34
// 生成域名： GMqvyDy.info
// 生成URL： https://www.lAXxXgd.net/
// 生成UUID： 8369bbc3e1944a1aa871dc51228ed1e
