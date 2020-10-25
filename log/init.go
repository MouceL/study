package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

type Conf struct {
	File string
	Level string
	MaxSize , MaxAge , MaxBackups int
}

var conf Conf

// zap 没有 rotate 的功能，它只负责写日志，所以要配合 lumberjack 进行日志 rotate

// zap 是对 zap.core 更高级的封装

func init(){
	var core zapcore.Core
	if conf.File != ""{
		luber := &lumberjack.Logger{
			Filename:   conf.File,
			MaxSize:    conf.MaxSize,
			MaxAge:     conf.MaxAge,
			MaxBackups: conf.MaxBackups,
		}
		encoderCfg := zap.NewProductionEncoderConfig()
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg),zapcore.AddSync(luber),zap.DebugLevel)
	}else{
		encoderCfg := zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeTime = TimeEncoder
		//core = zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg),os.Stdout,zap.DebugLevel)
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg),os.Stdout,zap.DebugLevel)
	}
	Logger = zap.New(core).WithOptions(zap.AddCaller(),zap.AddCallerSkip(1)).Sugar()
}

func TimeEncoder ( t time.Time, enc zapcore.PrimitiveArrayEncoder){
	enc.AppendString(t.Format("2006-01-02T15:04:05.000"))
}