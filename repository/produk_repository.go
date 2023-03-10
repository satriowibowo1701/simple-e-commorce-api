package repository

import (
	"context"
	"database/sql"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type ProductRepo interface {
	Create(ctx context.Context, tx *sql.Tx, produk model.ProdukRequest) error
	Update(ctx context.Context, tx *sql.Tx, produk model.ProdukUpdate) error
	FindById(ctx context.Context, tx *sql.Tx, produkId int) (*model.Produk, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (*model.Produk, error)
	FindAllAdmin(ctx context.Context, tx *sql.Tx) ([]*model.Produk, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Produk, error)
	DeleteById(ctx context.Context, tx *sql.Tx, productid int64) error
	UpdateQty(ctx context.Context, tx *sql.Tx, newqty int64, id int64) error
}
