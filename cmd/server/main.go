package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/eltonCasacio/go-grpc/internal/database"
	pbCategory "github.com/eltonCasacio/go-grpc/internal/pb/category"
	pbCourse "github.com/eltonCasacio/go-grpc/internal/pb/course"
	"github.com/eltonCasacio/go-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	courseDB := database.NewCourse(db)
	courseService := service.NewCourseService(*courseDB)

	grpcServer := grpc.NewServer()
	pbCategory.RegisterCategoryServiceServer(grpcServer, categoryService)
	pbCourse.RegisterCourseServiceServer(grpcServer, courseService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
