package main

import (
	"context"
	"io"
	"lets-learn-golang-grpc/chapter-01/product/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	serverAddr = "localhost:50051"
)

func printProduct(client pb.ProductInfoClient, productId *pb.ProductId) {
	log.Printf("Getting product for product id: %v", productId.Value)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product, err := client.GetProduct(ctx, productId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(product)
}

func printProductsByYearRange(client pb.ProductInfoClient, yearRange *pb.YearRange) {
	log.Printf("Looking for products %v", yearRange)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListProductsByYearRange(ctx, yearRange)
	if err != nil {
		log.Fatalf("%v.ListProductByYearRange(_) = _, %v", client, err)
	}
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListProductsByYearRange(_) = _, %v", client, err)
		}
		log.Printf("Product: name: %q, year: %v", product.GetName(), product.GetYear())
	}
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewProductInfoClient(conn)

	// Looking for a valid product (Unary)
	printProduct(client, &pb.ProductId{Value: 1})

	// Product missing
	printProduct(client, &pb.ProductId{Value: 0})

	// Looking for products between 1999 and 2004 (Stream)
	printProductsByYearRange(client, &pb.YearRange{
		Start: 1999,
		End:   2004,
	})
}
