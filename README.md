# Online Auction

## About The Project

![](https://i.imgur.com/27rEuH9.png)
The project is a full stack web application that allows users to bid on products in an online auction. It was built using Golang for the back-end, with a React front-end and a Postgres database. Redis is used to store session information for logged-in users, while JWT is used for authentication.

Users can create accounts and log in to view three product listings on the home page. Logged-in users can participate in live auctions by bidding on products. The frontend is constantly refreshed at regular intervals to keep the participants informed about the current prices.

## Technologies Used

![golang](https://badges.aleen42.com/src/golang.svg) ![react](https://badges.aleen42.com/src/react.svg)

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [React](https://react.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [GORM](https://gorm.io/)
- [Sass](https://sass-lang.com/)
- [Docker](https://www.docker.com/) with docker-compose

## Usage

![](https://i.imgur.com/r6xJn8B.gif)

## Run Locally

1. Clone the project
2. Enter the project folder

```bash
  $ cd kartaca-challenge
```

3. Run project with docker compose.

```bash
  $ docker compose up
```

You can reach to Application UI on [http://localhost:3000](http://localhost:3000)<br>
API Server will run on [http://localhost:8080](http://localhost:8080)


## How Does The Application Works ?

| Endpoints localhost:8080/ | Example Requests                                                                                   | Response                                                 |
| ------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------- |
| (POST) /v1/users/register | `{"email": "talhaunal@gmail.com", "password": "123456", "firstName": "talha", "lastName": "ünal"}` | Successfuly Registered                                   |
| (POST) /v1/users/login    | `{"email": "talhaunal@gmail.com", "password": "123456"}     `                                      | User Data <br> Set Cookie - <br> `Authorization : Token` |
| (POST) /v1/users/logout   | Set Cookie - ` Authorization : Token` (Token contains user id)                                     | Successfuly Logged Out                                   |
| (GET) /v1/products/all    | Set Cookie - ` Authorization : Token`                                                              | ` {Product 1 : Prd. Data}` `{Product 2 : Prd. Data} `    |
| (PUT) /v1/products/offer  | Set Cookie - `Authorization : Token` <br> `{ "productId":6, "offerPrice":400 }  `                  | Successfuly Offered                                      |
