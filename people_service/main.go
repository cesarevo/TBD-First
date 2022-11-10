package main

import (
	"fmt"

	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service/store"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn := "postgresql://korshunov:123cesar4ik@95.217.232.188:7777/korshunov"
	s := store.NewStore(conn)
	fmt.Println(s.ListPeople())
	fmt.Println(s.GetPeopleByID("2"))

}
