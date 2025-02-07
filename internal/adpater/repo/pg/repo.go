package pg

// import (
// 	"github.com/sean-minngah/sil-backend-assessment/internal/core/util"
// 	"github.com/uptrace/bun"
// )

// type sqlRepo struct {
// 	db          bun.IDB
// 	productRepo port.ProductRepository
// }

// func NewSQLRepository(config util.Config) (port.Repository, error) {
// 	db, err := db.New(config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return create(db.DB()), nil
// }

// func create(db bun.IDB) port.Repository {
// 	return &sql.Repo{
// 		db:          db,
// 		productRepo: NewProductRepository(db),
// 	}
// }
