Spaced repetition TUI application written in Go (bubbletea) in order to practice leetcode. Grabs leetcode problem information via their GraphQL API. The user is able to sign in via an assisted interface, likely through opening a browser because LeetCode has CAPTCHA. And then the user will be able to see their own problems solved, problems they've solved in the past, that sort of thing, in the application. If the user does not authenticate, then uh they're still able to use the space repetition algorithm, but it'll just be tracked locally.

```Project Structure
├── configs
├── go.mod
├── go.sum
├── internal
│   ├── API
│   │   ├── api.go
│   │   └── config.go
│   ├── daileet
│   │   └── queries.go
│   ├── data
│   │   ├── functions.go
│   │   ├── GQLQueries
│   │   └── patterns.go
│   ├── srs
│   │   └── scheduler.go
│   ├── storage
│   │   └── db.go
│   └── tui
│       └── app.go
├── pkg
├── _test
│   └── api_test.py
└── tests
```

## Log:
- setup primary project structure
- found LeetcodeAPI examples for proper queries from the site
- started writing API client

## TODO:
- setup API wrapper around graphQL queries (keeping very thin, just post requests over HTTP)
- return data necessary to track problems within a certain predefined database (blind 75, etc)
- setup basic tui operations
- generate logo
- write SRS algorithm (research online)