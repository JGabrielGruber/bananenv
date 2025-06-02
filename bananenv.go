package main

import (
	"fmt"
	"os"
	"strings"
)

const envFileVar = "BANANENV_FILE"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bananenv {set KEY=VALUE|unset KEY|list|init bash}")
		os.Exit(1)
	}

	cmd := os.Args[1]
	envFile := getEnvFile()

	switch cmd {
	case "set":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bananenv set KEY=VALUE [KEY=VALUE...]")
			os.Exit(1)
		}
		envs := loadEnvs(envFile)
		for _, pair := range os.Args[2:] {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "Invalid pair: %s\n", pair)
				os.Exit(1)
			}
			envs[parts[0]] = parts[1]
			fmt.Printf("export %s=%q\n", parts[0], parts[1])
		}
		saveEnvs(envFile, envs)
	case "unset":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bananenv unset KEY [KEY...]")
			os.Exit(1)
		}
		envs := loadEnvs(envFile)
		for _, key := range os.Args[2:] {
			delete(envs, key)
			fmt.Printf("unset %s\n", key)
		}
		saveEnvs(envFile, envs)
	case "list":
		envs := loadEnvs(envFile)
		for k, v := range envs {
			fmt.Printf("export %s=\"%s\"\n", k, v)
		}
	case "init":
		if len(os.Args) != 3 || os.Args[2] != "bash" {
			fmt.Println("Usage: bananenv init bash")
			os.Exit(1)
		}
		fmt.Print(generateInitScript(envFile))
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

func getEnvFile() string {
	envFile := os.Getenv(envFileVar)
	if envFile == "" {
		// Use a fixed temp file name
		envFile = "/tmp/bananenv.session"
		// Create file if it doesn't exist
		if _, err := os.Stat(envFile); os.IsNotExist(err) {
			file, err := os.Create(envFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create temp file: %v\n", err)
				os.Exit(1)
			}
			file.Close()
			if err := os.Chmod(envFile, 0600); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to set permissions: %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Printf("export %s=%q\n", envFileVar, envFile)
	}
	return envFile
}

func loadEnvs(envFile string) map[string]string {
	envs := make(map[string]string)
	data, err := os.ReadFile(envFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Failed to read env file: %v\n", err)
		os.Exit(1)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "export ") {
			parts := strings.SplitN(line[7:], "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := strings.Trim(parts[1], "\"")
				envs[key] = value
			}
		}
	}
	return envs
}

func saveEnvs(envFile string, envs map[string]string) {
	var lines []string
	for k, v := range envs {
		lines = append(lines, fmt.Sprintf("export %s=\"%s\"", k, v))
	}
	data := strings.Join(lines, "\n")
	if err := os.WriteFile(envFile, []byte(data), 0600); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write env file: %v\n", err)
		os.Exit(1)
	}
}

func generateInitScript(envFile string) string {
	return fmt.Sprintf(`
# bananenv initialization
export %s=%q
BANANENV_EXEC_READY="true"
__bananenv_last_mtime=0

# Pre-exec hook for commands
__bananenv_preexec() {
  local cmdline="$1"
  # Skip invalid or internal commands
  if [ "${BANANENV_EXEC_READY}" != "true" ] || [ -z "$cmdline" ] || [[ "$cmdline" =~ ^[[:space:]]*(DEBUG|trap|eval|source) ]] || [[ "$cmdline" =~ "\033" ]]; then
    return 0
  fi
  BANANENV_EXEC_READY="false"
  # Check if env file changed
  if [ -f "$%s" ]; then
    mtime=$(stat -c %%Y "$%s" 2>/dev/null || echo 0)
    if [ "$mtime" != "$__bananenv_last_mtime" ]; then
      if [ "$(stat -c %%U "$%s" 2>/dev/null)" = "$USER" ]; then
        source "$%s"
        __bananenv_last_mtime=$mtime
      fi
    fi
  fi
  # Handle @ commands
  if [[ "$cmdline" =~ ^@([[:alnum:]]+)(.*)$ ]]; then
    cmd="${BASH_REMATCH[1]}"
    args="${BASH_REMATCH[2]}"
    args=$(echo "$args" | xargs)
    if [ -n "$args" ]; then
      eval "bananenv \"$cmd\" $args"
    else
      bananenv "$cmd"
    fi
    return 127  # Signal command was handled
  fi
  # Non-@ commands are handled by Bash
  return 0
}

# Command not found handler for @ commands
command_not_found_handle() {
  if [[ "$1" =~ ^@([[:alnum:]]+)$ ]]; then
    # Silently ignore @ commands (already handled by trap)
    return 127
  fi
  echo "bash: $1: command not found" >&2
  return 127
}

# Pre-command hook for prompt
__bananenv_precmd() {
  BANANENV_EXEC_READY="true"
  # Run existing PROMPT_COMMAND if set
  if [ -n "${BANANENV_ORIG_PROMPT_COMMAND}" ]; then
    eval "${BANANENV_ORIG_PROMPT_COMMAND}"
  fi
}

# Set PROMPT_COMMAND
if [ -z "${PROMPT_COMMAND}" ]; then
  PROMPT_COMMAND="__bananenv_precmd"
else
  BANANENV_ORIG_PROMPT_COMMAND="$PROMPT_COMMAND"
  PROMPT_COMMAND="__bananenv_precmd"
fi

# Set DEBUG trap
trap '__bananenv_preexec "$BASH_COMMAND"' DEBUG
`, envFileVar, envFile, envFileVar, envFileVar, envFileVar, envFileVar)
}
