# Contributing to bananenv 🦍🍌

Welcome to the `bananenv` jungle! We’re thrilled you want to join the cyber ape crew. Whether you’re fixing bugs, adding features, or supporting new shells, your contributions make `bananenv` stronger.

## How to Contribute
1. **Fork the Repo**: Click “Fork” on [GitHub](https://github.com/jgabrielgruber/bananenv).
2. **Clone Your Fork**:
   ```bash
   git clone https://github.com/your-username/bananenv.git
   cd bananenv
   ```
3. **Create a Branch**:
   ```bash
   git checkout -b feature/your-feature
   ```
4. **Make Changes**: Hack on `bananenv.go` or docs.
5. **Test**: Build and test locally:
   ```bash
   go build -o bananenv
   ./bananenv init bash
   ```
6. **Commit**: Follow [Conventional Commits](https://www.conventionalcommits.org/):
   ```bash
   git commit -m "feat: add Zsh support"
   ```
7. **Push and PR**: Push to your fork and open a Pull Request.

## Development Setup
- **Requirements**: Go 1.18+, Bash (Zsh/Fish for testing other shells).
- **Build**: `go build -o bananenv bananenv.go`
- **Test**: Run `@set`, `@list`, `@unset` in your shell.

## Code Style
- Keep it simple and ape-friendly.
- Comment complex logic in `bananenv.go`.
- Follow Go conventions (run `gofmt`).

## Issues
- Check [Issues](https://github.com/jgabrielgruber/bananenv/issues) for tasks.
- Tackle Zsh/Fish support or add new `@` commands.
- Report bugs with steps to reproduce.

## Community
- Be kind and respectful (see [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)).
- Join the jungle on GitHub Discussions for ideas.

APES STRONG TOGETHER! 🦍🐒