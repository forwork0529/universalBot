package logger

import (
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var appLogger *zap.SugaredLogger

func New() {
	// 1
	writerSyncer := getLogWriter()
	// 2
	encoder := getEncoder()
	// 3 Определяем уровень логирования
	debugLevel := getDebugLevel()
	// 4 На основе того куда пишем, энкодера(конфига) , уровня логгирования создаём ядро логгера
	core := zapcore.NewCore(encoder, writerSyncer, debugLevel)
	// 5 на основании ядра и опций создаём сам логгер
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
	// 6 несколько видоизменяем интерфейс логгера
	appLogger = logger.Sugar()
	// 7 опусташаем буффер логгера
	defer appLogger.Sync()

}


// 1 определяем место записи, обычный io.Writer переделываем в zapcore.WriteSyncer
func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

// 2 здесь создаём конфиг файл, определяем его поля , на основе конфига создаём энкодер (интерфейс нужный для работы логгера)
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 3 Получаем желаемый уровень логирования из переменных окружения и преобразуем его в внутреннему типу логера
func getDebugLevel()zapcore.Level{
	got , ok := os.LookupEnv("DEBUG_LEVEL")
	if !ok {
		log.Printf("INFO: no DEBUG_LEVEL in env variables, chosed -1")
		return zapcore.DebugLevel
	}

	debugLevel, err := strconv.Atoi(got)
	if err != nil{
		log.Printf("getDebugLevel(): %v\n", err.Error())
		return zapcore.DebugLevel
	}
	if debugLevel < 0 || debugLevel > 5{
		log.Println("getDebugLevel(): invalid input DEBUG_LEVEL")
		return zapcore.DebugLevel
	}
	log.Printf("INFO: chosed DEBUG_LEVEL = %v\n", debugLevel)
	return zapcore.Level(debugLevel)
}

// Далее у нас идут функции, вызывающие в себе методы сконфигуиррованного выше глобального логгера

func Debug(args ...interface{}) {
	appLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	appLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	appLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	appLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	appLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	appLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	appLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	appLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	appLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	appLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	appLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	appLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	appLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	appLogger.Fatalf(template, args...)
}
