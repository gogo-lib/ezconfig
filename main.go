package main

import (
	"fmt"

	"github.com/gogo-lib/ezconfig/pkg/structconfig"
	"github.com/gogo-lib/ezconfig/pkg/varconfig"
)

func main() {
	structconfig := structconfig.Get()
	postgresDsn1 := structconfig.PostgresLog.FormatDSN()
	fmt.Println(postgresDsn1)

	postgresDsn2 := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		varconfig.PostgresUser,
		varconfig.PostgresPasswd,
		varconfig.PostgresHost,
		varconfig.PostgresPort,
		varconfig.PostgresDb)
	fmt.Println(postgresDsn2)

	fmt.Println(structconfig.DiscordApp.Key, structconfig.DiscordApp.Secret)
	fmt.Println(structconfig.FacebookApp.Key, structconfig.FacebookApp.Secret)
}
