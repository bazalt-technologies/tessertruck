package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"sync"
)

type Store struct {
	mu   *sync.Mutex
	pool *pgxpool.Pool
}

func New(conn string) (*Store, error) {
	p, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &Store{mu: &sync.Mutex{}, pool: p}, nil
}

// intToInt32Array преобразует массив int в массив int32.
func intToInt32Array(in []int) []int32 {
	var out []int32
	for _, val := range in {
		out = append(out, int32(val))
	}
	return out
}

func (s *Store) Init() {
	dat, err := ioutil.ReadFile("../pkg/SQL/schema.sql")
	if err != nil {
		log.Println(err.Error())
	}

	_, err = s.pool.Exec(context.Background(), string(dat))
	if err != nil {
		log.Fatal(err.Error())
	}

	dat, err = ioutil.ReadFile("../pkg/SQL/migrations.sql")
	if err != nil {
		log.Println(err.Error())
	}

	_, err = s.pool.Exec(context.Background(), string(dat))
	if err != nil {
		log.Fatal(err.Error())
	}
}
