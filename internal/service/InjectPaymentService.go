package service

type PaymentEngine struct{
	Srv *PaymentService
}

func NewPaymentEngine(srv *PaymentService) *PaymentEngine{
	return &PaymentEngine{
		Srv: srv,
	}
}