package logger

import "go.uber.org/zap"

var L *zap.Logger
var Sl *zap.SugaredLogger

func InitLogger() {
	L, _ = zap.NewDevelopment()
	Sl = L.Sugar()

}
