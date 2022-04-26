## Clean Gin

Trying to implement clean architecture with gin framework.

#### Environment Variables

| Key           | Value                    | Desc                          |
| ------------- | ------------------------ | ----------------------------- |
| `ServerPort`  | `:5000`                  | Port at which app runs        |
| `Environment` | `development,production` | App running Environment       |
| `LogOutput`   | `./server.log`           | Output Directory to save logs |
| `DBUsername`  | `username`               | Database Username             |
| `DBPassword`  | `password`               | Database Password             |
| `DBHost`      | `0.0.0.0`                | Database Host                 |
| `DBPort`      | `3306`                   | Database Port                 |
| `DBName`      | `test`                   | Database Name                 |
| `JWTSecret`   | `secret`                 | JWT Token Secret key          |

#### Migration Commands

| Command            | Desc                                           |
| -------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

#### Checklist

- [x] Implement Dependency Injection (go-fx)
- [x] Routing (gin web framework)
- [x] Environment Files
- [x] Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- [x] Middlewares (cors)
- [x] Database Setup (mysql)
- [x] Models Setup and Automigrate (gorm)
- [x] Repositories
- [x] Implementing Basic CRUD Operation
- [x] Authentication (JWT)
- [x] Migration
- [x] Dockerize Application with Debugging Support Enabled. Debugger runs at `5002`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application.


### Eform Gateway Checklist
- [x] CreateTransaction
      - Insert:
        - validasi prefix misal harus 5 huruf, validasi date format yyyy-mm-dd
      - fields request body
        - appname
        - data 
          - datanya bentuk json dari request
        - prefix
        - expired date
      - generate reference code
- [x] UpdateToExecuted
      - create index transactionExecuteds and remove index transactions
- [x] Inquiry search by reference number
      - inquery where reference_code and status = Open
      - Note : 
      - [x] querynya baru by ref_code,
- [x] Cronjob update expired date by timestime, Not including this service api, registered on crontab linux
      - search index where documen if expired_date = now , create to transactionExpireds and delete index from transactions
- [x] crontjob remove index reference_sequence
- [x] revisi code to status http
- [ ] Create TDD testing random refcode loop sejuta
- [ ] Create cmd docker
      - [ ] check connection elasticsearch
      - [ ] check connection mysql
- [ ] create endpoint to update env with authorization
### Feature Eform
- index elastic, 
  - logs
  - transaction
    - appname, object, prefix, expired date, reference_code, status
        - status 
            - Open,
            - Expired 
            - Executed

### Elasticsearch Reference
http://www.inanzzz.com/index.php/post/6drl/a-simple-elasticsearch-crud-example-in-golang
https://github.com/codenoid/golang-elasticsearch-crud

