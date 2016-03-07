package main

import (
	"fmt"
	"strings"
)

func parseImage(img string) (Image, error) {
	switch {
	case strings.HasPrefix(img, dockerPrefix):
		return parseDockerImage(strings.TrimPrefix(img, dockerPrefix))
		//case strings.HasPrefix(img, appcPrefix):
		//
	}
	return nil, fmt.Errorf("no valid prefix provided")
}
