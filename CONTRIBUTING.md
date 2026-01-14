# Contributing to Goway

First off, thanks for taking the time to contribute! 🎉

## Philosophy

We strive to keep the codebase clean, readable, and maintainable. We follow **Hexagonal Architecture**, **Clean Architecture**, and **Domain-Driven Design** strictly. When contributing, please ensure you respect the boundaries between layers.

- **Don't** import `adapter` code into `domain` or `application`.
- **Don't** put business logic in the HTTP handlers.
- **Do** test your business logic.

## How to Contribute

1. **Fork the repository.**
2. **Create a feature branch.** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes.** (`git commit -m 'Add some amazing feature'`)
4. **Push to the branch.** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request.**

## Coding Standards

- **Formatting**: Run `go fmt ./...` before committing.
- **Linting**: We recommend using `golangci-lint`.
- **Conventions**:
    - Use meaningful variable names.
    - Handle errors explicitly.
    - Keep functions small and focused.

## Reporting Bugs

If you find a bug, please look through existing issues to see if it has already been reported. If not, open a new issue with a clear description of the problem and steps to reproduce it.
