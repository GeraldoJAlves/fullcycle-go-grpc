package main

import (
	"database/sql"
	"net"

	"github.com/geraldojalves/go-grpc/internal/database"
	"github.com/geraldojalves/go-grpc/internal/pb"
	"github.com/geraldojalves/go-grpc/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/fullcycle")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)

	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
