package main

import (
	"fmt"
	"os"

	"github.com/cardthp/ecommerce-shop/config"
	"github.com/cardthp/ecommerce-shop/modules/servers"
	databases "github.com/cardthp/ecommerce-shop/pkg/databases/migrations"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1] //use env from .env.dev , bin variable is os
	}
}

func main() {

	cfg := config.LoadConfig(envPath())

	//fmt.Println(cfg.Db())        // &{127.0.0.1 4444 tcp thanaphat 123456 ecommerce_db_test disable 25}
	//fmt.Println(cfg.App())       // &{127.0.0.1 3000 ecommerce-shop v0.1.0 60000000000 60000000000 10490000 2097000 ecommerce-shop-dev-bucket}
	//fmt.Println(cfg.App().Url()) // 127.0.0.1:3000

	db := databases.DbConnect(cfg.Db())
	defer db.Close() // ทำงานเป็นตัวสุดท้ายก่อน return function

	//fmt.Println(db)
	fmt.Println(servers.NewServer(cfg, db)) // &{0xc00013b900 0xc000008588 0xc000164210}
	servers.NewServer(cfg, db).Start()
}
