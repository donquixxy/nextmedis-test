package service

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	payload2 "ecommerce-api/modules/cart/payload"
	"ecommerce-api/modules/order/interfaces"
	"ecommerce-api/modules/order/payload"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type orderService struct {
	opt commons.Model
}

func (s orderService) SubmitOrder(ctx context.Context, data payload.SubmitOrder) (*model.Order, string, error) {
	// Retrieve cart  detail
	uCart, err := s.opt.Repository.Cart.CartGet(ctx, payload2.CartFilter{
		ID:        &data.CartID,
		WithItems: true,
		UserID:    &data.UserID,
	})

	if err != nil {
		return nil, "user cart not found", err
	}

	tx := s.opt.Database.Begin()

	orderPld := payload.OrderCreate{
		ID:     uuid.NewString(),
		UserID: uCart.UserID,
		Total:  uCart.Total,
		DbTx:   tx,
	}

	var items []*model.OrderItem
	for _, x := range uCart.Items {
		// Create order items
		pldItem := payload.OrderItemCreate{
			ID:        uuid.NewString(),
			OrderID:   orderPld.ID,
			ProductID: x.ProductID,
			Quantity:  x.Quantity,
			Price:     x.Price,
			SubTotal:  x.SubTotal,
			DbTx:      tx,
		}
		oItem, err := s.opt.Repository.Order.CreateItem(ctx, pldItem)

		if err != nil {
			log.Errorf("[SubmitOrder] - %v | %+v", err, fmt.Sprintf("+%v", pldItem))
			tx.Rollback()
			return nil, "Unable to create order. Please try again.", err
		}

		// Delete cart item
		err = s.opt.Repository.Cart.CartItemDelete(ctx, payload2.CartItemDelete{
			ID:   x.ID,
			DbTx: tx,
		})

		if err != nil {
			log.Errorf("[SubmitOrder] - %v | %+v", err, fmt.Sprintf("+%v", pldItem))
			tx.Rollback()
			return nil, "Unable to create order. Please try again.", err
		}

		items = append(items, oItem)
	}

	// Create order
	oData, err := s.opt.Repository.Order.Create(ctx, orderPld)

	if err != nil {
		log.Errorf("[SubmitOrder] - %v | %+v", err, fmt.Sprintf("+%v", orderPld))
		tx.Rollback()
		return nil, "Unable to create order. Please try again.", err
	}

	oData.Items = items

	// Delete cart
	if err = s.opt.Repository.Cart.CartDelete(ctx, payload2.CartDelete{
		ID: data.CartID,
	}); err != nil {
		log.Errorf("[SubmitOrder] - %v | %+v", err, fmt.Sprintf("+%v", oData))
		tx.Rollback()
		return nil, "Unable to create order. Please try again.", err
	}

	tx.Commit()

	return oData, "order created", nil
}

func NewOrderService(opt commons.Model) interfaces.OrderService {
	return &orderService{
		opt: opt,
	}
}
