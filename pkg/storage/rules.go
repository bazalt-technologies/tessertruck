package storage

import (
	"context"
	"log"
	"tracflow/pkg/models"
)

func (s *Store) GetRulesForTractor(tractorID int) ([]models.Rule, error) {
	rows, err := s.pool.Query(context.Background(), `
		SELECT id, tractor_id, name, field_name, val_int, val_float FROM rules WHERE tractor_id = $1`,
		tractorID,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var data []models.Rule
	for rows.Next() {
		var item models.Rule
		err = rows.Scan(
			&item.ID,
			&item.TractorID,
			&item.Name,
			&item.FieldName,
			&item.ValInt,
			&item.ValFloat,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}
	return data, err
}

func (s *Store) NewRule(item models.Rule) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
		INSERT INTO rules(tractor_id, name, field_name, val_int, val_float) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		item.TractorID,
		item.Name,
		item.FieldName,
		item.ValInt,
		item.ValFloat,
	).Scan(&id)
	return id, err
}

func (s *Store) UpdateRule(item models.Rule) (int, error) {
	_, err := s.pool.Query(context.Background(), `
		UPDATE rules SET
		tractor_id= $1,
		name = $2,
		field_name = $3,
		val_int = $4,
		val_float = $5
		WHERE id = $6
	`,
		item.TractorID,
		item.Name,
		item.FieldName,
		item.ValInt,
		item.ValFloat,
		item.ID,
	)

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return item.ID, err
}

func (s *Store) DeleteRule(id int) error {
	_, err := s.pool.Query(context.Background(), `
		DELETE FROM rules WHERE id = $1`, id)
	return err
}
