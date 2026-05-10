package order

import (
	"time"

	"github.com/javohir01/eater-service/src/domain/order/models"
)

type OrderResponse struct {
	ID            string                 `json:"id"`
	EaterID       string                 `json:"eater_id"`
	Instruction   string                 `json:"instruction"`
	Restaurant    *models.RestaurantInfo `json:"restaurant"`
	Delivery      *models.DeliveryInfo   `json:"delivery"`
	Payment       *models.PaymentInfo    `json:"payment"`
	Items         []*models.OrderItem    `json:"items"`
	Status        string                 `json:"status"`
	PaymentStatus string                 `json:"payment_status"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}