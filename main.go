package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type configFile struct {
	Default string            `yaml:"default"`
	Links   map[string]string `yaml:"links"`
}

var config configFile

func init() {
	fileData, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	if err := yaml.Unmarshal(fileData, &config); err != nil {
		log.Fatalf("failed to parse file: %v", err)
	}
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect(config.Default)
		},
	})

	app.Get("/:link", func(c *fiber.Ctx) error {
		if link, ok := config.Links[c.Params("link")]; ok {
			return c.Redirect(link, fiber.StatusMovedPermanently)
		}

		return c.Redirect(config.Default)
	})

	app.Listen(":3000")
}
