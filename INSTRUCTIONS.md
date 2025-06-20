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

## Testing

- **End-to-end:** Playwright
- **Pattern:** Page Object Pattern for readable, maintainable tests that mimic user interactions.

## Development Practices

- **TDD:** Test-Driven Development is mandatory.
- **Simplicity:** Prioritize simple, clear solutions.
- **Functional Programming:** Prefer pure functions and avoid side effects, even though Go is not a functional language.
