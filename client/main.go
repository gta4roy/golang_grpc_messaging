package main

import (
	"context"
	"fmt"
	pb "grpc_messaging/order"

	"google.golang.org/grpc"
)

func main() {
	// Create NATS server connection
	var clientConnect *grpc.ClientConn

	opts := grpc.WithInsecure()
	clientConnect, err := grpc.Dial("localhost:7878", opts)
	if err != nil {
		fmt.Println(err)
	}

	purchaseServiceClient := pb.NewPurchaseOrderServiceClient(clientConnect)

	purchaseOrder := pb.PurchaseOrder{Productid: "73837837",
		Productname: "Sofa Set ",
		Cost:        737,
		Quantity:    2,
	}

	purchaseOrder2 := pb.PurchaseOrder{Productid: "83933",
		Productname: "Wash Room Set ",
		Cost:        1400,
		Quantity:    1,
	}

	purchaseOrder3 := pb.PurchaseOrder{Productid: "72873",
		Productname: "Laopala Dinner Set",
		Cost:        1400,
		Quantity:    1,
	}

	purchaseServiceClient.Add(context.Background(), &purchaseOrder)
	purchaseServiceClient.Add(context.Background(), &purchaseOrder2)
	purchaseServiceClient.Add(context.Background(), &purchaseOrder3)

	var emptyParams pb.EmptyParams

	allPurchases, _ := purchaseServiceClient.ListAll(context.Background(), &emptyParams)

	for _, orders := range allPurchases.GetOrders() {
		fmt.Println(*orders)
	}

	searchedOrders, _ := purchaseServiceClient.Search(context.Background(), &pb.SearchValue{Value: "83933"})

	fmt.Println(*searchedOrders)

	clientConnect.Close()
}
