package products

import (
	"context"
	"time"

	repo "github.com/Abir-Zayn/kotoNilo/internal/adapters/postgresql/sqlc"
)

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Service interface {
	ListProducts(ctx context.Context) ([]Product, error)
	CreateProduct(ctx context.Context, p Product) (Product, error)
}

type svc struct {
	q repo.Querier
}

func NewService(q repo.Querier) Service {
	return &svc{q: q}
}

func (s *svc) ListProducts(ctx context.Context) ([]Product, error) {
	dbProducts, err := s.q.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, p := range dbProducts {
		products = append(products, Product{
			ID:        p.ID,
			Name:      p.Name,
			Price:     float64(p.PriceInCenters) / 100.0,
			Quantity:  int(p.Quantity),
			CreatedAt: p.CreatedAt.Time,
			UpdatedAt: p.UpdatedAt.Time,
		})
	}
	return products, nil
}

func (s *svc) CreateProduct(ctx context.Context, p Product) (Product, error) {
	params := repo.CreateProductParams{
		Name:           p.Name,
		PriceInCenters: int32(p.Price * 100),
		Quantity:       int32(p.Quantity),
	}

	savedProduct, err := s.q.CreateProduct(ctx, params)
	if err != nil {
		return Product{}, err
	}

	return Product{
		ID:        savedProduct.ID,
		Name:      savedProduct.Name,
		Price:     float64(savedProduct.PriceInCenters) / 100.0,
		Quantity:  int(savedProduct.Quantity),
		CreatedAt: savedProduct.CreatedAt.Time,
		UpdatedAt: savedProduct.UpdatedAt.Time,
	}, nil
}
