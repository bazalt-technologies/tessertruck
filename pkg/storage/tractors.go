package storage

import (
	"context"
	"tracflow/pkg/models"
)

func (s *Store) GetTractors(ids []int) ([]models.Tractor, error) {
	rows, err := s.pool.Query(context.Background(), `
		SELECT id, name, create_date, use_date, use_place FROM tractors WHERE id = ANY($1) OR array_length($1,1) IS NULL`,
		intToInt32Array(ids),
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var data []models.Tractor
	for rows.Next() {
		var item models.Tractor
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.CreateDate,
			&item.UseDate,
			&item.UsePlace,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, err
}

func (s *Store) NewTractor(item models.Tractor) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
		INSERT INTO tractors(name, create_date, use_date, use_place, info ) VALUES ($1, $2, $3, $4, $5::JSONB) RETURNING id`,
		item.Name,
		item.CreateDate,
		item.UseDate,
		item.UsePlace,
		item.Teledata,
	).Scan(&id)
	return id, err
}

func (s *Store) UpdateTractor(item models.Tractor) (int, error) {
	_, err := s.pool.Query(context.Background(), `
		UPDATE tractors SET
		name = $1,
		create_date= $2, 
		use_date = $3,
		use_place = $4, 
		info = $5::JSONB,
		WHERE id = $6
	`,
		item.Name,
		item.CreateDate,
		item.UseDate,
		item.UsePlace,
		item.Teledata,
		item.ID,
	)

	if err != nil {
		return 0, err
	}
	return item.ID, err
}
