hello:
   echo "Hello, World"
   echo "This line will print if the file hello does not exist."

migrate-main-up:
   @migrate -database "postgresql\://postgres\:postgres@localhost\:5432/fetroshop?sslmode=disable" -path db/migrations up