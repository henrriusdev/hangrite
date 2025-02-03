# Run templ generation in watch mode
templ:
	templ generate --watch -v

# Run air for Go hot reload
server:
	air -c .air.toml

# Watch Tailwind CSS changes
tailwind:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

# Start development server with all watchers
dev:
	make -j3 templ server tailwind
