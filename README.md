# 🐹 gotrova

**The Ultimate Go Package Explorer**

`gotrova` is a highly interactive Terminal User Interface (TUI) for exploring and installing Go packages from `pkg.go.dev` right from your terminal. Built with the beautiful aesthetics of [Bubble Tea](https://github.com/charmbracelet/bubbletea) and inspired by the `@gafreax/trova` npm package.

## 🚀 Features

- **Fast & Interactive:** Search for Go packages with instant feedback and a clean TUI without ever leaving your terminal.
- **Beautiful UI:** Leveraging Bubble Tea, Bubbles, and Lipgloss for a stunning terminal experience.
- **Deep Dive:** View full package descriptions and details in a clean, scrollable interface.
- **One-Key Install:** Found what you need? Just press `i` on a package to install it directly to your project in the background.

## 📦 Installation

To install `gotrova` globally, make sure you have Go installed, then run:

```bash
go install github.com/gafreax/gotrova/cmd/gotrova@latest
```

Ensure that your `$(go env GOPATH)/bin` directory is in your system's `PATH`.

Alternatively, you can build it from source:

```bash
git clone https://github.com/gafreax/gotrova.git
cd gotrova
make build
# The binary will be available as ./gotrova
```

## 💻 Usage

You can start the interactive TUI directly, or pass a search query to get instant results.

```bash
# Launch the interactive explorer
gotrova

# Or start searching immediately
gotrova bubbletea
```

### Controls

- **Navigate:** Use the `Up`/`Down` arrow keys or `j`/`k` to scroll through the search results table.
- **Details:** Press `Enter` on a package to open the Details Card and read the full package synopsis.
- **Install:** While viewing a package, press `i` to automatically run `go get <package_path>` and install it.
- **Back/Quit:** Press `Esc` to go back from the details view, and `q` or `Ctrl+C` to exit the application.

## 🛠️ Tech Stack

Built on the shoulders of giants in the Go ecosystem:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): State machine & interactive TUI framework.
- [Bubbles](https://github.com/charmbracelet/bubbles): Reusable UI components (Spinner, Table, Viewport).
- [Lipgloss](https://github.com/charmbracelet/lipgloss): For that sweet, sweet terminal CSS styling.
