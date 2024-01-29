package domain

import "github.com/google/uuid"

type ItemId string // ItemId is a unique identifier for a item

type Consumer struct {
	Ref string // ref is a unique identifier for a consumer
}

type Item struct {
	ID        ItemId    `bson:"_id" gorm:"primaryKey;type:uuid"`
	Name      string    `bson:"name"`
	Qty       int       `bson:"qty"`
	Code      string    `bson:"code"`
	Version   int       `bson:"version"`
	Aggregate Aggregate `bson:"aggregate" gorm:"type:jsonb"`
}

func NewItem(name string, qty int, code string) (*Item, error) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	item := &Item{
		ID:      ItemId(ID.String()),
		Name:    name,
		Qty:     qty,
		Code:    code,
		Version: 1,
	}
	item.ItemCreatedEvent()
	return item, nil
}

func (i *Item) ItemCreatedEvent() {
	payload := map[string]interface{}{
		"itemId": i.ID,
		"name":   i.Name,
		"qty":    i.Qty,
		"code":   i.Code,
	}
	i.Aggregate.AppendEvent(ItemCreatedEvent, payload)
}

func (i *Item) Update(name string, qty int, code string) {
	i.Name = name
	i.Qty = qty
	i.Code = code
	i.ItemUpdatedEvent()
}

func (i *Item) ItemAllowcate(qty int, consumer Consumer) {
	i.Qty -= qty
	i.ItemAllowcatedEvent(consumer)
}

func (i *Item) ItemAllowcatedEvent(consumer Consumer) {
	payload := map[string]interface{}{
		"itemId": i.ID,
		"qty":    i.Qty,
		"ref":    consumer.Ref,
	}
	i.Aggregate.AppendEvent(ItemAllowcatedEvent, payload)
}

func (i *Item) ItemUpdatedEvent() {
	payload := map[string]interface{}{
		"itemId": i.ID,
		"name":   i.Name,
		"qty":    i.Qty,
		"code":   i.Code,
	}
	i.Aggregate.AppendEvent(ItemUpdatedEvent, payload)
}

func (i *Item) Delete() {
	i.ItemDeletedEvent()
}

func (i *Item) ItemDeletedEvent() {
	payload := map[string]interface{}{
		"itemId": i.ID,
	}
	i.Aggregate.AppendEvent(ItemDeletedEvent, payload)
}
