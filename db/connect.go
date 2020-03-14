package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func conectar() *sql.DB {
	fmt.Println("Go com mysql")
	db, err := sql.Open("mysql", "root:foxtro2008@tcp(172.17.0.1:3306)/golang")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Banco conectado com sucesso!")
	return db
}
