ORBIT_ENV=development

up:
	COMPOSE_PROFILES=dev,db,ipfs docker-compose --env-file=.env.${ORBIT_ENV} up

daemon:
	COMPOSE_PROFILES=dev,db,ipfs docker-compose --env-file=.env.${ORBIT_ENV} up -d

run-ipfs-as-daemon:
	COMPOSE_PROFILES=ipfs docker-compose --env-file=.env.${ORBIT_ENV} up --detach

down:
	COMPOSE_PROFILES=dev,db,ipfs docker-compose --env-file=.env.${ORBIT_ENV} down

build:
	COMPOSE_PROFILES=dev,db,ipfs docker-compose --env-file=.env.${ORBIT_ENV} build

cleanup:
	COMPOSE_PROFILES=dev,db,ipfs docker system prune --volumes

ganache:
	docker-compose --env-file=.env.${ORBIT_ENV} run --rm ganache

truffle:
	docker-compose --env-file=.env.${ORBIT_ENV} run --rm truffle bash

geth:
	docker-compose --env-file=.env.${ORBIT_ENV} run --rm geth

web:
	docker-compose --env-file=.env.${ORBIT_ENV} run --rm web

ipfs:
	docker-compose --env-file=.env.${ORBIT_ENV} run --rm ipfs