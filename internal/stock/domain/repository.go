package domain

type ItemRepository interface {
	Save(item *Item) error
	FindById(id ItemId) (*Item, error)
	FindAll() ([]*Item, error)
}
