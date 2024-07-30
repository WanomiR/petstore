package dbrepo

import (
	"backend/internal/lib/e"
	"backend/internal/modules/pet/entities"
	"context"
	"database/sql"
	"errors"
	"log"
)

func (db *PostgresDBRepo) GetPetById(ctx context.Context, petId int) (entities.Pet, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	// get pet
	query := `SELECT id, category_id, name, status FROM pets WHERE id = $1`

	var pet entities.Pet
	var categoryId int

	err := db.conn.QueryRowContext(ctx, query, petId).Scan(
		&pet.Id,
		&categoryId,
		&pet.Name,
		&pet.Status,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return entities.Pet{}, errors.New("pet not found")
	} else if err != nil {
		return entities.Pet{}, e.Wrap("failed to execute query", err)
	}

	// get pet category
	if pet.Category, err = db.GetPetCategoryById(ctx, categoryId); err != nil {
		return pet, e.Wrap("failed to get pet category", err)
	}

	// get pet photo urls
	if pet.PhotoUrls, err = db.GetPhotoUrlsByPetId(ctx, pet.Id); err != nil {
		return pet, e.Wrap("failed to get photo urls", err)
	}

	// get pet tags
	if pet.Tags, err = db.GetPetTagsById(ctx, pet.Id); err != nil {
		return pet, e.Wrap("failed to get pet tags", err)
	}

	return pet, nil
}

func (db *PostgresDBRepo) GetPetCategoryById(ctx context.Context, categoryId int) (entities.Category, error) {
	query := `SELECT id, name FROM categories WHERE id = $1`

	var category entities.Category

	err := db.conn.QueryRowContext(ctx, query, categoryId).Scan(&category.Id, &category.Name)

	if errors.Is(err, sql.ErrNoRows) {
		return entities.Category{}, errors.New("category not found")
	} else if err != nil {
		return entities.Category{}, e.Wrap("failed to execute query", err)
	}

	return category, nil

}

func (db *PostgresDBRepo) GetPhotoUrlsByPetId(ctx context.Context, petId int) ([]string, error) {
	query := `SELECT url FROM photo_urls WHERE pet_id = $1`

	rows, err := db.conn.QueryContext(ctx, query, petId)
	if err != nil {
		return nil, e.Wrap("failed to execute query", err)
	}
	defer rows.Close()

	urls := make([]string, 0)
	for rows.Next() {
		var url string
		if err = rows.Scan(&url); err != nil {
			return nil, e.Wrap("failed to scan row", err)
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (db *PostgresDBRepo) GetPetTagIdsByPetId(ctx context.Context, petId int) ([]int, error) {
	query := `SELECT tag_id FROM pet_tags WHERE pet_id = $1`

	rows, err := db.conn.QueryContext(ctx, query, petId)
	if err != nil {
		return nil, e.Wrap("failed to execute query", err)
	}
	defer rows.Close()

	tagIds := make([]int, 0)
	for rows.Next() {
		var tag int
		if err = rows.Scan(&tag); err != nil {
			return nil, e.Wrap("failed to scan row", err)
		}
		tagIds = append(tagIds, tag)
	}

	return tagIds, nil
}

func (db *PostgresDBRepo) GetPetTagsById(ctx context.Context, petId int) ([]entities.Tag, error) {
	tagIds, err := db.GetPetTagIdsByPetId(ctx, petId)
	if err != nil {
		return nil, e.Wrap("failed to get pet tags", err)
	}

	query := `SELECT id, name FROM tags WHERE id = $1`
	tags := make([]entities.Tag, 0)

	for _, tagId := range tagIds {
		var tag entities.Tag
		if err = db.conn.QueryRowContext(ctx, query, tagId).Scan(&tag.Id, &tag.Name); err != nil {
			log.Println(e.Wrap("failed to execute tags query", err).Error())
			continue
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
