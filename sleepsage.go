package main

import (
	"os"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

var (
	wakeUpMode bool
	strTime    string
	sleepTime  time.Time
	err        error
	logger     *log.Logger
)

func main() {
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "SleepSage💤",
	})

	huh.NewSelect[bool]().
		Title("SleepSage - Mode").
		Options(
			huh.NewOption("When should I wake up?", true),
			huh.NewOption("When should I go to sleep?", false),
		).
		Value(&wakeUpMode).Run()

	if wakeUpMode {
		huh.NewInput().
			Title("When do you want to go to sleep?").
			Placeholder("HH:MM").
			Validate(validateTime).
			Value(&strTime).Run()

		wakeUpTimes := calculateWakeUpTime(sleepTime)

		logger.Infof("You should try to waking up at one of the following times to feel refreshed:")
		for _, wakeUpTime := range wakeUpTimes {
			if wakeUpTime == wakeUpTimes[0] {
				logger.SetPrefix("🟡")
			} else {
				logger.SetPrefix("🟢")
			}
			logger.Infof(wakeUpTime)
		}
	} else {
		huh.NewInput().
			Title("When do you want to wake up?").
			Placeholder("HH:MM").
			Validate(validateTime).
			Value(&strTime).Run()

		sleepTimes := calculateSleepTime(sleepTime)

		logger.Infof("You should try to go to sleep at one of the following times to feel refreshed:")
		for _, sleepTime := range sleepTimes {
			if sleepTime == sleepTimes[0] {
				logger.SetPrefix("🟡")
			} else {
				logger.SetPrefix("🟢")
			}
			logger.Infof(sleepTime)
		}
	}
}

func validateTime(str string) error {
	sleepTime, err = time.Parse("15:04", strTime)
	if err != nil {
		return err
	}

	hour, min, _ := sleepTime.Clock()
	sleepTime = time.Date(0, 0, 0, hour, min, 0, 0, time.UTC)

	return nil
}

func calculateWakeUpTime(sleepTime time.Time) []string {
	sleepCycleDuration := 90 * time.Minute

	sleepTime = sleepTime.Add(20 * time.Minute)

	wakeUpTimes := make([]string, 3)

	for i := 4; i <= 6; i++ {
		wakeUpTime := sleepTime.Add(sleepCycleDuration * time.Duration(i))
		wakeUpTimes[i-4] = wakeUpTime.Format("15:04")
	}

	return wakeUpTimes
}

func calculateSleepTime(wakeUpTime time.Time) []string {
	sleepCycleDuration := 90 * time.Minute

	wakeUpTime = wakeUpTime.Add(-20 * time.Minute)

	sleepTimes := make([]string, 3)

	for i := 4; i <= 6; i++ {
		sleepTime := wakeUpTime.Add(-sleepCycleDuration * time.Duration(i))
		sleepTimes[i-4] = sleepTime.Format("15:04")
	}

	return sleepTimes
}
