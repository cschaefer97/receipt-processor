package scoring

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/cschaefer97/receipt-processor/model"
)

func CheckName(name string) int {
	//award one point for each alphanumeric character in retailer name.
	namePoints := 0
	for _, character := range name {
		if unicode.IsLetter(character) || unicode.IsNumber(character) {
			namePoints += 1
		}
	}
	return namePoints
}

func CheckPrice(total string) int {
	//award points based on total price. if no change, award 50 points. if only quarters, award 25.
	points := 0
	price := strings.Split(total, ".")
	cents, err := strconv.Atoi(price[1])
	if err != nil {
		log.Fatal(err)
	}
	if cents == 0 {
		points += 50
	}
	if cents%25 == 0 {
		points += 25
	}
	return points
}

func CheckNumItems(items []model.Item) int {
	//award points based on number of items purchased
	points := len(items) / 2 * 5
	return points
}

func CheckDescription(items []model.Item) int {
	//award points based on number of characters in item description
	points := 0
	for _, item := range items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			pricePoints := item.Price * 0.2
			points += int(math.Ceil(pricePoints))
		}
	}
	return points
}

func CheckDate(date string) int {
	//if data purchased is odd, return 6 points, otherwise 0.
	dateArr := strings.Split(date, "-")

	//split string to receive date, then evaluate if date is an even-numbered day or not.
	day, err := strconv.Atoi(dateArr[2])
	if err != nil {
		log.Fatal(err)
	}
	if day%2 != 0 {
		return 6
	} else {
		return 0
	}
}

func CheckTime(time string) int {
	//if time purchased between 2pm and 4pm, award 10 points, otherwise 0
	timeArr := strings.Split(time, ":")

	//after splitting the value into hour and minutes, cast both to ints to evaluate.
	hour, err := strconv.Atoi(timeArr[0])
	if err != nil {
		log.Fatal(err)
	}
	min, err := strconv.Atoi(timeArr[1])
	if err != nil {
		log.Fatal(err)
	}

	//evaluate if items purchased during correct time of day to receive points.
	if (hour > 14 && hour < 16) || (hour == 14 && min > 0) {
		return 10
	} else {
		return 0
	}
}
