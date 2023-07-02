package service

import (
	"context"

	"github.com/eltonCasacio/go-grpc/internal/database"
	pb "github.com/eltonCasacio/go-grpc/internal/pb/course"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB database.Course
}

func NewCourseService(db database.Course) *CourseService {
	return &CourseService{
		CourseDB: db,
	}
}

func (c *CourseService) CreateCourse(ctx context.Context, input *pb.NewCourse) (*pb.Course, error) {
	course, err := c.CourseDB.Create(input.Name, input.Description, input.CategoryId)
	if err != nil {
		panic(err)
	}

	return &pb.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryId:  course.CategoryID,
	}, nil
}

func (c *CourseService) GetCourses(ctx context.Context, in *pb.BlankCourse) (*pb.CoursesResponse, error) {
	courseDB, err := c.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}

	output := []*pb.Course{}
	for _, c := range courseDB {
		output = append(output, &pb.Course{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
			CategoryId:  c.CategoryID,
		})
	}

	return &pb.CoursesResponse{Courses: output}, nil
}

func (c *CourseService) GetCourseByID(ctx context.Context, in *pb.CourseID) (*pb.Course, error) {
	courseDB, err := c.CourseDB.FindByID(in.ID)
	if err != nil {
		return nil, err
	}
	return &pb.Course{
		ID:          courseDB.ID,
		Name:        courseDB.Name,
		Description: courseDB.Description,
		CategoryId:  courseDB.CategoryID,
	}, nil
}

func (c *CourseService) DeleteCourse(ctx context.Context, in *pb.CourseID) (*pb.BlankCourse, error) {
	err := c.CourseDB.DeleteCourse(in.ID)
	if err != nil {
		panic(err)
	}
	return nil, err
}
