
> go get github.com/joho/godotenv
> go get -u gorm.io/gorm
> go get -u gorm.io/driver/postgres
> go get -u github.com/gorilla/mux

> docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=quests -d postgres