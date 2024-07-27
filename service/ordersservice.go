package service

import (
	"context"
	"invoice/common/dto"
	"invoice/common/entity"
	"invoice/repository"
)

type orderService struct {
	orderRepository repository.OrderRepository
}

type OrderService interface {
	GetAllOrders(context.Context, int) ([]dto.OrdersDTO, error)
	CreateNewOrders(context.Context, int, *dto.OrdersDTO) error
	DeleteOrder(context.Context, int) error
}

func NewOrderService(ord repository.OrderRepository) OrderService {
	return &orderService{
		ord,
	}
}

func (o *orderService) GetAllOrders(ctx context.Context, invID int) ([]dto.OrdersDTO, error) {
	lists, err := o.orderRepository.GetOrders(ctx, invID)
	if err != nil {
		return nil, err
	}

	var orders = make([]dto.OrdersDTO, 0)
	for _, l := range lists {
		ord := dto.OrdersDTO{
			ItemID:    l.ID,
			ItemName:  l.ItemName,
			Qty:       l.Qty,
			UnitPrice: l.UnitPrice,
		}
		orders = append(orders, ord)
	}

	return orders, nil
}

func (o *orderService) CreateNewOrders(ctx context.Context, invID int, or *dto.OrdersDTO) error {
	ord := &entity.OrderDB{
		ItemID: or.ItemID,
		Qty:    or.Qty,
	}

	if err := o.orderRepository.Insert(ctx, invID, ord); err != nil {
		return err
	}

	return nil
}

func (o *orderService) DeleteOrder(ctx context.Context, ordID int) error {
	if err := o.orderRepository.Delete(ctx, ordID); err != nil {
		return err
	}

	return nil
}
