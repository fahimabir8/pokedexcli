# Pokedex CLI

A command-line interface for exploring the Pokémon world, catching Pokémon, and managing your own Pokedex. Built with Go, it features efficient caching and a reactive REPL.


[![Go Version](https://img.shields.io/github/go-mod/go-version/fahimabir8/pokedexcli)](https://golang.org/)

## Features

- **Explore Locations**: Discover Pokémon in various regions using the PokeAPI.
- **Catch & Collect**: Test your luck catching Pokémon and add them to your persistent Pokedex.
- **Inspect Pokémon**: View detailed stats and types of Pokémon you've caught.
- **Efficient Caching**: Built-in TTL-based caching to minimize network requests and ensure a smooth experience.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (1.18 or higher recommended)

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/fahimabir8/pokedexcli.git
   cd pokedexcli
   ```

2. Build the executable:
   ```bash
   go build -o pokedex
   ```

3. Run the application:
   ```bash
   ./pokedex
   ```

## Usage

Once started, the Pokedex CLI enters an interactive REPL mode. You can type commands directly into the prompt.

```bash
Pokedex > help
```

### Core Commands

| Command | Description |
| :--- | :--- |
| `map` | List the next 20 location areas. |
| `mapb` | List the previous 20 location areas. |
| `explore {area}` | List all Pokémon found in a specific location area. |
| `catch {pokemon}` | Attempt to catch a Pokémon and add it to your Pokedex. |
| `inspect {pokemon}` | View stats of a Pokémon you have caught. |
| `pokedex` | List all Pokémon currently in your Pokedex. |
| `exit` | Close the Pokedex CLI. |

### Interest
I completed this project from boot.dev. You can find the link to project here -> <a href="https://www.boot.dev/courses/build-pokedex-cli-golang"> Click </a>