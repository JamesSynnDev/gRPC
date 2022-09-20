package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	// address로 서버와 연결을 설정한다.
	// 서버 클라이언트 사이에 보안되지 않은 연결을 만듬 (임시, 실사용 안됨)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 모든 작업이 완료되면 연결 종료
	defer conn.Close()

	// 연결을 전달하고 스텁을 만든다.
	// 모든 원격 메서드가 포함되어 있다.
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 14"
	description := `Meet Apple iPhone 14.`
	price := float32(2000.0)

	// 원격 호출할 context를 생성한다.
	// 최종 사용자ID, 인증 토큰, 요청 기한 같은 메타데이터를 포함하고
	// 요청 수명동안 존재한다.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 제품 정보로 addProduct 함수를 호출
	// 성공하면 ProductID를 리턴
	// 아니면 에러를 리턴 한다
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	// ProductID로 getProduct 함수를 호출
	// 성공하면 제품 상세 정보를 리턴
	// 아니면 에러를 리턴 한다
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())

}
