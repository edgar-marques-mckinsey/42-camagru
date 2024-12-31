up:
	docker-compose up -d

down:
	docker-compose down --rmi all --volumes --remove-orphans

re: down up

.PHONY: up down re
