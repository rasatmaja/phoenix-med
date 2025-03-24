## API for an e-commerce platform
You are tasked with designing an API for an e-commerce platform. The system must support the following features:
- User registration and authentication
- Viewing and searching products
- Adding items to a shopping cart
- Completing a purchase

Design the RESTful endpoints for the above features. Describe your choice of HTTP methods (GET, POST, PUT, DELETE), URL structure, and the expected response formats. Assume that users need to authenticate before performing certain actions (e.g., adding items to the cart).

## Solution
The API design will follows URL structure pattern: `/api/v1/{domain}/{resource_id}/{action}`

Each component serves a specific purpose:
- `/api/v1`: Base URL that includes API namespace and version identifier for better maintainability
- `domain`: Represents the main resource or entity (e.g., products, cart, orders)
- `resource_id`: Unique identifier for accessing specific resource instances
- `action`: Optional segment for specific operations on the resource

> Example: `/api/v1/products/123/cart` - Adds product ID 123 to the shopping cart

The API implements standard HTTP methods following REST principles:
- `GET`: Retrieves resources from the server without modifying data (e.g., fetching product list, viewing cart contents)
- `POST`: Creates new resources or triggers actions (e.g., user registration, user login, adding items to cart)
- `PUT`: Updates existing resources (e.g., updating user profile, updating cart item quantity)
- `DELETE`: Removes resources from the server (e.g., removing an item from the cart)

The API will use JSON as the default response format for all endpoints and contains three main components:
1. `message`: A human-readable description of what happened with the request
2. `data`: The actual content/information returned by the API
3. `trace_id`: A unique identifier for tracking requests, especially useful for:
   - Error tracking across multiple services
   - Debugging issues in production
   - Users can reference this ID when reporting issues

```http
200 OK: Successful request.
{
  "message": "Successful request.",
  "data": {
    "key": "value"
  }
}

201 Created: Resource created successfully.
{
  "message": "Resource created successfully.",
  "data": {
    "key": "value"
  }
}

400 Bad Request: Invalid request parameters.
{
  "message": "Invalid request parameters.",
  "trace_id": "1234567890",
  "errors": [
    {
      "field": "username",
      "message": "Username is required."
    }
  ]
}

401 Unauthorized: Authentication required.
{
  "message": "Authentication required."
}

403 Forbidden: Access denied.
{
  "message": "Access denied."
}

404 Not Found: Resource not found.
{
  "message": "Resource not found."
}

500 Internal Server Error: Server error.
{
  "message": "Internal server error."
  "trace_id": "1234567890"
}
```

### Authentication Endpoints
- Register a new user.
  ```http
  POST /api/v1/register
  Content-Type: application/json
  {
    "username": "john_doe",
    "email": "john@email.com",
    "password": "password123"
  }

  Response: 201 Created
  Content-Type: application/json
  {
    "message": "User registered successfully"
  }
  ```

- Authenticate a user and generate a JWT token.
  ```http
  POST /api/v1/login
  Content-Type: application/json
  {
    "email": "EMAIL",
    "password": "password123"
  }

  Response: 200 OK
  Content-Type: application/json
  {
    "message": "Login successful",
    "data": {
      "access_token": "eyj....XXXX",
      "refresh_token": "eyj....YYYY"
    }
  }
  ```

### Product, Cart and Orders Endpoints
- Get a list and search of all products.
  ```http
  GET /api/v1/products?q=shirt&page=1&limit=10
  Response: 200 OK
  Content-Type: application/json
  {
    "message": "Product list retrieved successfully",
    "data": {
      "products": [
        {
          "id": "abc123",
          "name": "T-Shirt",
          "price": 20.0,
          "description": "A comfortable and stylish T-Shirt."
          "image_url": "assers/test.com/t-shirt.jpg"
        }
      ],
      "total": 1,
      "page": 1,
    }
  }
  ```

- Add a product to the shopping cart.
  ```http
  POST /api/v1/products/{product_id}/cart
  Authorization: Bearer {token}
  Content-Type: application/json
  
  Response: 201 Created
  Content-Type: application/json
  {
    "message": "Product added to cart successfully."
  }
  ```

- Get the contents of the shopping cart.
  ```http
  GET /api/v1/cart
  Authorization: Bearer {`token`}
  Response: 200 OK
  Content-Type: application/json
  {
    "message": "Cart retrieved successfully.",
    "data": {
      "items": [
        {
          "product_id": "abc123",
          "product_name": "T-Shirt",
          "product_price": 20.0,
          "product_description": "A comfortable and stylish T-Shirt."
          "product_image_url": "assers/test.com/t-shirt.jpg"
          "quantity": 2,
          "subtotal": 40.0
        }
      ],
      "total": 40.0
    }
  }
  ```

- Create a new order, reserve stock, and continue to payments.
  ```http
  POST /api/v1/orders
  Authorization: Bearer {`token`}
  Content-Type: application/json
  {
    "cart_items": [
      {
        "product_id": "abc123",
        "quantity": 2
      }
    ],
    "shipping_address": {
      "street": "123 Main St",
      "city": "Anytown",
      "state": "CA",
      "zip": "12345"
    },
    "payment_method": "credit_card"
  }

  Response: 200 OK
  Content-Type: application/json
  {
    "message": "Order created successfully.",
    "data": {
      "order_id": "order123",
      "payment_url": "/payments/order123",
      "payment_token": "payment_token_123"
    }
  }
  ```