package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed to run:", err)
		os.Exit(1)
	}
}

func run() error {
	var err error
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("usage: countdown [minutes]")
	}

	var countdown time.Time

	val, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("atoi %s: %w", args[1], err)
	}
	countdown = time.Now().Add(time.Duration(val) * time.Minute)

	err = update(countdown)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}

	ticker := time.NewTicker(1 * time.Second)
	for {
		<-ticker.C
		err = update(countdown)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}
	}
}

func update(countdown time.Time) error {
	w, err := os.Create("timer.txt")
	if err != nil {
		return err
	}
	defer w.Close()

	if countdown.Before(time.Now()) {
		_, err = w.WriteString("Starting Now")
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		os.Exit(0)
	}
	total := int(time.Until(countdown).Seconds())
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	if minutes >= 1.0 {
		_, err = w.WriteString(fmt.Sprintf("%d:%d", minutes, seconds))
		if err != nil {
			return fmt.Errorf("write minutes: %w", err)
		}
		return nil
	}
	_, err = w.WriteString(fmt.Sprintf("%d", seconds))
	if err != nil {
		return fmt.Errorf("write seconds: %w", err)
	}
	return nil
}
