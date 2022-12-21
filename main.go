package main

import (
	"encoding/json"
	"fmt"
	"github.com/oculius/oculi/v2/common/error"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/pkg/errors"
	"net/http"
)

func main() {
	const (
		green  = "\033[97;42m"
		yellow = "\033[97;33m"
		reset  = "\033[0m"
	)
	fmt.Printf("%s %s %s %s %s\n", green, "haiii", reset, yellow, "hello")
	y := gerr.NewError(errors.New("testing"), "test", http.StatusBadRequest, 456)
	y2 := response.NewOkResponse("hai", "123", 123)
	x, _ := json.Marshal(response.New(y))
	fmt.Println(string(x))
	x, _ = json.Marshal(response.New(y2))
	fmt.Println(string(x))
}
