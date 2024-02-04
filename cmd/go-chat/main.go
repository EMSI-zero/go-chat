package main

import "github.com/EMSI-zero/go-chat/infra/boot"

func main() {
	if err := boot.Boot(); err != nil {
		panic(err)
	}

	
}
