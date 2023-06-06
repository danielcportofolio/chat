package repositories

import (
	"database/sql"

	"github.com/danielcportofolio/chat-api/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (tag_name, avatar_url)
		VALUES ($1, $2)
		RETURNING id
	`

	err := ur.db.QueryRow(query, user.TagName, user.AvatarURL).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
