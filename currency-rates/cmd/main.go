package main

import (
	"context"
	"log"
	"time"

	"github.com/RinaDish/currency-rates/internal/app"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() { //отсюда начинаем
	var cfg app.Config
	err := envconfig.Process("currency-rates", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	config := zap.NewProductionConfig() //инициализация конфига логгера
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339) //устанавливает формат времени 2024-06-11T10:20:30Z

	l, _ := config.Build() //инициализация логгера
	defer l.Sync() // nolint:errcheck 
	// <-- Важно для завершения работы с логгером Если не вызвать Sync(), часть логов может остаться в буфере и не попасть в конечный лог-файл или другой выходной поток.
	logger := l.Sugar() //Как видно, "сахарный" логгер делает код более компактным и легким для восприятия.

	ctx := context.Background() //создает новый, пустой контекст

	application := app.NewApp(cfg, logger) // идем в internal/app создаем экземпляр структуры App
	if err := application.Run(ctx); err != nil { //запускаем приложение
		logger.Error(err)
	}
}
