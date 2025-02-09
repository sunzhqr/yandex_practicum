package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logrus.SetOutput(file)
	// устанавливаем вывод логов в формате JSON
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.WarnLevel)

	logrus.WithFields(logrus.Fields{
		"genre":  "pop",
		"singer": "The Weeknd",
	}).Info("The king of modern pop industry")

	logrus.WithFields(logrus.Fields{
		"omg":  true,
		"name": "Hurry Up Tomorrow",
	}).Warn("Hurry Up Tomorrow is the final installment of a trilogy following the Weeknd's previous two studio albums.")

	logrus.WithFields(logrus.Fields{
		"omg":  true,
		"name": "After Hours",
	}).Fatal("But the After Hours album the highest rated in media")
}
