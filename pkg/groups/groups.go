package groups

import (
	"github.com/alejo-lapix/products-go/pkg/products"
)

type Group struct {
	PrimaryProductID *string
	Associations     []*Association
}

type Association struct {
	Product *products.Product
	Ratio   *float64
}

func NewGroup(primaryProductID *string, associations []*Association) (*Group, error) {
	// TODO perform validation

	return &Group{
		PrimaryProductID: primaryProductID,
		Associations:     associations,
	}, nil
}

type GroupRepository interface {
	Store(*Group) error
	Remove(ID *string) error
	FindByProduct(primaryProductID *string) (*Group, error)
}
