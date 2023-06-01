package helper

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type PrependEncoder struct {
	zapcore.Encoder
	Pool buffer.Pool
}

func ZapGetConfig(production bool) zap.Config {
	var config zap.Config

	if production {
		config = zap.NewProductionConfig()
		config.Encoding = "console"
		config.EncoderConfig.TimeKey = ""
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stdout"}
	timeEncoder := func(t time.Time, e zapcore.PrimitiveArrayEncoder) {
		e.AppendString(time.Now().Local().Format("2006-01-02 15:04:05"))
	}
	config.EncoderConfig.EncodeTime = timeEncoder

	return config
}

func (e *PrependEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// new log buffer
	buf := e.Pool.Get()

	// prepend the JournalD prefix based on the entry level
	buf.AppendString(e.toJournaldPrefix(entry.Level))
	buf.AppendString(" ")

	// calling the embedded encoder's EncodeEntry to keep the original encoding format
	consolebuf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	// just write the output into your own buffer
	_, err = buf.Write(consolebuf.Bytes())
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// some mapper function
func (e *PrependEncoder) toJournaldPrefix(lvl zapcore.Level) string {
	switch lvl {
	case zapcore.DebugLevel:
		return "<7>"
	case zapcore.InfoLevel:
		return "<6>"
	case zapcore.WarnLevel:
		return "<4>"
	}
	return ""
}
