package expenses

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetexpensesByIdHandler(c echo.Context) error {
	id := c.Param("id")

	rowID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "id should be int " + err.Error()})
	}

	row := db.QueryRow("SELECT id, title, amount, note,tags FROM expenses WHERE id=$1", rowID)

	exp := expense{}
	err = row.Scan(&exp.ID, &exp.Title, &exp.Amount, &exp.NOTE, pq.Array(&exp.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	fmt.Printf("exp % #v\n", exp)

	return c.JSON(http.StatusOK, exp)
}

func GetExpenseHandler(c echo.Context) error {
	exp := []expense{}

	rows, err := db.Query("SELECT id, title, amount, note,tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	for rows.Next() {
		e := expense{}
		err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.NOTE, pq.Array(&e.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})

		}
		exp = append(exp, e)
	}

	return c.JSON(http.StatusOK, exp)
}
