package daysteps

import (
	"errors"
	"fmt"
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
	
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат данных: ожидалось два значения, разделенных запятой")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось распознать количество шагов: %w", err)
	}

	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
	return 0, 0, fmt.Errorf(
	"не удалось распознать продолжительность: %w", err,
		)
	}

	if duration <= 0 {
	return 0, 0, errors.New(
	"продолжительность активности должна быть больше нуля",
		)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка парсинга данных:", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceInMeters := float64(steps) * stepLength
	distanceInKm := distanceInMeters / float64(mInKm)

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка расчета калорий:", err)
		return ""
	}

	result := fmt.Sprintf(
 	"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
 	steps, distanceInKm, calories,
	)

return result
}
