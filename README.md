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

```bash
git clone https://github.com/sidharth-chauhan/imdb-cli.git
cd imdb-cli
go build -o imdb
./imdb --help
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
└── Makefile     # For running tests with pretty output
```

---

## 🧪 Running Tests

```bash
# Run test cases with detailed output
go test -v ./test/...
```

> Note: Tests require internet access for OMDb API and use a local SQLite test database (`data/favorites.db`).

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---
