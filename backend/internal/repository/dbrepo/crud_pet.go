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
	query := `SELECT id, category_id, name, status FROM pets WHERE id = $1 AND is_deleted = FALSE`

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
	if pet.Tags, err = db.GetTagsByPetId(ctx, pet.Id); err != nil {
		return pet, e.Wrap("failed to get pet tags", err)
	}

	return pet, nil
}

func (db *PostgresDBRepo) CreatePet(ctx context.Context, categoryId int, petName string, petStatus string) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT into pets (category_id, name, status, is_deleted) VALUES ($1, $2, $3, FALSE) returning id`

	var petId int

	if err := db.conn.QueryRowContext(ctx, query, categoryId, petName, petStatus).Scan(&petId); err != nil {
		return 0, e.Wrap("failed to execute query", err)
	}

	return petId, nil
}

func (db *PostgresDBRepo) DeletePet(ctx context.Context, petId int) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE pets SET is_deleted = TRUE WHERE id = $1`

	if _, err := db.conn.ExecContext(ctx, query, petId); err != nil {
		return e.Wrap("failed to execute query", err)
	}

	return nil
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

func (db *PostgresDBRepo) GetPetCategoryByName(ctx context.Context, categoryName string) (entities.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, name FROM categories WHERE name = $1`

	var category entities.Category

	err := db.conn.QueryRowContext(ctx, query, categoryName).Scan(&category.Id, &category.Name)

	if errors.Is(err, sql.ErrNoRows) {
		return entities.Category{}, errors.New("category not found")
	} else if err != nil {
		return entities.Category{}, e.Wrap("failed to execute query", err)
	}

	return category, nil
}

func (db *PostgresDBRepo) CreatePetCategory(ctx context.Context, categoryName string) (entities.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO categories (name) VALUES ($1) returning id`

	var categoryId int
	if err := db.conn.QueryRowContext(ctx, query, categoryName).Scan(&categoryId); err != nil {
		return entities.Category{}, e.Wrap("failed to execute query", err)
	}

	category := entities.Category{Name: categoryName, Id: categoryId}

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

func (db *PostgresDBRepo) CreatePetPhotoUrl(ctx context.Context, petId int, photoUrl string) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO photo_urls(pet_id, url) VALUES ($1, $2)`

	if _, err := db.conn.ExecContext(ctx, query, petId, photoUrl); err != nil {
		return e.Wrap("failed to execute query", err)
	}

	return nil
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

func (db *PostgresDBRepo) GetPetTagPair(ctx context.Context, petId int, tagId int) (entities.PetTag, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, pet_id, tag_id FROM pet_tags WHERE pet_id = $1 AND tag_id = $2`

	var petTag entities.PetTag
	err := db.conn.QueryRowContext(ctx, query, petId, tagId).Scan(&petTag.Id, &petTag.PetId, &petTag.TagId)

	if errors.Is(err, sql.ErrNoRows) {
		return entities.PetTag{}, errors.New("pet_tag not found")
	} else if err != nil {
		return entities.PetTag{}, e.Wrap("failed to execute query", err)
	}

	return petTag, nil
}

func (db *PostgresDBRepo) CreatePetTagPair(ctx context.Context, petId int, tagId int) (entities.PetTag, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO pet_tags(pet_id, tag_id) VALUES ($1, $2) returning id`

	var petTagId int
	if err := db.conn.QueryRowContext(ctx, query, petId, tagId).Scan(&petTagId); err != nil {
		return entities.PetTag{}, e.Wrap("failed to execute query", err)
	}

	petTag := entities.PetTag{Id: petTagId, PetId: petId, TagId: tagId}

	return petTag, nil
}

func (db *PostgresDBRepo) GetTagsByPetId(ctx context.Context, petId int) ([]entities.Tag, error) {
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

func (db *PostgresDBRepo) GetTagByName(ctx context.Context, tagName string) (entities.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, name FROM tags WHERE name = $1`

	var tag entities.Tag
	if err := db.conn.QueryRowContext(ctx, query, tagName).Scan(&tag.Id, &tag.Name); errors.Is(err, sql.ErrNoRows) {
		return entities.Tag{}, errors.New("tag not found")
	} else if err != nil {
		return entities.Tag{}, e.Wrap("failed to execute query", err)
	}

	return tag, nil
}
func (db *PostgresDBRepo) CreateTag(ctx context.Context, tagName string) (entities.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO tags(name) VALUES ($1) returning id`

	var tagId int
	if err := db.conn.QueryRowContext(ctx, query, tagName).Scan(&tagId); err != nil {
		return entities.Tag{}, e.Wrap("failed to execute query", err)
	}

	tag := entities.Tag{Id: tagId, Name: tagName}

	return tag, nil
}

func (db *PostgresDBRepo) UpdatePet(ctx context.Context, pet entities.Pet) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE pets SET category_id = $1, name = $2, status = $3 WHERE id = $4`

	if _, err := db.conn.ExecContext(ctx, query, pet.Category.Id, pet.Name, pet.Status, pet.Id); err != nil {
		return e.Wrap("failed to execute query", err)
	}

	return nil
}
