package groups

type Group struct {
	PrimaryProductID *string
	Associations     []*Association
}

type Association struct {
	ProductID *string
	Ratio     *float64
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
