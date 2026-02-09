package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

type Humano struct {
    Nome string `json:"nome"`
    Idade int   `json:"idade"`
}

func main() {

    e := echo.New()

    // GET

	e.GET("/humanos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Humano{})
	})

    // POST

    e.POST("/humanos", func(c echo.Context) error {
        x := new(Humano)
        if err := c.Bind(x); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"erro": "JSON inválido"})
        }
        
        return c.JSON(http.StatusCreated, map[string]interface{}{
            "mensagem": "Humano criado com sucesso!",
            "nome":     x.Nome,
            "idade":    x.Idade,
        })
    })

    // PUT

    e.PUT("/humanos", func(c echo.Context) error {
        x := new(Humano)
        if err := c.Bind(x); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"erro": "JSON inválido"})
        }
        
        return c.JSON(http.StatusOK, map[string]interface{}{
            "mensagem": "Humano atualizado com sucesso!",
            "nome":     x.Nome,
            "idade":    x.Idade,
        })
    })

    // DELETE

    e.DELETE("/humanos", func(c echo.Context) error {
        x := new(Humano)
        if err := c.Bind(x); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"erro": "JSON inválido"})
        }
        
        return c.JSON(http.StatusOK, map[string]interface{}{
            "mensagem": "Humano deletado com sucesso!",
            "nome":     x.Nome,
            "idade":    x.Idade,
        })
    })
e.Logger.Fatal(e.Start(":1323"))
}