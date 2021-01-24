package main

import (
	"context"
	pb "grpc_messaging/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

type PurchaseOrderServer struct {
	purchaseOrders []pb.PurchaseOrder
}

func (s *PurchaseOrderServer) Search(ctx context.Context, in *pb.SearchValue) (*pb.PurchaseOrder, error) {

	var purchaseOrder pb.PurchaseOrder
	purchaseId := in.Value
	for _, order := range s.purchaseOrders {
		if order.Productid == purchaseId {
			purchaseOrder = order
			break
		}
	}
	return &purchaseOrder, nil
}

func (s *PurchaseOrderServer) Add(ctx context.Context, in *pb.PurchaseOrder) (*pb.ResponseMessage, error) {
	s.purchaseOrders = append(s.purchaseOrders, *in)
	var response pb.ResponseMessage
	response.Value = "successfully saved"
	return &response, nil
}

func (s *PurchaseOrderServer) ListAll(ctx context.Context, in *pb.EmptyParams) (*pb.AllPurchase, error) {
	var allPurchase pb.AllPurchase

	for _, value := range s.purchaseOrders {
		allPurchase.Orders = append(allPurchase.Orders, &value)
	}

	return &allPurchase, nil
}

func main() {

	// Create server connection

	port := ":7878"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Failed to listen %v ", err)
	}
	s := grpc.NewServer()
	pb.RegisterPurchaseOrderServiceServer(s, &PurchaseOrderServer{})
	log.Println("Server listening on the port :", port)
	s.Serve(lis)

	// Keep the connection alive

}
