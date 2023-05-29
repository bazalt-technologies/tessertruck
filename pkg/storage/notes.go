package storage

import (
	"context"
	"tracflow/pkg/models"
)

func (s *Store) GetNotesForTractor(tractorID int) ([]models.Note, error) {
	rows, err := s.pool.Query(context.Background(), `
		SELECT id, tractor_id, message, time FROM notes WHERE tractor_id = $1`,
		tractorID,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var data []models.Note
	for rows.Next() {
		var item models.Note
		err = rows.Scan(
			&item.ID,
			&item.TractorID,
			&item.Message,
			&item.Time,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}
	return data, err
}

func (s *Store) NewNote(item models.Note) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
		INSERT INTO notes(tractor_id, message, time) VALUES ($1, $2, $3) RETURNING id`,
		item.TractorID,
		item.Message,
		item.Time,
	).Scan(&id)
	return id, err
}
