package repository

import (
	"auth/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type TokenRepository interface {
	Create(ctx context.Context, token *models.RefreshToken) error
	FindByTokenHash(ctx context.Context, tokenHash string) (*models.RefreshToken, error)
	DeleteByTokenHash(ctx context.Context, tokenHash string) error
	DeleteExpired(ctx context.Context) error
	DeleteByUserID(ctx context.Context, userID string) error
}

type tokenRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewTokenRepository(db *pgxpool.Pool, logger *zap.Logger) TokenRepository {
	return &tokenRepository{db: db, logger: logger}
}

func (r *tokenRepository) Create(ctx context.Context, token *models.RefreshToken) error {
	//query := `
	//	INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, created_at)
	//	VALUES ($1, $2, $3, $4, NOW())
	//`
	//_, err := r.db.Exec(ctx, query, token.ID, token.UserID, token.TokenHash, token.ExpiresAt)
	//if err != nil {
	//	return helpers.Wrap("failed to create refresh token", err)
	//}
	//return nil
	panic("implement me")
}

func (r *tokenRepository) FindByTokenHash(ctx context.Context, tokenHash string) (*models.RefreshToken, error) {
	//query := `
	//	SELECT id, user_id, token_hash, expires_at, created_at
	//	FROM refresh_tokens
	//	WHERE token_hash = $1
	//`
	//row := r.db.QueryRow(ctx, query, tokenHash)
	//var token models.RefreshToken
	//err := row.Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)
	//if err != nil {
	//	return nil, helpers.Wrap("refresh token not found", err)
	//}
	//return &token, nil
	panic("implement me")
}

func (r *tokenRepository) DeleteByTokenHash(ctx context.Context, tokenHash string) error {
	//query := `DELETE FROM refresh_tokens WHERE token_hash = $1`
	//_, err := r.db.Exec(ctx, query, tokenHash)
	//if err != nil {
	//	return helpers.Wrap("failed to delete refresh token", err)
	//}
	//return nil
	panic("implement me")
}

func (r *tokenRepository) DeleteExpired(ctx context.Context) error {
	//query := `DELETE FROM refresh_tokens WHERE expires_at < $1`
	//_, err := r.db.Exec(ctx, query, time.Now())
	//if err != nil {
	//	return helpers.Wrap("failed to delete expired tokens", err)
	//}
	//return nil
	panic("implement me")
}

func (r *tokenRepository) DeleteByUserID(ctx context.Context, userID string) error {
	//query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	//_, err := r.db.Exec(ctx, query, userID)
	//if err != nil {
	//	return helpers.Wrap("failed to delete user tokens", err)
	//}
	//return nil
	panic("implement me")
}
