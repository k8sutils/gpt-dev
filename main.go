package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	openai "github.com/openai/openai-go/v2"
)

// Define the OpenAI API credentials
var (
	openaiAPIKey = os.Getenv("OPENAI_API_KEY") // Set your OpenAI API key as an environment variable
	modelID      = "text-davinci-002"
)

// Define the response from the OpenAI API
type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Load the templates
	router.LoadHTMLGlob("templates/*")

	// Serve the static assets
	router.Static("/assets", "./assets")

	// Define the route for the home page
	router.GET("/", func(c *gin.Context) {
		// Render the home page template
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Define the route for submitting the form
	router.POST("/", func(c *gin.Context) {
		// Get the user's input from the form
		input := c.PostForm("input")

		// Query the OpenAI API
		client, err := openai.NewClient(openaiAPIKey)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		resp, err := client.Completions.Create(&openai.CompletionRequest{
			Model:      modelID,
			Prompt:     input,
			MaxTokens:  1024,
			Temperature: 0.7,
			N: 1,
			Stop:       "",
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Parse the response and extract the generated code
		code := ""
		for _, choice := range resp.Choices {
			code += choice.Text
		}

		// Render the results page template
		c.HTML(http.StatusOK, "results.html", gin.H{
			"input": input,
			"code":  template.HTML(code),
		})
	})

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
