package usecase

import "github.com/akshaybt001/DatingApp_Payment_Service/internal/adapters"

type PaymentUsecase struct {
	adapters adapters.PaymentAdapterInterface
}

func NewPaymentUsecase(adapters adapters.PaymentAdapterInterface) *PaymentUsecase{
	return &PaymentUsecase{
		adapters: adapters,
	}
}