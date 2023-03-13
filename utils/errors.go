package utils

import "github.com/sirupsen/logrus"

func HandleError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", msg, err)
	}
}
