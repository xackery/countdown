package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var Version string

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed to run:", err)
		os.Exit(1)
	}
}

func run() error {
	var err error
	text := ""
	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter minutes for countdown: ")
		text, err = reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read failed: %w", err)
		}
		text = strings.ReplaceAll(text, "\n", "")
	} else {
		text = os.Args[1]
	}

	var countdown time.Time

	val, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Errorf("atoi %s: %w", text, err)
	}
	countdown = time.Now().Add(time.Duration(val) * time.Minute)

	fmt.Printf("countdown v%s counting down %d minutes\n", Version, val)
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

	display := fmt.Sprintf("%d", seconds)
	if minutes >= 1.0 {
		display = fmt.Sprintf("%d:%d", minutes, seconds)
	}
	_, err = w.WriteString(display)
	if err != nil {
		return fmt.Errorf("write %s: %w", display, err)
	}
	fmt.Println("updated timer.txt:", display)
	return nil
}
