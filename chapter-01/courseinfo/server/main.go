package main

import (
	"context"
	"lets-learn-golang-grpc/chapter-01/courseinfo/pb"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct {
	courseMap map[string]*pb.Course
	pb.UnimplementedCourseInfoServer
}

// AddCourse implements pb.AddCourse
func (s *server) AddCourse(ctx context.Context, in *pb.Course) (*pb.CourseID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Course ID", err)
	}

	in.Id = id.String()
	if s.courseMap == nil {
		s.courseMap = make(map[string]*pb.Course)
	}
	s.courseMap[in.Id] = in
	return &pb.CourseID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetCourse inplements pb.GetCourse
func (s *server) GetCourse(ctx context.Context, in *pb.CourseID) (*pb.Course, error) {
	value, exists := s.courseMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Course does not exist.", in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseInfoServer(s, &server{})
	log.Printf("starting gRPC listener on port %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
