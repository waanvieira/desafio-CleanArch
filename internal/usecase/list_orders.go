package usecase

import (
	"github.com/waanvieira/desafio-CleanArch/internal/entity"
)

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderUseCase) Execute() ([]*entity.Order, error) {
	ordersDb, err := c.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}

	var ordersModel []*entity.Order
	for _, order := range ordersDb {
		ordersModel = append(ordersModel, &entity.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersModel, nil
}
