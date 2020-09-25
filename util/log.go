package util
import(
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"github.com/natefinch/lumberjack"
)

func NewZapCore(file_path string, level zapcore.Level, max_size int, max_backups int, max_age int, compress bool) zapcore.Core {
	var writer zapcore.WriteSyncer
	if len(file_path) > 0 {
		/// 日志文件路径配置
		hook := lumberjack.Logger{
			Filename:   file_path,   /// 日志文件路径
			MaxSize:    max_size,    /// 每个日志文件保存的最大尺寸 单位: M
			MaxBackups: max_backups, /// 日志文件最多保存多少个备份
			MaxAge:     max_age,     /// 文件最多保存多少天
			Compress:   compress,    // 是否压缩
		}
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	/// 设置日志级别
	atomic_level := zap.NewAtomicLevel()
	atomic_level.SetLevel(level)
	// 公用编码器
	encoder_config := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder_config),
		writer,
		atomic_level,
	)
}

func NewZapDevelopmenet(output_path string) *zap.Logger {
	core := NewZapCore(output_path, zap.DebugLevel, 20, 2, 30, false)
	return zap.New(core, zap.AddCaller(), zap.Development())
}

func NewZapProduction(output_path string) *zap.Logger {
	core := NewZapCore(output_path, zap.InfoLevel, 20, 2, 30, false)
	return zap.New(core, zap.AddCaller(), zap.Development())
}
