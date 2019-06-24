package useragents

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type userAgents []string

func New(filepath string) (uaPointer *userAgents, errDefer error) {
	var ua userAgents
	ua = []string{}

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла %s %s ", filepath, err.Error())
	}

	defer func() {
		if err := jsonFile.Close(); err != nil {
			errDefer = err
		}
	}()

	jsonParser := json.NewDecoder(jsonFile)
	if err := jsonParser.Decode(&ua); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга %s %s ", filepath, err)
	}

	return &ua, nil
}

func (ua userAgents) Random() string {
	if len(ua) == 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	return ua[rand.Intn(len(ua)-1)]
}

func (ua userAgents) IsEmpty() bool {
	return ua.Length() == 0
}

func (ua userAgents) Length() int {
	return len(ua)
}
