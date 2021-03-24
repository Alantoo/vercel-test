package api

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r* http.Request) {
	now := time.Now().Format(time.RFC3339Nano)
	fmt.Fprintf(w, "OK " + now)
}
