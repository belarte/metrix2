# Project Instructions: metrix2

## Purpose

Create a web application for tracking user-defined metrics. Users can:
- Define new metrics.
- Add values to metrics over time.
- Analyze tracked metrics.

## Tech Stack

- **Language:** Go
- **UI:** HTMX and templ
- **Persistence:** SQLite
- **CLI:** Cobra

## Project Structure

- `cmd/` — CLI entry points (minimal logic)
- `router/` — HTTP routing and server lifecycle (graceful shutdown)
- `handlers/` — HTTP handler functions (add new endpoints here)
- `web/templates/` — UI components (templ)
- `model/` — Domain models and business logic

## How to Run

- Start the server: `go run main.go serve`
- Run all Go tests: `go test ./...`

## Testing

- **End-to-end:** Playwright
- **Pattern:** Page Object Pattern for readable, maintainable tests that mimic user interactions.

## Development Practices

- **TDD:** Test-Driven Development is mandatory. Write tests before implementing features.
- **Simplicity:** Prioritize simple, clear solutions.
- **Functional Programming:** Prefer pure functions and avoid side effects, even though Go is not a functional language.

## Contribution Guidelines

- Keep commits small and focused.
- Add new HTTP endpoints in `handlers/` and register them in `router/`.
- Prefer pure functions and clear, simple code.
- Follow code review and style guidelines.
