package usecase

import "github.com/akshaybt001/DatingApp_Payment_Service/entities"

type PaymentUsecaseInterface interface {
	AddPayment(entities.Payment) error
	AddSubscriptionPlan(entities.Subscription) error
	UpdateSubscriptionPlan(req entities.Subscription) error
	GetAllSubscriptionPlans() ([]entities.Subscription, error)
	GetSubscriptionPlanById(Id string) (entities.Subscription, error)
	AddUserSubscription(userId, subId, duration string) error
	GetUserSubscription(userId string) (entities.UserSubscription, error)
	GetSubscriptionByDuration(duration string) (entities.Subscription, error)
}