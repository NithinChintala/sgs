package model

import (
	"strconv"
	"log"
)

func nilifyStr(str string) *string {
	if str == "<nil>" || str == "" {
		return nil
	}
	return &str
}

func mustAtoi(str string) int {
	conv, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return conv
}

func nilifyInt(str string) *int {
	if str == "<nil>" || str == "" {
		return nil
	}
	conv := mustAtoi(str)
	return &conv
}