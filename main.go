// Copyright 2023 Blink Labs Software
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleWebhook(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.JSON(
			http.StatusUnsupportedMediaType,
			"invalid request body, should be application/json",
		)
		return
	}
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, "empty request")
		return
	}
	rawBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to read request body")
		return
	}
	_ = c.Request.Body.Close()
	// fmt.Printf("received webhook payload: %s\n", string(rawBytes))
	fmt.Println(string(rawBytes))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.New()
	router.POST("/", handleWebhook)
	_ = router.Run()
	// fmt.Println("started webhook server")
	select {}
}
