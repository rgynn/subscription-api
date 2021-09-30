# Subscription API for a telecom startup

## Requirements
A .env file containing the follow variables:
```
HOST=0.0.0.0
PORT=3000
PTS_URL=http://api.pts.se/PTSNumberService/Pts_Number_Service.svc/json/SearchByNumber
TIMEOUT_CLIENT=5s
TIMEOUT_IDLE=5s
TIMEOUT_READ=5s
TIMEOUT_WRITE=5s
TEST_MSISDN_NUMBER=8-6785500
TEST_OPERATOR_NAME=Tele2 Sverige AB
```

## How to run

1. Make sure .env file is present
2. make run or go run main.go

## Endpoints

```
GET localhost:3000/api/0.1/subscriptions - List subscriptions
POST localhost:3000/api/0.1/subscriptions - Create new subscription
GET localhost:3000/api/0.1/subscriptions/{msidns} - Get subscription based on MSISDN
PUT localhost:3000/api/0.1/subscriptions/{msidns} - Update subscription activation date (if status pending)
POST localhost:3000/api/0.1/subscriptions/8-6785500/toggle_paused - Toggle subscription status paused/active
POST localhost:3000/api/0.1/subscriptions/8-6785500/cancel - Cancel subscription
```

## Curl commands to test api

```
curl 'localhost:3000/api/0.1/subscriptions' -d '{"msisdn": "8-6785500","activate_at": "2021-05-21T00:00:00Z","type": "PBX"}'
curl 'localhost:3000/api/0.1/subscriptions/8-6785500'
curl 'localhost:3000/api/0.1/subscriptions/8-6785500' -XPUT -H 'Content-Type: application/json' -d '{"msisdn": "8-6785500","activate_at": "2021-06-21T01:00:00Z","type": "PBX"}'
curl 'localhost:3000/api/0.1/subscriptions/8-6785500/toggle_paused' -XPOST
curl 'localhost:3000/api/0.1/subscriptions/8-6785500/toggle_paused' -XPOST
curl 'localhost:3000/api/0.1/subscriptions/8-6785500/cancel' -XPOST
```

## What is lacking?
* GraphQL endpoint
* More unit tests
* Integration tests
* Proper documentation, using openapi specs perhaps
* Logging (access logs and business logic logs)
* Metrics endpoint
* Validation for MSISDN number using regex that works with the PTS api
* Proper state machine for the subscription status
* A repository using database/sql and a postgres driver using the context in the Repository implementations for subscription
