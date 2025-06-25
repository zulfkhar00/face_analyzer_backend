set -e # Exit early if any commands fail

(
  cd "$(dirname "$0")" # Ensure compile steps are run within the repository directory
  go build -o /tmp/beauty-backend cmd/main.go
)

chmod +x /tmp/beauty-backend

# Run the Go backend
exec /tmp/beauty-backend "$@"
