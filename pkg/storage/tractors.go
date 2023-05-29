package storage

import (
	"context"
	"encoding/json"
	"log"
	"tracflow/pkg/models"
)

func (s *Store) GetTractors(ids []int) ([]models.Tractor, error) {
	rows, err := s.pool.Query(context.Background(), `
		SELECT id, name, info FROM tractors WHERE id = ANY($1) OR array_length($1,1) IS NULL`,
		intToInt32Array(ids),
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var data []models.Tractor
	for rows.Next() {
		var item models.Tractor
		var info models.Info
		var infoStr string
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&infoStr,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(infoStr), &info)
		item.Teledata = info
		data = append(data, item)
	}
	return data, err
}

func (s *Store) NewTractor(item models.Tractor) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
		INSERT INTO tractors(name, info) VALUES ($1, $2::JSONB) RETURNING id`,
		item.Name,
		item.Teledata,
	).Scan(&id)
	return id, err
}

func (s *Store) UpdateTractor(item models.Tractor) (int, error) {
	_, err := s.pool.Query(context.Background(), `
		UPDATE tractors SET
		name = $1,
		info = $2::JSONB,
		WHERE id = $3
	`,
		item.Name,
		item.Teledata,
		item.ID,
	)

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return item.ID, err
}
