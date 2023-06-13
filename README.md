
# Synapsis Backend Developer

Challenge ini bertujuan untuk membuat Restful API menggunakan Java, GOlang .Untuk Itu saya membuat Online Store API ini menggunakan .

API ini adalah implementasi dari basic Vehicle API.

- Implementasi autentikasi, untuk mengakses API diperlukan token dari hasil login. 
- Autentikasi menggunakan JWT(JSON web Tokens) Tokens.
- Customer dapat melihat daftar produk berdasarkan Kategori
- Customer dapat menambahkan produk ke keranjang belanja
- Customer dapat melihat produk yang sudah ditambahkan ke keranjang belanja
- Customer dapat menghapus daftar belanjaan yang ada di keranjang
- Customers dapat checkout dan membuat pembayaran terhadap daftar belanjaan yang ada di keranjang
- Login and register customers
## RUN

- Clone this repository
- Run `docker-compose up` to run the API
- Run `docker-compose down` to stop the API

or if you want to setup the database in postgresql and configure the database name etc in **config/database.go and run the server

```bash
  go mod tidy
  go run main.go
```


## ERD:
 ![seru_backend_test_erd](https://user-images.githubusercontent.com/90734992/244950440-332dc314-3fb0-4c43-9696-b996a132fa37.jpeg)


 ## Directory Structure

saya menggunakan Konsep Clean Achitecture untuk membangun API ini, dimana setiap layer memiliki tanggung jawab masing-masing.
```

/config                                    * contains db configuration 
    |- database.go
/controller                                *contains all handler to handle request
    |- cart.go
    |- category.go
    |- product.go
    |- user.go
/docs                                     *contains json documentation of API
    |- docs.go
    |- swagger.json
    |- swagger.yaml
/middleware                               *contains middleware to check the request contains right token and validate he is admin
    |- auth.go
/models                                   *contains model of every entity to intteract with db
    |- cart.go
    |- category.go
    |- product.go
    |- user.go
/repositories                             *contains repository of every entity to intteract with db
    |- cart.go
    |- category.go
    |- product.go
    |- user.go
/usecase                                  *contains business logic of every entity
    |- cart.go
    |- category.go
    |- product.go
    |- user.go
/routes                                  *contains routes in API
    |- route.go
docker-compose.yml                       *docker compose file to run the API
Dockerfile                               *docker file to build the API
main.go                                  *Entry point of the API
  
```



## Run Locally

if we have docker installed in our machine, we can run this API using docker-compose

```bash
  docker-compose up
  docker-compose down
```

or 
run locally setup the database in postgresql and configure the database name etc in **config/db.config.js and run the server

```bash
  go mod tidy
  go run main.go
```


## API Reference

all the documentation of API can be found in **/docs/swagger.yaml
and you can see in swagger ui in https://rest-api-7qon5jxieq-et.a.run.app/swagger/

Before you can use This API first you need to register and login to get the token. and use the token to access the API

Actually if we want to make transaction . 
- First we need to add product to cart  (POST /cart)
- we can checkout the product at cart (POST /cart/checkout)
- automatically the api will created the order and detail product. After we finish checkout all the product in cart will be deleted. 
- lastly make payment (POST /payment) and the order will be updated to paid status
