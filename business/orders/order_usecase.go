package orders

import (
	"context"
	"time"
)

type OrderUsecase struct {
	contextTimeout time.Duration
	repository     OrderRepository
}

func NewOrderUsecase(repo OrderRepository, timeout time.Duration) OrderUseCaseDomain {
	return &OrderUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *OrderUsecase) GetAllOrders(ctx context.Context) ([]OrderDomain, error) {
	Orders, err := uc.repository.GetAllOrders(ctx)
	if err != nil {
		return []OrderDomain{}, err
	}
	return Orders, nil
}

func (uc *OrderUsecase) GetOrderById(ctx context.Context, id int) (OrderDomain, error) {
	Order, err := uc.repository.GetOrderById(ctx, id)
	if err != nil {
		return OrderDomain{}, err
	}
	return Order, nil
}

func (uc *OrderUsecase) GetOrderByCustomerId(ctx context.Context, id int) ([]OrderDomain, error) {
	Orders, err := uc.repository.GetOrderByCustomerId(ctx, id)
	if err != nil {
		return []OrderDomain{}, err
	}
	return Orders, nil
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, createOrder OrderDomain) (OrderDomain, error) {
	Order, err := uc.repository.CreateOrder(ctx, createOrder)
	if err != nil {
		return OrderDomain{}, err
	}
	return Order, nil
}

func (uc *OrderUsecase) UpdateOrder(ctx context.Context, updateOrder OrderDomain) (OrderDomain, error) {
	Order, err := uc.repository.UpdateOrder(ctx, updateOrder)
	if err != nil {
		return OrderDomain{}, err
	}
	return Order, nil
}

func (uc *OrderUsecase) DeleteOrder(ctx context.Context, id int) (OrderDomain, error) {
	Order, err := uc.repository.DeleteOrder(ctx, id)
	if err != nil {
		return OrderDomain{}, err
	}
	return Order, nil
}
