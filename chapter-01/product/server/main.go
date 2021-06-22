package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"lets-learn-golang-grpc/chapter-01/product/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	port       = ":50051"
	jsonDBFile = "sample_data/product_db.json"
)

type productInfoServer struct {
	pb.UnimplementedProductInfoServer
	savedProducts []*pb.Product
}

// GetProduct returns the product infomation
func (s *productInfoServer) GetProduct(ctx context.Context, productId *pb.ProductId) (*pb.Product, error) {
	for _, product := range s.savedProducts {
		if proto.Equal(product.Id, productId) {
			return product, nil
		}
	}
	// No product was found, return and unnamed product
	return &pb.Product{Id: productId}, nil
}

// ListProductsByYearRange lists all products published within the given year range (server-to-cient stream)
func (s *productInfoServer) ListProductsByYearRange(yearRange *pb.YearRange, stream pb.ProductInfo_ListProductsByYearRangeServer) error {
	for _, product := range s.savedProducts {
		if inRange(product.Year, yearRange) {
			if err := stream.Send(product); err != nil {
				return err
			}
		}
	}
	return nil
}

// inRange
func inRange(productYear int32, yearRange *pb.YearRange) bool {
	if productYear >= yearRange.Start && productYear <= yearRange.End {
		return true
	}
	return false
}

// loadProduccts loads products from a JSON file.
func (s *productInfoServer) loadProducts(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load default products: %v", err)
		}
	} else {
		data = exampleData
	}
	if err := json.Unmarshal(data, &s.savedProducts); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func newServer() *productInfoServer {
	s := &productInfoServer{}
	s.loadProducts(jsonDBFile)
	return s
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductInfoServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

var exampleData = []byte(`[
    {
        "id": {
            "value": 1
        },
        "name": "Harry Potter 1st",
        "Description": "Philosopher's Stone",
        "Year": 1997,
        "Price": 13.39
    },
    {
        "id": {
            "value": 2
        },
        "name": "Harry Potter 2nd",
        "description": "Chamber of Secrets",
        "year": 1998,
        "price": 14.59
    },
    {
        "id": {
            "value": 3
        },
        "Name": "Harry Potter 3rd",
        "Description": "Prisoner of Azkaban",
        "Year": 1999,
        "Price": 12.79
    },
    {"id": {
        "value": 4
        },
        "Name": "Harry Potter 4th",
        "Description": "Goblet of Fire",
        "Year": 2000,
        "Price": 15.49
    },
    {"id": {
        "value": 5
        },
        "Name": "Harry Potter 5th",
        "Description": "Order of the Phoenix",
        "Year": 2003,
        "Price": 13.79
    },
    {"id": {
        "value": 6
        },
        "Name": "Harry Potter 6th",
        "Description": "Half-Blood Prince",
        "Year": 2005,
        "Price": 14.29
    },
    {"id": {
        "value": 7
        },
        "Name": "Harry Potter 7th",
        "Description": "Deathly Hallows",
        "Year": 2007,
        "Price": 15.99
    }
]`)
