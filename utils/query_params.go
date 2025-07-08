package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type QueryParams struct {
	Page     int
	Limit    int
	Sort     string
	Order    int
	Filter   string
	Search   string
}

func ParseQueryParams(c *fiber.Ctx) QueryParams {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	order := -1 
	if c.Query("order", "desc") == "asc" {
		order = 1
	}

	return QueryParams{
		Page:   page,
		Limit:  limit,
		Sort:   c.Query("sort", "created_at"),
		Order:  order,
		Filter: c.Query("filter", ""),
		Search: c.Query("search", ""),
	}
}
