package service

import (
	"github.com/akshaybt001/DatingApp_Payment_Service/usecase"
	"github.com/akshaybt001/DatingApp_proto_files/pb"
	"google.golang.org/grpc"
)

type PaymentService struct {
	usecases usecase.PaymentUsecaseInterface
	UserConn pb.UserServiceClient
}

func NewPaymentService(usecases usecase.PaymentUsecaseInterface) *PaymentService{
	userConn,_:=grpc.Dial("localhost:8081",grpc.WithInsecure())
	return &PaymentService{
		usecases: usecases,
		UserConn: pb.NewUserServiceClient(userConn),
	}
}