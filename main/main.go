package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/codegangsta/martini"
)

type pers struct {
	Id  		int	`sql:"AUTO_INCREAMENT" gorm:"primary_key"`
	Name 		string	`sql:"size:255;unique;index"`
	Vorname 	string	`sql:"size:255;unique;index"`
	Tel	 	string	`sql:"size:255;unique;index"`
	Email 		string	`sql:"size:255;unique;index"`
	ORT	 	string	`sql:"size:255;unique;index"`
	Region 		string	`sql:"size:255;unique;index"`
}

type session struct {
	Inf string
}

//var database []*person2 = make([]*person2,0)
var db *gorm.DB
var err error
//var sessionData person2 = person2{Id:-1}

var people  map[int]*pers

func index(rdr render.Render){
	rdr.HTML(200,"index", people)
}

var data [7]string

func newHandler(rdr render.Render, r *http.Request)  {

	//infM := r.FormValue("inf")
	nameM := r.FormValue("name")
	vornameM := r.FormValue("vorname")
	telM := r.FormValue("tel")
	emailM := r.FormValue("email")
	ortM := r.FormValue("ort")
	regionM := r.FormValue("region")

	//info[0] = infM
	data[1] = nameM
	data[2] = vornameM
	data[3] = telM
	data[4] = emailM
	data[5] = ortM
	data[6] = regionM

	db, err = gorm.Open("postgres","user=oosy password=oo dbname=gorm1 host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}



	db.CreateTable(pers{})
	jayna := pers{Name:nameM, Vorname:vornameM}
	db.Create(&jayna)
	u := []pers{}
	db.Find(&u)

	for _,value := range u{
		fmt.Println(value)
	}


	rdr.HTML(200,"new", data)

}

func checkInputHandler(rdr render.Render, r *http.Request)  {
/*
	db, err = gorm.Open("postgres","user=oosy password=oo dbname=gorm1 host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}


	
	db.CreateTable(person2{})
	jayna := person2{Name:data[1], Vorname:data[2], Tel:data[3], Email:data[4], ORT:data[5], Region:data[6]}
	db.Create(&jayna)
	u := []person2{}
	db.Find(&u)

	for _,value := range u{
		fmt.Println(value)
	}*/

	rdr.HTML(200, "ende", people)

}

func admin(rdr render.Render){
	rdr.HTML(200,"admin", people)
}


func main()  {
	fmt.Println("Application starts!")



	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "html",
		Extensions:[]string{".tmpl",".html"},
		Charset: "UTF-8",
		IndentJSON:true,
	}))

	m.Get("/", index)
	m.Get("/admin", admin)
	m.Get("/new", newHandler)
	m.Get("/ende", checkInputHandler)


	m.Run()
}