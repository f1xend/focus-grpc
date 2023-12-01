package controller

import (
	"encoding/json"
	"fmt"
	"github.com/f1xend/focus-grpc/internal/illusionist/domain"
	"net/http"
)

func (c Controller) Show(res http.ResponseWriter, req *http.Request) {
	rabbit := c.show.Rabbit(req.Context())
	result := domain.Animal{
		Name:  "rabbit",
		Color: rabbit.Color,
	}

	err := json.NewEncoder(res).Encode(result)
	if err != nil {
		c.log.Println(fmt.Errorf("err encode json: %w", err))
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
