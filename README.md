# go-telkomsel-cart


## Description

This repo is to solve Telkomsel Technical Test

## Getting started

### Prerequisite
```
* golang 1.15
* mysql 5.7
```

### Run Locally
```
make dev
```


## Database
![Database Diagram](/docs/database/go_cart_database.png)

## Event Diagram
### User Event
![User Event Diagram](/docs/sequence_diagram/user_event.png)

### Shop Event
![Shop Event Diagram](/docs/sequence_diagram/shop_event_by_user.png)

### Transaction Event
![Update Transaction Event Diagram](/docs/sequence_diagram/update_transaction_by_admin.png)