package main

import (
	"context"
	"lets-learn-golang-grpc/chapter-01/courseinfo/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	// Set up an unsecured connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Pass the connection and create a stub
	client := pb.NewCourseInfoClient(conn)

	// Smaple data
	name := "gRPC 101"
	title := "Let's learn gRPC with golang."
	var credit int32 = 4
	instructor := "abc101"

	// Create a Context to pass with the remote call.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call addCourse method with course details
	// Return: Course ID when it success, otherwise an error
	r, err := client.AddCourse(ctx, &pb.Course{
		Name:       name,
		Title:      title,
		Credit:     credit,
		Instructor: instructor})
	if err != nil {
		log.Fatalf("Could not add a course: %v", r.Value)
	}
	log.Printf("Course ID: %s added successfully", r.Value)

	// Call getCourse with the course ID
	// Return: course detail when it exists otherwise an error
	course, err := client.GetCourse(ctx, &pb.CourseID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get the course: %v", err)
	}
	log.Printf("Course: %v", course.String())
}
