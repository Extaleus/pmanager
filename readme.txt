docker run --name=pmanager-db -e POSTGRES_PASSWORD=Gagarin2 -p 5436:5432 -d --rm postgres
migrate -path ./schema -database 'postgres://postgres:Gagarin2@localhost:5436/postgres?sslmode=disable' up