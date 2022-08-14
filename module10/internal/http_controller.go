package internal

import (
	"cncamp/module10/framework"
	"fmt"
	"net/http"
)

func HttpController(c *framework.Context) {

	req, err := http.NewRequest("GET", "http://simple-web-1", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	client := &http.Client{}
	if _, err := client.Do(req); err != nil {
		fmt.Printf("%s", err)
	}

	c.SetOkStatus().Json("ok, HttpController")
}
