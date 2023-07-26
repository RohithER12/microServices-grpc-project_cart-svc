package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RohithER12/cart-svc/pkg/config"
	"github.com/RohithER12/cart-svc/pkg/db"
	"github.com/RohithER12/cart-svc/pkg/pb"
	"github.com/RohithER12/cart-svc/pkg/services"
	"github.com/RohithER12/cart-svc/repo"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var cartModule = wire.NewSet(
	repo.NewCartImpl,
	repo.NewCartItemsImpl,
	db.Init,
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Cart Svc on", c.Port)

	cart := InitializeCartImpl(&h)

	s := services.Server{
		H:    h,
		Cart: cart,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

func InitializeCartImpl(h *db.Handler) repo.Cart {
	wire.Build(cartModule)
	return nil
}
