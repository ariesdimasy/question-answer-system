package order_details

import (
	"context"
	"time"
)

type OrderDetailUsecase struct {
	contextTimeout time.Duration
	repository     OrderDetailRepository
}

func NewOrderDetailUsecase(repo OrderDetailRepository, timeout time.Duration) OrderDetailUseCaseDomain {
	return &OrderDetailUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *OrderDetailUsecase) GetAllOrderDetails(ctx context.Context) ([]OrderDetailDomain, error) {
	OrderDetails, err := uc.repository.GetAllOrderDetails(ctx)
	if err != nil {
		return []OrderDetailDomain{}, err
	}
	return OrderDetails, nil
}

func (uc *OrderDetailUsecase) GetByOrderId(ctx context.Context, order_id int) ([]OrderDetailDomain, error) {
	OrderDetails, err := uc.repository.GetByOrderId(ctx, order_id)
	if err != nil {
		return []OrderDetailDomain{}, err
	}
	return OrderDetails, nil
}

func (uc *OrderDetailUsecase) GetByOrderDetailId(ctx context.Context, id int) (OrderDetailDomain, error) {
	OrderDetail, err := uc.repository.GetByOrderDetailId(ctx, id)
	if err != nil {
		return OrderDetailDomain{}, err
	}
	return OrderDetail, nil
}

func (uc *OrderDetailUsecase) CreateOrderDetail(ctx context.Context, createOrderDetail OrderDetailDomain) (OrderDetailDomain, error) {
	OrderDetail, err := uc.repository.CreateOrderDetail(ctx, createOrderDetail)
	if err != nil {
		return OrderDetailDomain{}, err
	}
	return OrderDetail, nil
}

func (uc *OrderDetailUsecase) UpdateOrderDetail(ctx context.Context, updateOrderDetail OrderDetailDomain) (OrderDetailDomain, error) {
	OrderDetail, err := uc.repository.UpdateOrderDetail(ctx, updateOrderDetail)
	if err != nil {
		return OrderDetailDomain{}, err
	}
	return OrderDetail, nil
}

func (uc *OrderDetailUsecase) DeleteOrderDetail(ctx context.Context, id int) (OrderDetailDomain, error) {
	OrderDetail, err := uc.repository.DeleteOrderDetail(ctx, id)
	if err != nil {
		return OrderDetailDomain{}, err
	}
	return OrderDetail, nil
}
