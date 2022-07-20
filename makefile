include .env
export

go_migrate:
	migrate  -path ./migrations  -database mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp"(127.0.0.1:3306)"/${MYSQL_DBNAME} up   