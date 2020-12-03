package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"homework01/service"
)

func Application() ([]byte, error) {
	file, err := service.Service()
	if err != nil {
		return nil, err
	}
	awesome, err := json.Marshal(file)
	if err != nil {
		return nil, errors.Wrap(err, "json.Marshal Error")
	}
	return awesome, nil
}
