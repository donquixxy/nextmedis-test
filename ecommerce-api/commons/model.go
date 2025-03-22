package commons

import (
	cart "ecommerce-api/modules/cart/interfaces"
	order "ecommerce-api/modules/order/interfaces"
	product "ecommerce-api/modules/product/interfaces"
	"ecommerce-api/modules/user/interfaces"
)

type Model struct {
	Options
	Service    *Service
	Repository *Repository
}

type Service struct {
	User    interfaces.UserService
	Product product.ProductService
	Cart    cart.CartService
	Order   order.OrderService
}

type Repository struct {
	User    interfaces.UserRepository
	Product product.ProductRepository
	Cart    cart.CartRepository
	Order   order.OrderRepository
}
