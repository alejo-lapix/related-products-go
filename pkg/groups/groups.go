package groups

import (
	"time"
)

type Group struct {
	PrimaryProductID     *string
	AssociatedProductIDs []*string
	CreatedAt            *string
}

func NewGroup(primaryProductID *string, associatedProductIDs []*string) (*Group, error) {
	createdAt := time.Now().Format(time.RFC3339)

	return &Group{
		AssociatedProductIDs: associatedProductIDs,
		CreatedAt:            &createdAt,
	}, nil
}

type GroupRepository interface {
	Store(*Group) error
	Remove(ID *string) error
	FindByProduct(primaryProductID *string) (*Group, error)
}
