package main

import (
	"fmt"

	"github.com/kevinyang423/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
