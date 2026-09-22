package main

import (
	"context"
	"io"
	"log"

	pb "github.com/idiotcoffee/gRPC-learning/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	//unary
	feature, err := client.GetFeature(
		context.Background(),
		&pb.Point{
			Latitude:  406411633,
			Longitude: -741722051,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Feature: %s", feature.Name)

	// server-streaming
	stream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Latitude:  400000000,
			Longitude: -750000000,
		},
		Hi: &pb.Point{
			Latitude:  420000000,
			Longitude: -730000000,
		}})
	if err != nil {
		log.Fatal(err)
	}

	for {
		feature, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Printf(
			"Feature: %s (%d, %d)",
			feature.Name,
			feature.Location.Latitude,
			feature.Location.Longitude,
		)
	}

	recordStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	recordStream.Send(&pb.Point{
		Latitude:  1,
		Longitude: 2,
	})

	recordStream.Send(&pb.Point{
		Latitude:  3,
		Longitude: 4,
	})
	recordStream.Send(&pb.Point{
		Latitude:  5,
		Longitude: 6,
	})

	summary, err := recordStream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Points received: %d", summary.PointCount)
}
