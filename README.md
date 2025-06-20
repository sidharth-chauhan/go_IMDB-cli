# 🎬 go_IMDB-cli

[![Go](https://img.shields.io/badge/Go-1.18%2B-blue)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)](#)

An interactive and powerful command-line application to **search, compare, and manage your favorite movies** using the **OMDb API** and local **SQLite** database — built in **Go**.

---

## ✨ Features

- 🔍 Search movies by title
- 📖 View full movie details by IMDb ID
- ➕ Add movies to your favorites (saved locally)
- 🗑️ Delete movies from your favorites
- ⭐ List all saved favorite movies
- ⚖️ Compare two movies by IMDb rating
- 💾 Uses SQLite (via GORM) for persistent local storage

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

- Go 1.18 or newer
- Internet connection (for API calls)
- OMDb API key (already set in code as `cad09794`)

---

## 📝 Usage Examples

```bash
# 🔍 Search for a movie
./imdb search --title "Inception"

# 📖 Get full movie details
./imdb detail --id tt1375666

# ➕ Add to favorites
./imdb add --id tt1375666

# ⭐ List all favorites
./imdb list

# 🗑️ Remove from favorites
./imdb delete --id tt1375666

# ⚖️ Compare two movies by rating
./imdb compare --id1 tt0111161 --id2 tt0068646
```

---

## 📚 Command Reference

| Command   | Description                                 | Example                                      |
|-----------|---------------------------------------------|----------------------------------------------|
| `search`  | Search movies by title                      | `./imdb search --title "Inception"`          |
| `detail`  | View detailed info using IMDb ID            | `./imdb detail --id tt1375666`               |
| `add`     | Save a movie to local favorites             | `./imdb add --id tt1375666`                  |
| `list`    | List all favorite movies                    | `./imdb list`                                |
| `delete`  | Remove a movie from favorites               | `./imdb delete --id tt1375666`               |
| `compare` | Compare two movies based on IMDb rating     | `./imdb compare --id1 tt0111161 --id2 tt0068646` |

---

## 📂 Project Structure

```
go_IMDB-cli/
├── cmd/         # All CLI command logic (search, add, list, etc.)
├── db/          # Database connection setup and helper functions
├── models/      # Struct definitions for API and database models
├── utils/       # API interaction logic
├── test/        # Test cases for all features
├── Makefile     # For running tests with pretty output
└── main.go      # CLI entry point
```

---

## 🧪 Running Tests

```bash
make test  # Pretty test output
# or
go test -v ./test/...
```

---

## 🎉 Pretty Test Output Example

```
==============================
🚀 Running all tests for go_IMDB-cli
==============================
=== RUN   TestAddFavoriteValid
✅ --- PASS: TestAddFavoriteValid (0.45s)
=== RUN   TestSearchMovie
✅ --- PASS: TestSearchMovie (0.33s)
...
🎉 All tests passed!
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---
