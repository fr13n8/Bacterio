package utils

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/internal/utils"
)

func Validate(v []string, x map[string]string) (map[string]string, error) {
	for key := range x {
		err := errors.New(color.YellowString(" [!] You should set a %s!", key))
		if !utils.Contains(v, fmt.Sprintf("%s=", key)) {
			return nil, err
		}
		attr := utils.SplitAfterIndex(utils.Find(v, fmt.Sprintf("%s=", key)), '=')
		if len(attr) == 0 {
			return nil, err
		}
		x[key] = attr
	}
	return x, nil
}
