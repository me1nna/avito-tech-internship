## Description:

A small REST service. Providing methods to interact with user's money accounts,
for example crediting money and debiting(by creating service order) money, 
getting user's transaction history, getting "income-per-service" report.

## How to start:
```bash
$ docker-compose up
```
This command will create two containers:
* PostreSQL on 54320 port
* App on 8080 port
