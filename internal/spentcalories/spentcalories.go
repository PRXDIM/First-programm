package spentcalories

import (
	"time"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	// TODO: реализовать функцию

	parts := strings.Split(data, ",")

	// Проверить, чтобы длина слайса была равна 3
	if len(parts) != 3 {
		return 0, "", 0, errors.New("неверный формат данных: ожидалось три значения, разделенных запятой")
	}

	// Преобразовать количество шагов в тип int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("не удалось распознать количество шагов: %w", err)
	}

	// Количество шагов должно быть больше 0
	if steps <= 0 {
		return 0, "", 0, errors.New("количество шагов должно быть больше нуля")
	}

	activity := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("не удалось распознать продолжительность: %w", err)
	}

	// Продолжительность должна быть больше 0
	if duration <= 0 {
		return 0, "", 0, errors.New("продолжительность активности должна быть больше нуля")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	// Рассчитать длину шага на основе роста
	stepLength := height * stepLengthCoefficient

	// Умножить пройденное количество шагов на длину шага
	distanceInMeters := float64(steps) * stepLength

	// Разделить полученное значение на число метров в километре
	distanceInKm := distanceInMeters / float64(mInKm)

	return distanceInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if duration <= 0 {
		return 0
	}

	// Вычислить дистанцию с помощью функции distance()
	dist := distance(steps, height)

	// Вычислить и вернуть среднюю скорость в часах
	hours := duration.Hours()
	
	return dist / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	// Получить значения из строки данных с помощью функции parseTraining()
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		// Обработать возможные ошибки и вывести их в лог.
		log.Println(err)
		return "", err
	}

	var calories float64
	var calErr error

	// Проверить, какой вид тренировки был передан
	switch activity {
	case "Бег":
		calories, calErr = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, calErr = WalkingSpentCalories(steps, weight, height, duration)
	default:
		// Если был передан неизвестный тип тренировки, вернуть ошибку.
		return "", errors.New("неизвестный тип тренировки")
	}

	// Обработка ошибок, которые могли возникнуть при расчете калорий
	if calErr != nil {
		log.Println(calErr)
		return "", calErr
	}

	// Общие расчеты для дистанции и скорости
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	hours := duration.Hours()

	// Сформировать и вернуть строку в заданном формате
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, hours, dist, speed, calories,
	)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверить входные параметры на корректность
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, errors.New("вес пользователя должен быть больше нуля")
	}
	if height <= 0 {
		return 0, errors.New("рост пользователя должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность активности должна быть больше нуля")
	}

	// Рассчитать среднюю скорость с помощью функции meanSpeed()
	speed := meanSpeed(steps, height, duration)

	// Рассчитать и вернуть количество калорий
	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / float64(minInH)

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверить входные параметры на корректность
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, errors.New("вес пользователя должен быть больше нуля")
	}
	if height <= 0 {
		return 0, errors.New("рост пользователя должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность активности должна быть больше нуля")
	}

	// Рассчитать среднюю скорость с помощью функции meanSpeed()
	speed := meanSpeed(steps, height, duration)

	// Рассчитать количество калорий
	durationInMinutes := duration.Minutes()

	baseCalories := (weight * speed * durationInMinutes) / float64(minInH)

	calories := baseCalories * walkingCaloriesCoefficient

	return calories, nil
}
