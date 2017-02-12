# warehouse
Go service handling Expresso inventory

## API

### Item

#### `POST /api/item` creates and adds a new item to the database.

Example:

*Request:*
```
POST localhost:8080/api/item
{
    "roasterID": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "name": "Spring Blend",
    "coffeeType": "dark roast",
    "inStockBags": 50,
    "providerPrice": 4.00,
    "consumerPrice": 9.00,
    "ozInBag": 12.50
}
```

*Response:*
```
{
  "data": {
    "ID": "deb0a02f-f159-11e6-9563-acbc32977aaf",
    "RoasterID": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "Name": "Spring Blend",
    "Picture": "",
    "CoffeeType": "dark roast",
    "InStockBags": 50,
    "ProviderPrice": 4,
    "ConsumerPrice": 9,
    "OzInBag": 12.5
  },
  "success": true
}
```

#### `GET /api/item?offset=0&limit=20` returns up to `limit` items starting from `offset` when ordered by itemID

Example:

*Request:*
```
GET localhost:8080/api/item?offset=0&limit=20
```

*Response:*
```
{
  "data": [
    {
      "ID": "deb0a02f-f159-11e6-9563-acbc32977aaf",
      "RoasterID": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
      "Name": "Spring Blend",
      "Picture": "",
      "CoffeeType": "dark roast",
      "InStockBags": 50,
      "ProviderPrice": 4,
      "ConsumerPrice": 9,
      "OzInBag": 12.5
    }
  ],
  "success": true
}
```

#### `GET /api/item/:itemID` returns the item with the given itemID

Example:

*Request:*
```
GET localhost:8080/api/item/deb0a02f-f159-11e6-9563-acbc32977aaf
```

*Response:*
```
{
  "data": {
    "ID": "deb0a02f-f159-11e6-9563-acbc32977aaf",
    "RoasterID": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "Name": "Spring Blend",
    "Picture": "",
    "CoffeeType": "dark roast",
    "InStockBags": 50,
    "ProviderPrice": 4,
    "ConsumerPrice": 9,
    "OzInBag": 12.5
  },
  "success": true
}
```
#### `PUT /api/item/:itemID` updates the item with the given itemID to match the provided data. This just overrides values, so anything not present in the request will be set to NULL

Example:

*Request:*
```
PUT localhost:8080/api/item/1a7c63af-f15c-11e6-bad7-acbc32977aaf
{
    "id": "1a7c63af-f15c-11e6-bad7-acbc32977aaf",
    "roasterID": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "name": "Spring Blend",
    "pictureURL": "imgur.com",
    "coffeeType": "dark roast",
    "inStockBags": 50,
    "providerPrice": 4.00,
    "consumerPrice": 9.00,
    "ozInBag": 12.50
}
```
*Response:*
```
{
  "data": {
    "ID": "1a7c63af-f15c-11e6-bad7-acbc32977aaf",
    "RoasterID": "dd82cc65-d79d-11e6-9d4c-0242ac120004",
    "Name": "Spring Blend",
    "PictureURL": "imgur.com",
    "CoffeeType": "dark roast",
    "InStockBags": 50,
    "ProviderPrice": 4,
    "ConsumerPrice": 9,
    "OzInBag": 12.5
  },
  "success": true
}
```
#### `DELETE /api/item/:itemID` removes the item

Example:

*Request:*
```
DELETE localhost:8080/api/item/1a7c63af-f15c-11e6-bad7-acbc32977aaf
```

*Response:*
```
{
  "data": null,
  "success": true
}
```

### Order

#### `POST /api/order` creates and adds a new order to the database.

Example:

*Request:*
```
POST localhost:8080/api/order
{
    "userID": "dd82cc65-d79d-11e6-9d4c-0242ac120006",
    "subscriptionID": "dd82cc65-d79d-11e6-9d4c-0242ac120005"
}
```

*Response:*
```
{
  "data": {
    "ID": "9772d7ea-f15e-11e6-bad7-acbc32977aaf",
    "UserID": "dd82cc65-d79d-11e6-9d4c-0242ac120006",
    "SubscriptionID": "dd82cc65-d79d-11e6-9d4c-0242ac120005",
    "RequestDate": "2017-02-12T14:05:25.148465376-06:00",
    "ShipDate": "2020-01-01T01:00:00Z"
  },
  "success": true
}
```

#### `GET /api/order?offset=0&limit=20` returns up to `limit` orders starting from `offset` when ordered by orderID

Example:

*Request:*
```
GET localhost:8080/api/order?offset=0&limit=20
```

*Response:*
```
{
  "data": [
    {
      "ID": "9772d7ea-f15e-11e6-bad7-acbc32977aaf",
      "UserID": "dd82cc65-d79d-11e6-9d4c-0242ac120006",
      "SubscriptionID": "dd82cc65-d79d-11e6-9d4c-0242ac120005",
      "RequestDate": "2017-02-12T20:05:25Z",
      "ShipDate": "2020-01-01T01:00:00Z"
    }
  ],
  "success": true
}
```

#### `GET /api/order/:orderID` returns the order with the given orderID

Example:

*Request:*
```
GET localhost:8080/api/order/9772d7ea-f15e-11e6-bad7-acbc32977aaf
```

*Response:*
```
{
  "data": {
    "ID": "9772d7ea-f15e-11e6-bad7-acbc32977aaf",
    "UserID": "dd82cc65-d79d-11e6-9d4c-0242ac120006",
    "SubscriptionID": "dd82cc65-d79d-11e6-9d4c-0242ac120005",
    "RequestDate": "2017-02-12T20:05:25Z",
    "ShipDate": "2020-01-01T01:00:00Z"
  },
  "success": true
}
```
#### `PUT /api/order/:orderID` updates the order with the given orderID to match the provided data. This just overrides values, so anything not present in the request will be set to NULL

Example:

*Request:*
```
PUT localhost:8080/api/order/9772d7ea-f15e-11e6-bad7-acbc32977aaf
{
    "id": "9772d7ea-f15e-11e6-bad7-acbc32977aaf",
    "userID": "dd82cc65-d79d-11e6-9d4c-0242ac120006",
    "subscriptionID": "dd82cc65-d79d-11e6-9d4c-0242ac120005",
    "requestDate": "2017-02-12T20:05:25Z",
    "shipDate": "2021-01-01T01:00:00Z"
}
```
*Response:*
```
{
  "data": {
    "ID": "9772d7ea-f15e-11e6-bad7-acbc32977aaf",
    "UserID": "dd82cc65-d79d-11e6-9d4c-0242ac120006",
    "SubscriptionID": "dd82cc65-d79d-11e6-9d4c-0242ac120005",
    "RequestDate": "2017-02-12T20:05:25Z",
    "ShipDate": "2021-01-01T01:00:00Z"
  },
  "success": true
}
```
#### `DELETE /api/order/:orderID` removes the order

Example:

*Request:*
```
DELETE localhost:8080/api/order/9772d7ea-f15e-11e6-bad7-acbc32977aaf
```

*Response:*
```
{
  "data": null,
  "success": true
}
```