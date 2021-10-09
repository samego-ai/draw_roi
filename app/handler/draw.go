package handler

import (
	"github.com/samego-ai/draw_roi/app/service"
	"github.com/samego-ai/draw_roi/pb"
	"golang.org/x/net/context"
)

// DrawServer
type DrawServer struct {
}

// MultiRectangle /**
func (server *DrawServer) MultiRectangle(ctx context.Context, in *pb.MultiRectangleRequest) (*pb.MultiRectangleResponse, error) {

	rectangles := make([]service.Rectangle, len(in.Rectangle))
	for index, rectangle := range in.Rectangle {
		rectangles[index] = service.Rectangle{
			X: int(rectangle.X),
			Y: int(rectangle.Y),
			H: int(rectangle.H),
			W: int(rectangle.W),
		}
	}

	svc := service.NewDrawService()
	image, err := svc.DrawMultiRectangle(in.Image, rectangles, in.Title)
	if nil != err {
		return &pb.MultiRectangleResponse{Code: 1, Message: err.Error()}, err
	}

	return &pb.MultiRectangleResponse{
		Code:    0,
		Message: "success",
		Data:    &pb.MultiRectangleResponseData{Image: image},
	}, nil
}
