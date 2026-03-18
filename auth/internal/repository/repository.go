package repository

import (
	"github.com/fin/tools/pkg/dbtools"
	"go.uber.org/zap"
)

type AuthRepository struct {
	dbtools.TxRepository
	logger *zap.Logger
}
