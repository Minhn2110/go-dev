package main

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func main() {
	err := errors.New("this is a sample error")
	logrus.WithError(err).Error("Something went wrong")
	//logrus.Error("Get Import By Id Failed", err)

}
