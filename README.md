# e-com
### :handbag: A simple e-com microservices based projecct with Products and Order service.

**Product Service** : It Provides information about the product like availability, price, category.

**Order service**: It Provides information about the order like Order Value, DispatchDate, Order Status, Product Quantity.

User able to get the product catalogue and using that info should be able to place an order.


### How To Deploy on new instance?
Prerequisite
- ubuntu instance with docker and docker-compose installed [installation steps](https://docs.docker.com/engine/install/ubuntu/)

### Steps
- create new docker-compose.yml file in instance and paste following 
<pre>
version: "3.9"
services:
  ordersvc:
    build:
      context: https://github.com/Omkeshs/e-com.git
      dockerfile: ./order/Dockerfile
    ports:
      - 8080:8080

  productsvc:
    build:
      context: https://github.com/Omkeshs/e-com.git
      dockerfile: ./product/Dockerfile
    restart: always
    ports:
      - 8000:8000</pre>
      
run command
- sudo  docker-compose up --build
      

