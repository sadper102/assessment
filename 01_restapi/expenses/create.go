package expenses

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type Err struct {
	Message string `json:"message"`
}

func CreateHandler(c echo.Context) error {
	exp := expense{}
	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses(title, amount, note,tags) VALUES($1, $2, $3,$4) RETURNING id;", exp.Title, exp.Amount, exp.NOTE, pq.Array(&exp.Tags))
	err = row.Scan(&exp.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	fmt.Printf("id : % #v\n", exp)

	return c.JSON(http.StatusCreated, exp)
}
