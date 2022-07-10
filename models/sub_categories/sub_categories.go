package models

type SubCategories struct {
	ID         int64
	CategoryID int64
	Name       string
	Form       string
	Status     bool
	UpdatedAt  *string
	CreatedAt  *string
}
