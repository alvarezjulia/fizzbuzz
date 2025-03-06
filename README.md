# FizzBuzz API Service

A RESTful API service that provides a configurable FizzBuzz implementation with request statistics tracking.

## Description

This service implements a customizable version of the classic FizzBuzz problem through a REST API. It allows users to:
- Specify custom divisors and replacement words
- Set a custom limit for the sequence
- Track and retrieve statistics about the most frequently used parameters

## Available commands

To run the application `make run`

To run the tests `make test`

## API Documentation

The API is documented using OpenAPI (Swagger) specification in `api/api.yml`.

You can use online editors like [Swagger Editor](https://editor.swagger.io/) by copying the contents of `api/api.yml`.

## Curl examples

```
curl -X GET http://localhost:8080/api/stats

curl -X POST http://localhost:8080/api/fizzbuzz 
    -H "Content-Type: application/json" 
    -d '{    
            "firstDivisor": 3,
            "secondDivisor": 5,
            "limit": 15,
            "firstWord": "Fizz",
            "secondWord": "Lololo"
        }'
```