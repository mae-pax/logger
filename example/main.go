package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	baselogger "github.com/mae-pax/logger"
	"github.com/spf13/pflag"
)

type UTCTime time.Time

func (t *UTCTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.UTC)
	*t = UTCTime(now)
	return err
}

func (t UTCTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len("2006-01-02 15:04:05")+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, "2006-01-02 15:04:05")
	b = append(b, '"')
	return b, nil
}

type InfoServerStruct struct {
	ReqType  string  `gorm:"column:ReqType"`
	ReqTime  UTCTime `gorm:"column:ReqTime;type:DateTime;"`
	BuildVer string  `gorm:"column:BuildVer"`
}

var (
	confPath string
)

func init() {
	pflag.StringVar(&confPath, "conf", "configs/config.yaml", "default configs path")
}

func main() {
	// c := baselogger.New()
	// c.SetDivision("time")            // 设置归档方式，"time"时间归档 "size"文件大小归档，文件大小等可以在配置文件配置
	// c.SetTimeUnit(baselogger.Minute) // 时间归档 可以设置切割单位
	// c.SetEncoding("json")            // 输出格式 "json" 或者 "console"
	// //c.Stacktrace = true

	// c.SetInfoFile("./logs/server.log")      // 设置info级别日志
	// c.SetErrorFile("./logs/server_err.log") // 设置warn级别日志

	// c.SentryConfig = logger.SentryLoggerConfig{
	// 	DSN:              "http://b13b4964b50a41fb91a82bbb2932ca33@localhost:9000/2",
	// 	Debug:            true,
	// 	AttachStacktrace: true,
	// 	Environment:      "dev",
	// 	Tags: map[string]string{
	// 		"source": "demo",
	// 	},
	// }

	c := baselogger.NewFromYaml(confPath)
	c.CloseConsoleDisplay()
	c.SetCaller(false, 2)
	logger := c.InitLogger("time", "level", false, false)

	c.SampleRate = 0.5
	sampleLogger := c.InitSampleLogger("time", "level", false, true)

	logger.Info("info level test")
	logger.Error("dsdadadad level test", baselogger.WithError(errors.New("sabhksasas")))
	logger.Error("121212121212 error")
	logger.Error("error message", baselogger.With("foo", "bar"))
	logger.Warn("warn level test")
	logger.Debug("debug level test")
	// mock
	rawLog := `{"ReqType":"req","ReqTime":"2024-12-31 07:50:00","BuildVer":"0.0.1"}`
	var logData InfoServerStruct
	_ = json.Unmarshal([]byte(rawLog), &logData)
	// write interface{} with json format
	logger.Sugar().Infow("", baselogger.KeysAndValues(logData)...)

	sampleLogger.Info("test sampleLogger 1")
	sampleLogger.Info("test sampleLogger 2")
	sampleLogger.Info("test sampleLogger 3")
	sampleLogger.Info("test sampleLogger 4")
	sampleLogger.Info("test sampleLogger 5")
	sampleLogger.Info("test sampleLogger 6")
	sampleLogger.Info("test sampleLogger 7")
	sampleLogger.Info("test sampleLogger 8")
	sampleLogger.Info("test sampleLogger 9")
	sampleLogger.Info("test sampleLogger 10")

	// // time.Sleep(1 * time.Minute) // 避免程序结束太快，没有上传sentry

	// logger.Infof("info level test: %s", "111")
	// logger.Errorf("error level test: %s", "111")
	// logger.Warnf("warn level test: %s", "111")
	// logger.Debugf("debug level test: %s", "111")

	// logger.Info("this is a log", baselogger.With("Trace", "12345677"))
	// logger.Info("this is a log", baselogger.WithError(errors.New("this is a new error")))

	// for i := 0; i < 100; i++ {
	// 	logger.Info(myLogMsg(strconv.Itoa(i)))
	// }
}

func myLogMsg(i string) string {
	var buf bytes.Buffer
	buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	buf.WriteString("\t")
	buf.WriteString("p_id_" + i)
	buf.WriteString("\t")
	buf.WriteString("a_id_" + i)
	buf.WriteString("\t")
	buf.WriteString("u_id_" + i)
	return buf.String()
}
