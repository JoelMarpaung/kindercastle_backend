# KinderCastle Backend API

This is the documentation of how to set up and run this project and also explanation for each API request 
Here is the [link](https://kindercastle.up.railway.app/) of demo API
## Preparation

1. Make sure you already install [go](https://go.dev/doc/install).
2. Make sure mysql already installed in your system

## Installation
Clone this project
```bash
git clone https://github.com/JoelMarpaung/kindercastle_backend.git
```

Move into the directory
```bash
cd kindercastle_backend
```

Copy the .env.example to .env
```bash
cp .env.example .env
```

Now make the database and use that setting on the .env

Run the Migration using make command
```bash
make migrate
```

Run the project using make command
```bash
make run
```

## List API
To Access all the API, you must provide Authorization Header which is the token that you get from login using the firebase

```python
Authorization Bearer {{token}}
```

### 1. /v1/image [POST]
This API for uploading image file and then get the image_url to be used on image_url payload for the other API
```python
Payload
Image FormFile

Response : Code 200 OK
{
  "image_url": "string"
}
```


### 2. /v1/books [GET]
This API for getting all the book list
```python
Parameter
search string => For querying the book title
limit int => For limit the result data
page int => For setting the page of the result data

Response : Code 200 OK
{
  "items": [
    {
      "author": "string",
      "book_ownership": true,
      "created_at": "string",
      "description": "string",
      "edition": "string",
      "format": "string",
      "genre": "string",
      "id": "string",
      "image_url": "string",
      "isbn": "string",
      "language": "string",
      "number_of_pages": 0,
      "price": 0,
      "publication_date": "string",
      "publisher": "string",
      "title": "string",
      "updated_at": "string",
      "user_id": "string"
    }
  ],
  "limit": 60,
  "page": 1,
  "search": "keyword",
  "total_page": 1,
  "total_row": 1
}
```

### 3. /v1/books [POST]
This API for creating the book data.
```python
Payload
{
  "author": "string",
  "description": "string",
  "edition": "string",
  "format": "string",
  "genre": "string",
  "image_url": "string",
  "isbn": "string",
  "language": "string",
  "number_of_pages": 0,
  "price": 0,
  "publication_date": "string",
  "publisher": "string",
  "title": "string",
  "user_id": "string"
}

Response : Code 201 Created
```

### 4. /v1/books/me [GET]
This API for getting all the book list that has ownership by the login user.
```python
Parameter
search string => For querying the book title
limit int => For limit the result data
page int => For setting the page of the result data

Response : Code 200 OK
{
  "items": [
    {
      "author": "string",
      "book_ownership": true,
      "created_at": "string",
      "description": "string",
      "edition": "string",
      "format": "string",
      "genre": "string",
      "id": "string",
      "image_url": "string",
      "isbn": "string",
      "language": "string",
      "number_of_pages": 0,
      "price": 0,
      "publication_date": "string",
      "publisher": "string",
      "title": "string",
      "updated_at": "string",
      "user_id": "string"
    }
  ],
  "limit": 60,
  "page": 1,
  "search": "keyword",
  "total_page": 1,
  "total_row": 1
}
```

### 5. /v1/books/:book_id [GET]
This API for getting the detail of the book data.
```python
Parameter
book_id string => For getting the id book

Response : Code 200 OK
{
  "data": {
    "author": "string",
    "book_ownership": true,
    "created_at": "string",
    "description": "string",
    "edition": "string",
    "format": "string",
    "genre": "string",
    "id": "string",
    "image_url": "string",
    "isbn": "string",
    "language": "string",
    "number_of_pages": 0,
    "price": 0,
    "publication_date": "string",
    "publisher": "string",
    "title": "string",
    "updated_at": "string",
    "user_id": "string"
  }
}
```

### 6. /v1/books/:book_id [PUT]
This API for updating the book data.
```python
Parameter
book_id string => For getting the id book

Payload
{
  "author": "string",
  "description": "string",
  "edition": "string",
  "format": "string",
  "genre": "string",
  "image_url": "string",
  "isbn": "string",
  "language": "string",
  "number_of_pages": 0,
  "price": 0,
  "publication_date": "string",
  "publisher": "string",
  "title": "string",
  "user_id": "string"
}

Response : Code 200 OK
```

### 7. /v1/books/:book_id [DELETE]
This API for deleting the book data that related to the ownership of the login user.
```python
Parameter
book_id string => For getting the id book

Response : Code 204 No Content
```
