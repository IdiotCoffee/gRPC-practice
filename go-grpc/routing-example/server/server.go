package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/idiotcoffee/gRPC-learning/router"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedRouteGuideServer
	features []*pb.Feature
}

// unary
func (s *Server) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if feature.Location.Latitude == point.Latitude && feature.Location.Longitude == point.Longitude {
			return feature, nil
		}
	}
	return &pb.Feature{
		Location: point,
	}, nil
}

// server streaming
func (s *Server) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.features {
		if feature.Location.Latitude >= rect.Lo.Latitude && feature.Location.Latitude <= rect.Hi.Latitude && feature.Location.Longitude >= rect.Lo.Longitude && feature.Location.Longitude <= rect.Hi.Longitude {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

// client streaming
func (s *Server) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount: int32(count),
			})

		}
		if err != nil {
			return err
		}
		count++
	}
}

// bi-directional
func (s *Server) RouteChat(
	stream pb.RouteGuide_RouteChatServer,
) error {
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Recieved: %s", note.Message)

		if err := stream.Send(note); err != nil {
			return err
		}
	}

}

func main() {
	data, err := os.ReadFile("/home/void/Projects/gRPC-learning/routing-example/testdata/route_guide_db.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var features []*pb.Feature
	if err := json.Unmarshal(data, &features); err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", ":50051")
	if err := json.Unmarshal(data, &features); err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &Server{
		features: features,
	})
	log.Println("Server listening on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
