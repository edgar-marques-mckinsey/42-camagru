up:
	docker-compose up -d

backend-up:
	docker-compose up -d backend

down:
	docker-compose down --rmi all --volumes --remove-orphans

backend-down:
	docker-compose down --rmi all --remove-orphans backend

re: down up

backend-re: backend-down backend-up

.PHONY: up backend-up down backend-down re backend-re
