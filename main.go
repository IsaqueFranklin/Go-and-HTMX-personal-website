package main

import (
  //"fmt"
  //"log" 
  //"html/template"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
  //"github.com/gofiber/template/html/v2"
)

func main() {

  //Conectando ao MongoDB com MongoDriver para Go.
  clientOptions := options.Client().ApplyURI("")
  
  //Context mongo.
  ctx := context.Background()

  //Conectando ao cliente mongo.
  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal(err)
  }

  //Verificando a conexão com o banco de dados
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Conectado ao MongoDB.")

  //Definindo database e collection
  collection := client.Database("").Collection("") 

  app := fiber.New(fiber.Config{
    Views: html.New("./views", ".html"),
  })

  app.Static("/", "./public", fiber.Static{
    Compress: true,
  })

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
  })

  app.Get("/about", func(c *fiber.Ctx) error {
    return c.Render("about", fiber.Map{})
  })

  app.Get("/hackerman", func(c *fiber.Ctx) error {
    return c.Render("hackerman", fiber.Map{})
  })

  app.Listen(":4000")
}
//bloating go text.
