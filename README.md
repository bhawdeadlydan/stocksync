# Stock Info service


## Pre requisites

- Docker
- Golang v1.14+
 
 
## Running App 

1. Export DB_PASSWORD
`export DB_PASSWORD="S3cretP@ssw0rd"` 

2. Bring up the mysql container using:

`make infra-local`

3. Run the migrations on the local DB 
 
`make setup`

4. Build and run the app container.

`make app`

5. Inspect logs using docker 

`docker logs stock-service-go -f`

## Verifying the Functionality

Create Stocks Request
```
curl -X POST \
  http://localhost:8888/stock \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
"fsym": "BTC",
"tsym": "USD",
"CHANGE24HOUR": "-13.25",
"CHANGEPCT24HOUR": "-0.18152873223073468",
"OPEN24HOUR": "7299.12",
"VOLUME24HOUR": "47600.120073200706",
"VOLUME24HOURTO": "348033250.4911315",
"LOW24HOUR": "7197.22",
"HIGH24HOUR": "7426.64",
"PRICE": "7285.87",
"SUPPLY": "18313937",
"MKTCAP": "133432964170.19"
}'
```


