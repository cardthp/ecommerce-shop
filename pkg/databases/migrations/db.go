package databases

import (
	"log"

	"github.com/cardthp/ecommerce-shop/config"
	_ "github.com/jackc/pgx/v5/stdlib" //force import > if not insert stdlib code can't find inside package
	"github.com/jmoiron/sqlx"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	// จริงๆสามารถดึง IConfig มาทั้งก้อนเลยได้ แต่ว่าควรแยกเป็นส่วนๆ ,*sqlx.DB เป็น output ที่ออกมาจาก function sqlx.Connect
	//Connect
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
