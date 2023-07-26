package services

import (
	"context"
	"net/http"

	"github.com/RohithER12/cart-svc/pkg/db"
	"github.com/RohithER12/cart-svc/pkg/models"
	"github.com/RohithER12/cart-svc/pkg/pb"
	"github.com/RohithER12/cart-svc/repo"
)

type Server struct {
	H         db.Handler
	Cart      repo.Cart
	CartItems repo.CartItems
	pb.UnimplementedCartServiceServer
}

func (s *Server) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {

	cart, err := s.Cart.GetByUserId(req.UserId)
	if err != nil {
		cart := models.Cart{
			UserId: req.UserId,
		}
		if err := s.Cart.CreateCart(cart); err != nil {
			return &pb.AddCartResponse{
				Status: http.StatusConflict,
				Error:  err.Error(),
			}, err
		}
	}
	if cart.Id == 0 {
		fetchCart, err := s.Cart.GetByUserId(req.UserId)
		if err != nil {
			return &pb.AddCartResponse{
				Status: http.StatusConflict,
				Error:  err.Error(),
			}, err
		}

		cart.Id = fetchCart.Id
	}

	cartItem := models.CartItems{
		CartId:    cart.Id,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}
	if err := s.CartItems.AddItem(cartItem); err != nil {
		return &pb.AddCartResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, err
	}

	return &pb.AddCartResponse{
		Status: http.StatusCreated,
	}, nil

}

func (s *Server) RemoveCart(ctx context.Context, req *pb.RemoveCartRequest) (*pb.RemoveCartResponse, error) {

	cart, err := s.Cart.GetByUserId(req.UserId)
	if err != nil {
		return &pb.RemoveCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	cartItem, err := s.CartItems.GetByCartIdAndProductId(cart.Id, req.ProductId)
	if err != nil {
		return &pb.RemoveCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	if err := s.CartItems.RemoveOne(cartItem); err != nil {
		return &pb.RemoveCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.RemoveCartResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) DisplayCart(ctx context.Context, req *pb.DisplayCartRequest) (*pb.DisplayCartResponse, error) {
	cart, err := s.Cart.GetByUserId(req.UserId)
	if err != nil {
		return &pb.DisplayCartResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}
	carts, err := s.Cart.DisplayCart(cart.Id)
	if err != nil {
		return &pb.DisplayCartResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	var response pb.DisplayCartResponse
	for _, cart := range carts {
		data := &pb.DisplayCartItem{
			ProductId: cart.ProductId,
			Quantity:  cart.Quantity,
			Price:     cart.Amount,
		}
		response.CartItems = append(response.CartItems, data)
	}

	return &response, nil
}
