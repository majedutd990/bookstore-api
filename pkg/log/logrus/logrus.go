package logrus

import (
	"io"

	"github.com/alecthomas/units"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/majedutd990/bookstore-api/pkg/log"
	"github.com/sirupsen/logrus"
)

type logBundle struct {
	logger *logrus.Logger
}

func New(path, pattern, maxAgeStr, rotationTimeStr, rotationSizeStr string) (log.Logger, error) {

	l := &logBundle{logger: logrus.New()}
	writer, err := getLoggerWriter(path, pattern, maxAgeStr, rotationTimeStr, rotationSizeStr)
	if err != nil {
		return nil, err
	}
	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})
	return l, nil

}

func getLoggerWriter(path, pattern, maxAgeStr, rotationTimeStr, rotationSizeStr string) (io.Writer, error) {
	maxAge, err := str2duration.ParseDuration(maxAgeStr)
	if err != nil {
		return nil, err
	}
	rotationTime, err := str2duration.ParseDuration(rotationTimeStr)
	if err != nil {
		return nil, err
	}
	// converts 20mb in string to actual byte
	rotationSize, err := units.ParseBase2Bytes(rotationSizeStr)
	if err != nil {
		return nil, err
	}

	writer, err := rotators.New(
		path+pattern,
		rotators.WithLinkName(path),
		rotators.WithMaxAge(maxAge),
		rotators.WithRotationTime(rotationTime),
		rotators.WithRotationSize(int64(rotationSize)),
	)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func (l *logBundle) Error(fields log.LogFields) {
	l.logger.WithFields(logrus.Fields{
		"section":  fields.Section,
		"function": fields.Function,
		"params":   fields.Params,
	}).Error(fields.Message)
}
func (l *logBundle) Info(fields log.LogFields) {
	l.logger.WithFields(logrus.Fields{
		"section":  fields.Section,
		"function": fields.Function,
		"params":   fields.Params,
	}).Info(fields.Message)
}
func (l *logBundle) Warning(fields log.LogFields) {
	l.logger.WithFields(logrus.Fields{
		"section":  fields.Section,
		"function": fields.Function,
		"params":   fields.Params,
	}).Warning(fields.Message)
}
