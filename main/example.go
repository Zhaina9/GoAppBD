package main

import(
	_"github.com/lib/pq"
	"database/sql"
	"log"
	)

type person struct{
	Id int
	Name string
	Surname string
	Age int
}

func main(){
	db,err := sql.Open("postgres", "user=oosy dbname=gorm1 password=oo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	people := make(map[int]*person,0)

	for rows.Next()  {

		pers := new(person)

		err := rows.Scan(&pers.Id,&pers.Name,&pers.Surname,&pers.Age)

		if err != nil{
			log.Fatal(err)
		}
		people[pers.Id] = pers
	}


}
