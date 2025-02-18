up:
	"docker compose up -d"
down:
	"docker compose down"
run:
	"go run cmd/main.go"
migrate-up:
	"go run database/migration/migrate.go"