# Payment Transaction Program

Payment transaction is a Restful API written in Golang, which simulates a purchase made with card and according to certain rules, approve or not approve the financial transaction.

## Install

To run this program, you will need to:

```sh
$ export GOPATH=$(go env GOPATH)
$ go get github.com/vinixavier/payment_transaction
$ cd $GOPATH/vinixavier/payment_transaction
$ go install
$ $GOPATH/bin/payment_transaction &
```

Once this is done, the program will listen on `localhost:3000`.

## Test

To test this program:

```sh
$ cd $GOPATH/github.com/vinixavier/payment_transaction
$ go test
```

## Using

To make a payment, use the `curl` command or the `postman` application to perform a POST for the API.

Example in Linux with the `curl` command:

```sh
$ curl -X POST -d @examples/input.json http://localhost:3000/payment

{"approved":false,"deniedReasons":["Card is blocked!"]}
```

See examples of request files in the `examples/*.json` directory.

## Payment

### `POST /payment`

Make a payment.

### Request

```json
{
    "account": {
        "cardIsActive": "Boolean",
        "limit": "Number",
        "blacklist": [ "String" ],
        "isWhitelisted": "Boolean"
    },
    "transaction": {
        "merchant": "String",
        "amount": "Number",
        "time": "String"
    },
    "lastTransaction":
        [ <Transaction> ]
    }
}
```

### Response

```json
{
  "approved": "Boolean",
  "newLimit": "Number",
  "deniedReasons": [ "String" ]
}
```

## Payment Rules

This API considers the following rules to be applied at the time of the transaction:

1. The transaction amount should not be above limit.
2. No transaction should be approved when the card is blocked.
3. The first transaction shouldn't be above 90% of the limit.
4. There should not be more than 10 transactions on the same merchant.
5. Merchant blacklist.
6. There should not be more than 3 transactions on a 2 minutes interval

## ToDo

1. Include `application/json` header.