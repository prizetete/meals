package core

import "errors"

type OrderResponse struct {
	ID    uint    `json:"id"`
	Total float64 `json:"total"`
}

type OrderService interface {
	CreateOrder(order Order) error
	GetAllOrders() ([]OrderResponse, error)
	GetOrderByID(id uint) (*OrderResponse, error)
	UpdateOrder(id uint, order Order) (*OrderResponse, error)
	DeleteOrder(id uint) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	// Business logic function
	if order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}

func (s *orderServiceImpl) GetAllOrders() ([]OrderResponse, error) {
	orders, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]OrderResponse, 0, len(orders))
	for _, o := range orders {
		responses = append(responses, OrderResponse{
			ID:    o.ID,
			Total: o.Total,
		})
	}
	return responses, nil
}

func (s *orderServiceImpl) GetOrderByID(id uint) (*OrderResponse, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &OrderResponse{ID: order.ID, Total: order.Total}, nil
}

func (s *orderServiceImpl) UpdateOrder(id uint, order Order) (*OrderResponse, error) {
	if order.Total <= 0 {
		return nil, errors.New("Total must be positive")
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Total = order.Total
	if err := s.repo.Update(*existing); err != nil {
		return nil, err
	}
	return &OrderResponse{ID: existing.ID, Total: existing.Total}, nil
}

func (s *orderServiceImpl) DeleteOrder(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
