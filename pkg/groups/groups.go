package groups

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID                   *string
	PrimaryProductID     *string
	AssociatedProductIDs *[]string
	CreatedAt            *string
}

func NewGroup(primaryProductID *string, associatedProductIDs *[]string) (*Group, error) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)

	return &Group{
		ID:                   &id,
		AssociatedProductIDs: associatedProductIDs,
		CreatedAt:            &createdAt,
	}, nil
}

type GroupRepository interface {
	Store(*Group) error
	Remove(ID *string) error
	FindByProduct(primaryProductID *string) ([]*Group, error)
}
