# bananenv 🦍🍌

A lightweight shell environment manager for setting and syncing environment variables across terminals, with a cyber ape vibe. Set your `WORKDIR` or custom vars like a boss and swing through your shell sessions with ease!

## Features
- **Set and Sync Vars**: Use `@set FOO=BAR` to set variables that persist across terminals.
- **Auto-Navigate**: Automatically `cd` to `WORKDIR` in new terminals.
- **Bash Integration**: Seamless setup with `eval "$(bananenv init bash)"`.
- **Single Temp File**: Vars stored in `/tmp/bananenv.$USER.session`, cleared on reboot.
- **Cyber Ape Aesthetic**: Built with 🦍 strength and 🍌 flair!

## Installation
1. **Prerequisites**:
   - Go (1.18+)
   - Bash
2. **Clone and Build**:
   ```bash
   git clone https://github.com/jgabrielgruber/bananenv.git
   cd bananenv
   go build -o ~/bin/bananenv bananenv.go
   ```
3. **Set Up Bash**:
   Add to `~/.bashrc`:
   ```bash
   export PATH="~/bin:$PATH"
   eval "$(bananenv init bash)"
   cd $WORKDIR # Optional, but the reason for the tool
   ```
   Source it:
   ```bash
   source ~/.bashrc
   ```

## Usage
- Set a variable:
  ```bash
  @set FOO="BAR"
  ```
- List variables:
  ```bash
  @list
  ```
- Unset a variable:
  ```bash
  @unset FOO
  ```
- Set and navigate to a working directory:
  ```bash
  @set WORKDIR=/path/to/project
  ```
- Check vars in a new terminal:
  ```bash
  echo $FOO
  pwd  # Shows WORKDIR
  ```

## Contributing
We love apes and bananums! 🦍🍌 Check out [CONTRIBUTING.md](CONTRIBUTING.md) to join the jungle:
- Submit bug reports or feature requests via [Issues](https://github.com/jgabrielgruber/bananenv/issues).
- Add support for Zsh, Fish, or other shells.
- Improve performance or add new `@` commands.

## License
MIT License. See [LICENSE](LICENSE) for details.

## Roadmap
- Support for Zsh and Fish shells (see [Issues](https://github.com/jgabrielgruber/bananenv/issues)).
- Add `@clear` command to reset all vars.
- Optional persistent storage for vars across reboots.

## Acknowledgments
Built with love by @jgabrielgruber and cyber ape assistant 🐒 @grok. Inspired by Starship’s shell magic.

APES STRONG TOGETHER! 🦍🐒
