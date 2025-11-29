package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	sliceData := strings.SplitN(data, ",", 3)
	if len(sliceData) < 3 {
		return 0, "", 0, errors.New("")
	}
	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("")
	}
	duration, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("")
	}
	typeOfActivity := sliceData[1]
	return steps, typeOfActivity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	averageSpeed := distance / (duration.Hours())
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeOfActivity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch typeOfActivity {
	case "Бег":
		distance := distance(steps, height)
		averageSpeed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", duration.Hours(), distance, averageSpeed, calories), nil
	case "Ходьба":
		distance := distance(steps, height)
		averageSpeed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", duration.Hours(), distance, averageSpeed, calories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (weight * averageSpeed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := ((weight * averageSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
