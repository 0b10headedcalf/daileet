<pre align="center">
██████╗  █████╗ ██╗██╗     ███████╗███████╗████████╗
██╔══██╗██╔══██╗██║██║     ██╔════╝██╔════╝╚══██╔══╝
██║  ██║███████║██║██║     █████╗  █████╗     ██║   
██║  ██║██╔══██║██║██║     ██╔══╝  ██╔══╝     ██║   
██████╔╝██║  ██║██║███████╗███████╗███████╗   ██║   
╚═════╝ ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝   
</pre>

<p align="center">
  <strong>A terminal-based spaced-repetition app for LeetCode</strong>
</p> Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea). Track the problems you've solved, review them on an SM-2 schedule, and import your existing LeetCode submissions.

![Go](https://img.shields.io/badge/Built%20with-Go-00ADD8?style=flat&logo=go)

---

## Features

- **Spaced Repetition Scheduling** — Uses the SM-2 algorithm to calculate optimal review intervals based on how well you remember each problem.
- **LeetCode Integration** — Authenticate with your LeetCode session cookie to import problems you've already solved.
- **Problem Presets** — Quickly load curated problem lists (Blind 75, Grind 75) or import your own custom JSON list.
- **Local Persistence** — All data (problems, review history, session cookie) is stored in a local SQLite database (`daileet.db`).

---

## Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.26.2+

### Build from source

```bash
git clone https://github.com/0b10headedcalf/daileet.git
cd daileet
go build -o daileet ./cmd/daileet
```

### Run

```bash
./daileet
```

The first time you run it, a `daileet.db` SQLite file will be created in the project root.

---

## Authentication

To import your solved problems, you need to provide your LeetCode session cookie:

1. From the **main menu**, select **"Log In With LeetCode"**.
2. Press `o` to open LeetCode in your browser.
3. Log in to LeetCode if you aren't already.
4. Open your browser's developer tools → **Application** (or **Storage**) → **Cookies** → `leetcode.com`.
5. Copy the value of the `LEETCODE_SESSION` cookie.
6. Paste it into the terminal prompt (supports bracketed paste) and press **Enter**.

Once saved, the menu will show **"Logged in to LeetCode"**.

---

## Usage

| Option | Description |
|--------|-------------|
| **Due Problems** | Problems scheduled for review today (or overdue). Select one and press **Enter** to review. |
| **See Previously Solved Problems** | All problems with at least one review, sorted by most recent. Shows repetition count and ease factor. |
| **Manually Edit Problem List** | Browse, add (`a`), or delete (`d`) problems. You can add any LeetCode problem by its title slug (e.g. `two-sum`). |
| **Log In With LeetCode** | Paste your `LEETCODE_SESSION` cookie. |
| **Import My Solved Problems** | Fetches every problem you've solved on LeetCode and imports them as reviewed items. |
| **Presets** | Load the Blind 75 or Grind 75 lists, or import a custom JSON file of title slugs. |
| **Clear All User Data** | Wipes all problems and the stored session cookie. Use for testing. |

### Reviewing a Problem

When you select a due problem, the **Review** screen shows:
- Title, difficulty, pattern, and URL
- Last reviewed date, repetition count, and ease factor

Rate your recall with:
- `1` — **Again** (reset interval)
- `2` — **Hard** (short interval)
- `3` — **Good** (standard interval)
- `4` — **Easy** (longer interval)

Press `esc` to cancel, `q` to quit.

---

## Spaced Repetition (SM-2)

Daileet implements the **SuperMemo-2 (SM-2)** algorithm:

- Each problem has an **ease factor** (default 2.5), **repetition count**, and **interval**.
- Grading a review adjusts the ease factor and calculates the next due date.
- Problems start as `new`, become `review` after the first successful review, and stay on schedule forever.


## License

MIT
