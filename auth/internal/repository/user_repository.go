package repository

import (
	"auth/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.UserWithPassword) error
	FindByEmail(ctx context.Context, email string) (*models.UserWithPassword, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
}

//===================================================================================

type userRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *zap.Logger) UserRepository {
	return &userRepository{db: db, logger: logger}
}

func (r *userRepository) Create(ctx context.Context, user *models.UserWithPassword) error {
	//query := `
	//	INSERT INTO users (id, email, username, password_hash, created_at)
	//	VALUES ($1, $2, $3, $4, NOW())
	//`
	//_, err := r.db.Exec(ctx, query, user.ID, user.Email, user.Username, user.PasswordHash)
	//if err != nil {
	//	return helpers.Wrap("failed to create user", err)
	//}
	//return nil
	panic("implement me")
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.UserWithPassword, error) {
	//query := `
	//	SELECT id, email, username, password_hash, created_at
	//	FROM users
	//	WHERE email = $1
	//`
	//row := r.db.QueryRow(ctx, query, email)
	//
	//var user models.UserWithPassword
	//err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.CreatedAt)
	//if err != nil {
	//	return nil, helpers.Wrap("user not found", err)
	//}
	//user.Scan()
	//return &user, nil
	panic("implement me")
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	//query := `
	//	SELECT id, email, username, created_at
	//	FROM users
	//	WHERE id = $1
	//`
	//row := r.db.QueryRow(ctx, query, id)
	//
	//var user models.User
	//err := row.Scan(&user.ID, &user.Email, &user.Username, &user.createdAt)
	//if err != nil {
	//	return nil, helpers.Wrap("user not found", err)
	//}
	//user.Scan()
	//return &user, nil
	panic("implement me")
}
