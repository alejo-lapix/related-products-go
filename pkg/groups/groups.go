package groups

type Group struct {
	PrimaryProductID *string        `json:"primaryProductId"`
	Associations     []*Association `json:"associations"`
}

type Association struct {
	ProductID *string  `json:"productId"`
	Ratio     *float64 `json:"ratio"`
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
