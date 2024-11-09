package grpchandler

import (
	"context"

	"github.com/jonathangunawan/go-grpc/constant"
	"github.com/jonathangunawan/go-grpc/entity"
	"github.com/jonathangunawan/go-grpc/pb"
	"github.com/jonathangunawan/go-grpc/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductHandler struct {
	productUc usecase.ProductUsecaseItf
	pb.UnimplementedProductSvcServer
}

func NewProductHandler(puc usecase.ProductUsecaseItf) ProductHandler {
	return ProductHandler{
		productUc: puc,
	}
}

func (ph ProductHandler) GetAllProduct(context.Context, *emptypb.Empty) (*pb.GetAllProductResponse, error) {
	data := []*pb.Product{}
	res := ph.productUc.FindProduct()

	// compose response
	for _, val := range res {
		data = append(data, &pb.Product{
			Id:          val.ID,
			Name:        val.Name,
			Description: val.Description,
		})
	}

	return &pb.GetAllProductResponse{
		Data: data,
	}, nil
}

func (ph ProductHandler) InsertProduct(ctx context.Context, req *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {
	// input validation
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, constant.ErrValidationEmptyName)
	}

	if req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, constant.ErrValidationEmptyDescription)
	}

	// convert request to entity
	data := entity.Product{
		Name:        req.Name,
		Description: req.Description,
	}

	res, err := ph.productUc.AddProduct(data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// compose response from entity
	resp := &pb.InsertProductResponse{
		Id:          res.ID,
		Name:        res.Name,
		Description: res.Description,
	}

	return resp, nil
}
