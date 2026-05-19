package core

type OrderRepository interface {
	Save(order Order) error
	FindAll() ([]Order, error)
	FindByID(id uint) (*Order, error)
	Update(order Order) error
	Delete(id uint) error
}
