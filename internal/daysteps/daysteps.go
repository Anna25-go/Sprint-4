package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	sliceData := strings.Split(data, ",")
	if len(sliceData) != 2 {
		return 0, 0, errors.New("неверный формат")
	}
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("пройденное растояние должно быть больше 0")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		log.Println(err)
		return ""
	}
	distanceM := float64(steps) * stepLength
	distanceKM := distanceM / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKM, calories)
}
