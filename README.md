<h1 align="center">🎬 go_IMDB-cli</h1>

<p align="center">
  <b>A powerful command-line tool to search, compare, and manage your favorite movies using the OMDb API and SQLite — built in Go.</b>
</p>

<p align="center">
  <a href="https://golang.org"><img alt="Go" src="https://img.shields.io/badge/Go-1.18%2B-blue" /></a>
  <a href="LICENSE"><img alt="MIT License" src="https://img.shields.io/badge/License-MIT-green.svg" /></a>
  <img alt="Build" src="https://img.shields.io/badge/build-passing-brightgreen" />
</p>

---

## ✨ Features

- 🔍 **Search** movies by title
- 📖 **View full details** using IMDb ID
- ➕ **Add** movies to your favorites (saved locally)
- ⭐ **List** all your favorite movies
- 🗑️ **Delete** a movie from favorites
- ⚖️ **Compare** two movies by IMDb rating
- 💾 Local **SQLite** storage (via GORM)

---

## 🚀 Quick Start

### ⚙️ Option 1: Run as Binary (Recommended)

```bash
git clone https://github.com/sidharth-chauhan/go_IMDB-cli.git
cd go_IMDB-cli
go build -o imdb
./imdb --help
```

### 🐳 Option 2: Run with Docker Compose

```bash
# Build and start the CLI container
docker-compose up
```

Then in a separate terminal window, run commands like:

```bash
docker-compose run imdb search --title "Inception"
docker-compose run imdb add --id tt1375666
docker-compose run imdb list
docker-compose run imdb delete --id tt1375666
docker-compose run imdb compare --id1 tt0111161 --id2 tt0068646
```

To stop and clean up:

```bash
docker-compose down
```

---

## 🛠️ Requirements

| Requirement         | Description                                 |
|---------------------|---------------------------------------------|
| 🧠 Go Version        | Go 1.18 or newer                             |
| 🌍 Internet          | Needed for OMDb API calls                    |
| 🔑 OMDb API Key      | Already set in code (`cad09794`)             |

---

## 📝 Usage Examples

### 🔧 Binary

```bash
# 🔍 Search for a movie
./imdb search --title "Inception"

# 📖 Get full movie details
./imdb detail --id tt1375666

# ➕ Add a movie to favorites
./imdb add --id tt1375666

# ⭐ List all favorite movies
./imdb list

# 🗑️ Delete a favorite
./imdb delete --id tt1375666

# ⚖️ Compare two movies by IMDb rating
./imdb compare --id1 tt0111161 --id2 tt0068646
```

### 🐳 Docker Compose

```bash
# 🔍 Search for a movie
docker-compose run imdb search --title "Inception"

# 📖 Get full movie details
docker-compose run imdb detail --id tt1375666

# ➕ Add a movie to favorites
docker-compose run imdb add --id tt1375666

# ⭐ List all favorite movies
docker-compose run imdb list

# 🗑️ Delete a favorite
docker-compose run imdb delete --id tt1375666

# ⚖️ Compare two movies by IMDb rating
docker-compose run imdb compare --id1 tt0111161 --id2 tt0068646
```

---

## 📚 Command Reference

| Command    | Description                                 | Example                                      |
|------------|---------------------------------------------|----------------------------------------------|
| `search`   | Search movies by title                      | `./imdb search --title "Batman"`             |
| `detail`   | View detailed info by IMDb ID               | `./imdb detail --id tt1234567`               |
| `add`      | Save a movie to favorites                   | `./imdb add --id tt1234567`                  |
| `list`     | List all favorite movies                    | `./imdb list`                                |
| `delete`   | Delete a movie from favorites               | `./imdb delete --id tt1234567`               |
| `compare`  | Compare two movies by IMDb rating           | `./imdb compare --id1 tt0111161 --id2 tt0068646` |

---

## 📂 Project Structure

```
go_IMDB-cli/
├── cmd/         # CLI command logic (search, add, list, etc.)
├── db/          # Database connection and helpers
├── models/      # Structs for API and database models
├── utils/       # Functions for API calls and helpers
├── test/        # Test cases (WIP)
├── main.go      # Entry point
├── Dockerfile   # Build CLI inside container
├── docker-compose.yml # Compose file for volume + reuse
└── Makefile     # Pretty test output
```

---

## 🧪 Running Tests

```bash
go test -v ./test/...
```

> ⚠️ Note: Tests require internet access and use `data/favorites.db`.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
