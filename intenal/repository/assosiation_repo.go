package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Association struct {
	ID          int64
	Title       string
	Description string
	DateTime    time.Time
	Done        bool
}

type AssociationRepository struct {
	dbPool *pgxpool.Pool
}

func NewAssociationRepository(conn *pgxpool.Pool) *AssociationRepository {
	return &AssociationRepository{
		dbPool: conn,
	}
}
