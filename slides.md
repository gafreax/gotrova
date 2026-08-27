---
marp: true
theme: default
paginate: true
style: |
  @import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;700&display=swap');

  @import url('https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap')

  :root {
    --font-family: 'Inter', sans-serif;
    --font-size: 20px;
  }
    
  body {
    font-family: var(--font-family);
    font-size: var(--font-size);
  }

  section {
    background-color: #2C001E; /* Ubuntu Dark Aubergine */
    color: #F8FAFC;
    font-size: 30px;
  }
  
  h1 {
    color: #E95420; /* Ubuntu Orange */
    font-size: 60px;
    margin-bottom: 0.5em;
  }
  
  h2, h3 {
    color: #B14AE0; /* Vibrant Purple Accent */
  }

  strong {
    color: #E95420; /* Orange highlights */
  }
  
  em {
    color: #772953; /* Light Aubergine */
    font-style: normal;
  }
  
  /* Bubble Tea specific styling (Rounded borders, terminal-like) */
  pre {
    background-color: #1a0011;
    border-radius: 12px;
    border: 2px solid #5E2750; /* Mid Aubergine */
    padding: 20px;
    font-size: 24px;
    font-family: 'JetBrains Mono', monospace;
  }
  
  code {
    font-family: 'JetBrains Mono', monospace;
    color: #E95420;
    background-color: rgba(233, 84, 32, 0.1);
    padding: 8px;
    border-radius: 4px;
    margin: 4px;
  }

  /* Fancy accent bar at the top */
  section::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 8px;
    background: linear-gradient(90deg, #E95420, #B14AE0, #77216F);
  }
---

# 🐹 gotrova

**The Ultimate Go Package Explorer**

A highly interactive Terminal User Interface (TUI) for `pkg.go.dev`.

---

## 🚀 Why gotrova?

As Go developers, we constantly search for packages. 

Switching context to a web browser breaks flow. We wanted something:
- **Fast:** Instant feedback and search results.
- **Terminal-Native:** Never leave your IDE or terminal window.
- **Beautiful:** Built with the aesthetics of modern TUIs.

*Inspired by the `@gafreax/trova` npm package.*

---

## 🛠️ Tech Stack

Built on the shoulders of giants in the Go ecosystem:

- **Bubble Tea**: State machine & interactive TUI framework (`charmbracelet`)
- **Bubbles**: Reusable UI components (Spinner, Table, Viewport)
- **Lipgloss**: For that sweet, sweet terminal CSS styling
- **Standard Library**: Clean API fetching using only `net/http` and `regexp`

---

## 💻 How it Works (1/2)

Start your search right from the command line:

```bash
$ ./gotrova
```

Or pass your query directly for instant results:

```bash
$ ./gotrova bubbletea
```

---

## 💻 How it Works (2/2)

1. **Search:** `gotrova` scrapes `pkg.go.dev`, parsing results instantly.
2. **Sort:** It intelligently prioritizes matches in the package *Name* and *Path* over generic descriptions.
3. **View:** Results are formatted in a clean, scrollable **Table**.

---

## 📦 Deep Dive & Install

Select a package from the table and hit `Enter`. 

`gotrova` smoothly transitions into a **Details Card** (using the Viewport bubble) offering plenty of space to read the full synopsis.

Ready to use it? Just press `i`.

```bash
# gotrova runs this for you in the background:
go get <package_path>
```

You'll see a beautiful loading spinner until the installation completes!

---

## 💜 Get Started Today

Available now on GitHub. Install it instantly anywhere you have Go:

```bash
# Install globally
go install github.com/gafreax/gotrova/cmd/gotrova@latest

# Run!
gotrova
```

*Built with ♥️ using Go and Bubble Tea.*
