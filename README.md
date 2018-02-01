# Web Pub/Sub Architecture
Example with simple chat

## ARCHI:
![schema](docs/schema.png)

## Start project:
* Copy env vars in wsb:
```
> $ cp wsb/.env.dist wsb/.env
```

* Start containers:
```
> $ docker-compose up
```
* Create database and setup API:
```
> $ docker-compose exec php composer install
> $ docker-compose exec php bin/console doctrine:schema:create
> $ docker-compose exec php bin/console app:setup
```
* Setup RabbitMq:
```
> $ docker-compose exec rabbitmq sh /etc/rabbitmq/setup.sh
```

* Restart containers:
```
> $ docker-compose up
```

## Run CLI Client:
* Get dependencies:
```
> $ cd client
> $ go get .
```

* Start Chat with user test1:
```
> $ go run main.go auth test1 test1 --api=symfony.dev --ws=localhost:4000
yo <enter>
```

* Start Chat with user test2:
```
> $ go run main.go auth test1 test1 --api=symfony.dev --ws=localhost:4000
<new message>
```
