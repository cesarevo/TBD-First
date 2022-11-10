package store

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   string
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// make migration

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.New("file://migrations",
		connString)
	if err != nil {
		fmt.Println(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
	err = m.Force(1)
	if err := m.Down(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("migration is done")

	//11 is migrations version number, you may use your latest version
	if err != nil {
		fmt.Println(err)
	}
	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	var PeopleList []People
	rows, err := s.conn.Query(context.Background(), "select id, name from people")
	if err != nil {
		return nil, fmt.Errorf("list people: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var man People

		if err = rows.Scan(&man.ID, &man.Name); err != nil {
			fmt.Println("err is ", err)
		}
		PeopleList = append(PeopleList, man)
	}

	return PeopleList, err
}

func (s *Store) GetPeopleByID(id string) (People, error) {
	var name string
	err := s.conn.QueryRow(context.Background(), "select name from people where id=$1", id).Scan(&name)
	if err != nil {
		fmt.Print(err)
	}
	return People{ID: id, Name: name}, err
}
