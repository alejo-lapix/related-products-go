package groups

import (
	"github.com/alejo-lapix/products-go/pkg/products"
	"time"
)

type Group struct {
	PrimaryProductID *string
	Associations     []*Association
	CreatedAt        *string
}

type Association struct {
	Product *products.Product
	Ratio   *float
}

func NewGroup(primaryProductID *string, associations []*Association) (*Group, error) {
	createdAt := time.Now().Format(time.RFC3339)

	return &Group{
		PrimaryProductID: primaryProductID,
		Associations:     associations,
		CreatedAt:        &createdAt,
	}, nil
}

type GroupRepository interface {
	Store(*Group) error
	Remove(ID *string) error
	FindByProduct(primaryProductID *string) (*Group, error)
}
