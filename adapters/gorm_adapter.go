package adapters

import (
	"komkrit/core"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormOrderRepository) FindAll() ([]core.Order, error) {
	var orders []core.Order
	if result := r.db.Find(&orders); result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (r *GormOrderRepository) FindByID(id uint) (*core.Order, error) {
	var order core.Order
	if result := r.db.First(&order, id); result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *GormOrderRepository) Update(order core.Order) error {
	if result := r.db.Save(&order); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormOrderRepository) Delete(id uint) error {
	if result := r.db.Delete(&core.Order{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
