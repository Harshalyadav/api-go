package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CheckError(e error){

	if e !=nil{
      log.Fatalln(e)
	}
}

type Data struct{
	id int
	name string
}

func main() {
	ConnectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306/%v)",DbUser,DbPassword,DBName)

	db, err := sql.Open("mysql",ConnectionString)

	CheckError(err)
	defer db.Close()


	result, err:= db.Exec("insert into data values(4,'xyz')")
	CheckError(err)

	lastInsertId, err :=result.LastInsertId()
	fmt.Println("lastinsertid :",lastInsertId )

	CheckError(err)

	rowsAffected,err := result.RowsAffected()
	fmt.Println("rowAffected:",rowsAffected)



	rows, err := db.Query("SELECT * from data")
	CheckError(err)

	for rows.Next(){
		var data Data
		err:= rows.Scan(&data.id, &data.name)
		CheckError(err)
		fmt.Println(data)

		// ,,,,,,,,,,,,,,, go Build then run.....
	}
}