package adapters

import (
	"github.com/akshaybt001/DatingApp_Payment_Service/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentAdapter struct {
	DB *gorm.DB
}

func NewPaymentAdapter(db *gorm.DB) *PaymentAdapter{
	return &PaymentAdapter{
		DB: db,
	}
}

func (p *PaymentAdapter) AddPayment(req entities.Payment) error{
	id := uuid.New()
	insertQuery:=`INSERT INTO payments (id,user_id,payment_ref,time) VALUES ($1,$2,$3,NOW())`
	if err:=p.DB.Exec(insertQuery,id,req.UserId,req.PaymentRef).Error;err!=nil{
		return err
	}
	return nil
}

func (p *PaymentAdapter) GetPayment(userId string) (entities.Payment,error){
	return entities.Payment{},nil
}

func (p *PaymentAdapter) AddSubscriptionPlan(req entities.Subscription)error{
	id :=uuid.New()
	addSubPlan:=`INSERT INTO subscriptions (id,amount,duration) VALUES ($1,$2,$3)`
	if err:=p.DB.Exec(addSubPlan,id,req.Amount,req.Duration).Error;err!=nil{
		return err
	}
	return nil
}

func (p *PaymentAdapter) UpdateSubscriptionPlan(req entities.Subscription) error {
	updateSubscriptionPlanQuery := `UPDATE subscriptions SET amount=$1,duration=$2 WHERE id=$3`
	if err := p.DB.Exec(updateSubscriptionPlanQuery, req.Amount, req.Duration, req.Id).Error; err != nil {
		return err
	}
	return nil
}

func (p *PaymentAdapter) GetAllSubsriptionPlans() ([]entities.Subscription,error){
	selctQuery:=`SELECT * FROM subscriptions`
	var res []entities.Subscription
	if err:=p.DB.Raw(selctQuery).Scan(&res).Error;err!=nil{
		return []entities.Subscription{},err
	}
	return res,nil
}


func (p *PaymentAdapter) GetSubscriptionPlanById(Id string) (entities.Subscription,error){
	selectQuery:=`SELECT * FROM subscriptions WHERE id=?`
	var res entities.Subscription
	if err:=p.DB.Raw(selectQuery,Id).Scan(&res).Error;err!=nil{
		return entities.Subscription{},err
	}
	return res,nil
}

