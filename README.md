## Microservices Demo

A microservice architecture is consists of services that have single responsibilities and do one thing well and implement a single business capability.

| Service                                              | Language      | Description                                                                                                                       |
| ---------------------------------------------------- | ------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| [frontend](/frontend)                           | Go            | Exposes an HTTP server to serve the website. Does not require signup/login and generates session IDs for all users automatically. |
| [cartservice](/cartservice)                     | Go            | Stores the items in the user's shopping cart in Redis and retrieves it.                                                           |
| [catalogservice](/productservice) | Go            | Provides the list of products from PostgreSQL searchs products and get individual products.                        |
| [paymentservice](/paymentservice)               | Node.js       | Charges the given credit card info (mock) with the given amount and returns a transaction ID.                                     |
| [shippingservice](/shippingservice)             | Go            | Gives shipping cost estimates based on the shopping cart. Ships items to the given address (mock).                                 |
| [emailservice](/emailservice)                   | Go        | Sends users an order confirmation email (mock).
| [identityservice](/identityservice)                   | Go        | Generates token and sends it to krakend to be signed and make operations on user .                                                                       
| [checkoutservice](/checkoutservice)             | Go            | Retrieves user cart, prepares order and plays a role as producer for the payment, shipping and the email services through rabbitmq.                             
| [api-gateway](/krakend)                 | Krakend | Sits between frontend and backend microservices and handles incoming requests and forward them to necessary services. 
| [rabbitmq]                 | RabbitMQ | Provides async. communication  between  checkout, payment, shipping and email services via fanout exchange.
| [user-db]                 | RabbitMQ | Stores user information.
| [catalog-db]                 | RabbitMQ | Stores product information.
| [redis]                 | Redis | Caches users' basket for 3 hours in map inside set structure.