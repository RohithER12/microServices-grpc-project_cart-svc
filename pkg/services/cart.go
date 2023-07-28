package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/RohithER12/cart-svc/pkg/db"
	"github.com/RohithER12/cart-svc/pkg/models"
	"github.com/RohithER12/cart-svc/pkg/pb"
	"github.com/RohithER12/cart-svc/repo"
)

type Server struct {
	H    db.Handler
	Cart repo.Cart
	// CartItems repo.CartItems
	pb.UnimplementedCartServiceServer
}

// func NewServer(h db.Handler, cart repo.Cart, cartItems repo.CartItems) *Server {
// 	return &Server{
// 		H:         h,
// 		Cart:      cart,
// 		// CartItems: cartItems,
// 	}
// }

func (s *Server) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	fmt.Println("--------------\n\n\n", "userid", req.UserId)

	cart, err := s.Cart.GetByUserId(3)
	if err != nil {
		cart := models.Cart{
			UserId: req.UserId,
		}
		fmt.Println("--------------\n\n\n", cart)
		if err := s.Cart.CreateCart(cart); err != nil {
			return &pb.AddCartResponse{
				Status: http.StatusConflict,
				Error:  err.Error(),
			}, err
		}
	}
	fmt.Println("========================\n\n\n", cart)

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
	if err := s.Cart.AddCartItem(cartItem); err != nil {
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

	cartItem, err := s.Cart.GetCartItemByCartIdAndProductId(cart.Id, req.ProductId)
	if err != nil {
		return &pb.RemoveCartResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	if err := s.Cart.RemoveOneCartItem(cartItem); err != nil {
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
			Status: http.StatusNoContent,
			Error:  err.Error(),
		}, nil
	}
	carts, err := s.Cart.DisplayCart(cart.Id)
	if err != nil {
		return &pb.DisplayCartResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, errors.New("nothing inside cart")
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
