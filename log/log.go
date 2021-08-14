package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func LogRequest(req interface{}, leveltype interface{}) {

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetFormatter(&logrus.JSONFormatter{})
	switch v := leveltype.(type) {
	case error:
		log.WithFields(logrus.Fields{
			"request": req,
			"error": v,
		}).Error("Request")
		return
	}
	log.WithFields(logrus.Fields{
		"request": req,
	}).Info("Request")

}

func LogResponse(req interface{}) {

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"request": req,
	}).Info("Request")

}
