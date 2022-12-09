# jungle
Its a libary for using annotation comments With struct methods

### ```// @jungle:register```
Its a annotation that registers fiber routes from with in struct method that returns the [register.Route](https://github.com/0-Mqix/jungle/blob/c7398dbbdf459efc650094a537152cb35d95f135/src/register/route.go#L5-L10) struct from it
The methods its uses is depended on the structs you pass in the [register.JungleRoutes](https://github.com/0-Mqix/jungle/blob/c7398dbbdf459efc650094a537152cb35d95f135/src/register/register.go#L16-L75) function then it looks at all the methods and its comments
then checks the if its valid and passes it to fiber

### example
```go
package main

import (
	"github.com/0-Mqix/jungle/src/register"
	"github.com/gofiber/fiber/v2"
)

type App struct{}

// @jungle:register
func (a *app) Test() register.Route {
	return register.Route{
		Method:  "GET",
		Path:    "/app/test.:name",
		Handler: a.test,
	}
}

func (a *App) test(c *fiber.Ctx) error {
	return c.SendString("hey thanks for testing, " + c.Params("name", "nameless person"))
}

func main() {
	f := fiber.New()
	app := App{}

	register.JungleRoutes("./", f, true, &app)

	f.Listen(":3000")
}
```
