package utils

import (
	"regexp"
	"strconv"
	"time"
)


func ParseIdentityCard(identityCard string) (bool, time.Time) {
	regular := "^\\d{6}(\\d{8})\\d{2}(\\d)[0-9X]$"
	reg := regexp.MustCompile(regular)

	result := reg.FindStringSubmatch(identityCard)

	if len(result) != 3 {

	}

	sex, _ := strconv.Atoi(result[2])

	const format = "2006-01-02"

	birthday, _ := time.Parse(format, result[1][:4] + "-" + result[1][4:6] + "-" + result[1][6:])

	return sex % 2 == 1, birthday
}
