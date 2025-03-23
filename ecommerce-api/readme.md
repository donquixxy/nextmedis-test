# Ecommerce-api

Nextmedis backend test for number 1 question

# How to run

```bash
## Change directory to ecommerce-api folder
cd ecommerce-api
docker-compose up -d
```

# Collection
Import postman collection to test :)

## Explanation
- This app, has 2 api group; Public and Private.
- Public is the route group that should be able accessed by user without login / auth.
- Private is the route group that only accessible by logged-in user.




## API Reference

#### Get all Products (Public API)

```http
  GET /product
```

| Query | Type     | Description                | Default
| :-------- | :------- | :------------------------- | -----|
| `name` | `string` | **Optional**. name of product will be queried | null
| `page` | `integer` | **Optional**. page data| 1
| `limit` | `integer` | **Optional**. limit per page | 10
| `all` | `boolean` | **Optional**. true or false. if its true,  return all data without pagination ignoring query page and limit | false

#### Register User (Public API)

```http
  POST /register
```

| Payload | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of  user |
| `email`      | `string` | **Required**. Email of  user |
| `Password`      | `string` | **Required**. Password of  user |

#### Login (Public API)

```http
  POST /login
```

| Payload | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required**. Email of user |
| `Password`      | `string` | **Required**. Password of  user |


#### Add Item to Cart (Private API)

```http
  POST /api/cart
```

> 🔐 **Authorization:**  
> This endpoint requires a valid **JWT Token** in the `Authorization` header:
Authorization: Bearer <your-jwt-token-from-login>

| Payload | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `product_id`      | `string` | **Required**. Product ID |
| `quantity`      | `integer` | **Required**. Quantity of product |

#### Get Cart (Private API)

```http
  GET /api/cart
```

#### Place Order (Private API)

```http
  POST /api/order
```

| Payload | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `cart_id`      | `string` | **Required**. Cart id that will be  checked-out |
