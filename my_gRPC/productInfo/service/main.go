package main

import (
	"log"
	"net"

	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	// gPRC서버가 바인딩할 TCP 수신기 포트 생성
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 새로운 gPRC 서버 인스터스를 gRPC Go API에서 호출하여 생성
	s := grpc.NewServer()
	// 생성된 API를 호출하여 새로 생성된 gPRC 서버에 등록
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port " + port)
	// 포트에서 들어오는 메세지 수신
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
