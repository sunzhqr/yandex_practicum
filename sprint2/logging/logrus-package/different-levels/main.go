package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Musician struct {
	name string
	song string
}

func main() {
	DifferentLevels(logrus.WarnLevel, nil)
}

func DifferentLevels(level logrus.Level, out io.Writer) {
	if out == nil {
		out = os.Stdout
	}
	logrus.SetOutput(out)
	logrus.SetLevel(level)

	logrus.WithFields(logrus.Fields{
		"singer": "Justin Bieber",
		"song":   "WDYM",
	}).Info("Every girls crush in 2015-2016")

	logrus.WithFields(logrus.Fields{
		"singer": "Michael Jackson",
		"song":   "Beat it",
	}).Warn("The GOAT of pop industry")
	musician := &Musician{
		name: "The Weeknd",
		song: "Starboy",
	}
	contextLogger := logrus.WithFields(logrus.Fields{
		"singer": musician.name,
		"song":   musician.song,
	})

	contextLogger.Warn("The weekend best for me")
	contextLogger.Error("The weekend best for everyone")

	logrus.WithFields(logrus.Fields{
		"singer": "Bruno Mars",
		"song":   "Die with a smile",
	}).Fatal("If the world was ending, i will be next to you")
}
