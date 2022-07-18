startPostgres:
	sudo docker container start postgres

createDB:
	sudo docker exec -it postgres createdb --username=postgres --owner=postgres gorm

dropDB:
	sudo docker exec -it postgres dropdb --username=postgres gorm

migrateUp:
	migrate -path db/migrations -database "postgresql://postgres:StrongPassw0rd@172.17.0.2:5432/gorm?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migrations -database "postgresql://postgres:StrongPassw0rd@172.17.0.2:5432/gorm?sslmode=disable" -verbose down

buildUp:	startPostgres createDB migrateUp

