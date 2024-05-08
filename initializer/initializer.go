package initializer

import (
	"github.com/akshaybt001/DatingApp_Payment_Service/concurrency"
	"github.com/akshaybt001/DatingApp_Payment_Service/internal/adapters"
	"github.com/akshaybt001/DatingApp_Payment_Service/internal/service"
	"github.com/akshaybt001/DatingApp_Payment_Service/usecase"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.PaymentEngine {
	adapter := adapters.NewPaymentAdapter(db)
	usecases := usecase.NewPaymentUsecase(adapter)
	services := service.NewPaymentService(usecases)
	c := concurrency.NewConcurrency(db, services)
	c.Concurrency()

	return service.NewPaymentEngine(services)
}
