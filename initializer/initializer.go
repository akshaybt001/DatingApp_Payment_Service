package initializer

import (
	"github.com/akshaybt001/DatingApp_Payment_Service/internal/adapters"
	"github.com/akshaybt001/DatingApp_Payment_Service/internal/service"
	"github.com/akshaybt001/DatingApp_Payment_Service/usecase"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.PaymentEngine{
	adapter:=adapters.NewPaymentAdapter(db)
	usecases:=usecase.NewPaymentUsecase(adapter)
	services:=service.NewPaymentService(usecases)

	return service.NewPaymentEngine(services)
}