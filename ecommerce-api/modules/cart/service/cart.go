package service

import (
	"context"
	"ecommerce-api/commons"
	"ecommerce-api/model"
	"ecommerce-api/modules/cart/interfaces"
	"ecommerce-api/modules/cart/payload"
	payload2 "ecommerce-api/modules/product/payload"
	payload3 "ecommerce-api/payload"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type cartService struct {
	opt commons.Model
}

func (s *cartService) Get(ctx context.Context, filter payload.CartFilter) (*model.Cart, string, error) {
	result, err := s.opt.Repository.Cart.CartGet(ctx, filter)

	if err != nil {
		return nil, "cart not found", err
	}

	return result, "successfully retrieved", nil
}

func NewCartService(opt commons.Model) interfaces.CartService {
	return &cartService{
		opt: opt,
	}
}

func (s *cartService) AddItemCart(ctx context.Context, data payload.CartUpsert) (*model.Cart, string, error) {
	// Check if user already has cart; assuming 1 user 1 cart
	uCart, _ := s.opt.Repository.Cart.CartGet(ctx, payload.CartFilter{
		UserID: &data.UserID,
	})

	// Validate product
	pData, err := s.opt.Repository.Product.Get(ctx, payload2.ProductFilter{
		ID: &data.ProductID,
	})

	tx := s.opt.Database.Begin()

	if err != nil {
		log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
		tx.Rollback()
		return nil, "product not found", err
	}

	//  IF cart is not found for user;
	if uCart == nil {
		total := 0.0

		// Create new
		pld := payload.CartCreate{
			ID:     uuid.NewString(),
			UserID: data.UserID,
			DbTx:   tx,
		}

		if data.Quantity == nil && *data.Quantity == 0 {
			tx.Rollback()
			return nil, "quantity cannot be null", errors.New("quantity cannot be null")
		}

		subTotal := pData.Price * float64(*data.Quantity)
		cartItem, err := s.opt.Repository.Cart.CartItemCreate(ctx, payload.CartItemCreate{
			ID:        uuid.NewString(),
			CartID:    pld.ID,
			ProductID: pData.ID,
			Quantity:  *data.Quantity,
			SubTotal:  subTotal,
			Price:     pData.Price,
			DbTx:      tx,
		})

		if err != nil {
			tx.Rollback()
			log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
			return nil, "an error occurred in our system. please try again later", err
		}

		cartItem.Product = pData

		total += subTotal
		pld.Total = total
		cart, err := s.opt.Repository.Cart.CartCreate(ctx, pld)

		if err != nil {
			tx.Rollback()
			log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
			return nil, "an error occurred in our system. please try again later", err
		}

		cart.Items = append(cart.Items, cartItem)

		tx.Commit()
		return cart, "successfully created", nil
	}

	// Else if already has cart
	if data.Quantity == nil {
		tx.Rollback()
		return nil, "quantity cannot be null", errors.New("quantity cannot be null")
	}

	// Get cart item  under cart id
	ci, _ := s.opt.Repository.Cart.CartItemGet(ctx, payload.CartItemFilter{
		CartID:    &uCart.ID,
		ProductID: &data.ProductID,
	})

	// If not found; create  new cart item
	subTotal := pData.Price * float64(*data.Quantity)
	total := 0.0
	if ci == nil {
		ci, err = s.opt.Repository.Cart.CartItemCreate(ctx, payload.CartItemCreate{
			ID:        uuid.NewString(),
			CartID:    uCart.ID,
			ProductID: pData.ID,
			Quantity:  *data.Quantity,
			SubTotal:  subTotal,
			Price:     pData.Price,
			DbTx:      tx,
		})

		if err != nil {
			tx.Rollback()
			log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
			return nil, "an error occurred in our system please try again later", err
		}

		// Recalculate total on cart
		items, _, err := s.opt.Repository.Cart.CartItemGetAll(ctx, payload.CartItemFilter{
			NotInID: []string{ci.ID},
			Pagination: payload3.Pagination{
				All: true,
			},
		})

		if err != nil {
			tx.Rollback()
			log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
			return nil, err.Error(), err
		}

		items = append(items, ci)
		for _, x := range items {
			total += x.SubTotal
		}

		uCart, err = s.opt.Repository.Cart.CartUpdate(ctx, payload.CartUpdate{
			ID:    uCart.ID,
			Total: &total,
			DbTx:  tx,
		})

		if err != nil {
			tx.Rollback()
			log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
			return nil, err.Error(), err
		}
	} else {
		switch *data.Quantity {
		case 0:
			// Delete item in the cart
			if err = s.opt.Repository.Cart.CartItemDelete(ctx, payload.CartItemDelete{
				ID:   ci.ID,
				DbTx: tx,
			}); err != nil {
				tx.Rollback()
				log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
				return nil, "an error occurred in our system please try again later", err
			}

			// Recalculate
			items, _, err := s.opt.Repository.Cart.CartItemGetAll(ctx, payload.CartItemFilter{
				CartID:  &uCart.ID,
				NotInID: []string{ci.ID},
				Pagination: payload3.Pagination{
					All: true,
				},
			})

			for _, x := range items {
				total += x.SubTotal
			}

			uCart, err = s.opt.Repository.Cart.CartUpdate(ctx, payload.CartUpdate{
				ID:    uCart.ID,
				Total: &total,
				DbTx:  tx,
			})

			if err != nil {
				tx.Rollback()
				log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
				return nil, err.Error(), err
			}
			uCart.Items = items

		// If selected product exist in the cart item; update the quantity
		default:
			newSubTotal := pData.Price * float64(*data.Quantity)
			// Update quantity in the cart
			ci, err = s.opt.Repository.Cart.CartItemUpdate(ctx, payload.CartItemUpdate{
				ID:        uuid.NewString(),
				ProductID: &data.ProductID,
				Quantity:  data.Quantity,
				SubTotal:  &newSubTotal,
				Price:     &pData.Price,
				DbTx:      tx,
			})

			if err != nil {
				tx.Rollback()
				log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
				return nil, "an error occurred in our system please try again later", err
			}

			// Recalculate
			items, _, err := s.opt.Repository.Cart.CartItemGetAll(ctx, payload.CartItemFilter{
				CartID:  &uCart.ID,
				NotInID: []string{ci.ID},
				Pagination: payload3.Pagination{
					All: true,
				},
			})

			items = append(items, ci)
			for _, x := range items {
				total += x.SubTotal
			}

			uCart, err = s.opt.Repository.Cart.CartUpdate(ctx, payload.CartUpdate{
				ID:    uCart.ID,
				Total: &total,
				DbTx:  tx,
			})

			if err != nil {
				tx.Rollback()
				log.Errorf("[AddItemCart] - %v | %+v", err, fmt.Sprintf("+%v", data))
				return nil, err.Error(), err
			}
			uCart.Items = items
		}
	}

	tx.Commit()

	return uCart, "successfully updated", nil
}
