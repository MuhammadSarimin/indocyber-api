package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type CustomLog struct {
	log *logrus.Logger
}

func NewCustomLog() *CustomLog {

	fileName := path() + time.Now().Format("20060102150405") + "_indocyber_api.log"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logger.SetLevel(logrus.InfoLevel)

	multipleOutput := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(multipleOutput)

	return &CustomLog{log: logger}
}

func (c *CustomLog) Info(message string, req, res interface{}) {

	cl := c.log.WithFields(logrus.Fields{})

	if req != nil {
		cl = cl.WithField("request", req)
	}

	if res != nil {
		cl = cl.WithField("response", res)
	}

	cl.Info(message)
}

func (c *CustomLog) Error(err error, description string, req, res interface{}) {

	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])
	str := strings.ReplaceAll(stFormat, "\t", "")
	strs := strings.Split(str, "\n")

	cl := c.log.WithFields(logrus.Fields{
		"error_stack": []string{
			strs[1],
			strs[2],
		},
		"error_cause": err.Error(),
	})

	if req != nil {
		cl = cl.WithField("request", req)
	}

	if res != nil {
		cl = cl.WithField("response", res)
	}

	cl.Error(description)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func path() string {
	p := Env.PathLog
	mode := os.FileMode(0755)

	_, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(p, mode)
			if err != nil {
				log.Fatalf("Failed to mkdir path log: %v", err)
			}
		} else {
			log.Fatalf("Failed to open log file - check path log: %v", err)
		}
	}

	return "./" + p + "/"

}
