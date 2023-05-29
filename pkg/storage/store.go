package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
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
		log.Println(err.Error())
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
