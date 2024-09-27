package service

import (
	"context"

	"github.com/waanvieira/desafio-CleanArch/internal/infra/grpc/pb"
	"github.com/waanvieira/desafio-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrderUseCase    usecase.GetOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getOrderUseCase usecase.GetOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrderUseCase:    getOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, input *pb.Blank) (*pb.OrderList, error) {
	orders, err := s.GetOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersReponse []*pb.Order

	for _, order := range orders {
		orderResponse := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}

		ordersReponse = append(ordersReponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersReponse}, nil
}
