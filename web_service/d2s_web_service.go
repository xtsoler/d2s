package main

import (
	"encoding/json"
	"fmt"

	//"io/ioutil"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/xtsoler/d2s"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/d2s2json", func(c *fiber.Ctx) error {
		log.Print("received request")
		// Get first file from form field "document":
		file, err := c.FormFile("data")
		if err != nil {
			return err
		}
		log.Print("File received. Size is ", file.Size)
		c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

		//c.SendString("File received. Size is ", file.Size)
		out, _ := os.Open(file.Filename)
		showDebugOutput := false
		char, err := d2s.Parse(out, showDebugOutput)
		if err != nil {
			log.Fatal(err)
		}
		data, err := json.Marshal(char)
		if err != nil {
			log.Fatal(err)
		}
		out.Close()
		e := os.Remove(file.Filename)
		if e != nil {
			log.Fatal(e)
		}

		//log.Print(string(data))
		//c.SendString(string(data))
		// Save file to root directory:
		//return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		return c.SendString(string(data))
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
