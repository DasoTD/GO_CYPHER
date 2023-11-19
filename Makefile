DB_URL=postgresql://root:secret@localhost:5432/cypherdb?sslmode=disable

postgres:
	sudo docker run -d --name cypher -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
creatdb:
	sudo docker exec -it cypher createdb --username=root --owner=root cypherdb

dropdb:
	# sudo docker exec -it cypher dropdb --username=root --owner=root cypherdb
	sudo docker exec -it cypher dropdb cypherdb

connectDB:
	sudo docker exec -it cypher psql -U root

makemigrattion:
	migrate create -ext sql -dir db/migration -seq add_users

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate
server:
	go run main.go
.PHONY: creatdb dropdb postgres migratedown migrateup migratedown1 migrateup1 test server mock docker sqlc
