package api

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io"
	// "net/http"
	queries "github.com/0b10headedcalf/daileet/internal/API/Queries"
)

type GQLRequest struct {
	Query     string
	Variables map[string]any
}

func api() {
	// client := &http.Client{}
}

func test() {
	fmt.Println(queries.TestQ)
}

func main() {
	test()
}
