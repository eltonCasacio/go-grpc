package service

import (
	"context"
	"io"

	"github.com/eltonCasacio/go-grpc/internal/database"
	"github.com/eltonCasacio/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	db database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.NewCategoryRequest) (*pb.Category, error) {
	category, err := c.db.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) Categories(ctx context.Context, in *pb.Blanks) (*pb.CategoriesResponse, error) {
	categories, err := c.db.FindAll()
	if err != nil {
		return nil, err
	}

	output := []*pb.Category{}
	for _, c := range categories {
		output = append(output, &pb.Category{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}

	return &pb.CategoriesResponse{Categories: output}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.db.Find(in.ID)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoriesResponse{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryCreated, err := c.db.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			ID:          categoryCreated.ID,
			Name:        categoryCreated.Name,
			Description: categoryCreated.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryBidirectional(stream pb.CategoryService_CreateCategoryBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		categoryCreated, err := c.db.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			ID:          categoryCreated.ID,
			Name:        categoryCreated.Name,
			Description: categoryCreated.Description,
		})
		if err != nil {
			return err
		}
	}
}
