package logger

import (
	"context"
)

func formatArgs(ctx context.Context,args []interface{}) interface{} {
	args = append(args, "traceId:",ctx.Value(TraceKey))
	return args
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sugar.Debug(args)
}
func DebugWithCtx(ctx context.Context,args ...interface{})  {
	sugar.Debug(formatArgs(ctx,args))
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sugar.Info(args)
}

func InfoWithCtx(ctx context.Context,args ...interface{}) {
	sugar.Info(formatArgs(ctx,args))
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sugar.Warn(args)
}

func WarnWithCtx(ctx context.Context,args ...interface{}) {
	sugar.Warn(formatArgs(ctx,args))
}


// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sugar.Error(args)
}

func ErrorWithCtx(ctx context.Context,args ...interface{}) {
	//args = append(args, "traceId",ctx.Value(TraceKey))
	//sugar.Error(args)
	sugar.Error(formatArgs(ctx,args))
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	sugar.DPanic(args)
}

func DPanicWithCtx(ctx context.Context,args ...interface{}) {
	sugar.DPanic(formatArgs(ctx,args))
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	sugar.Panic(args)
}

func PanicWithCtx(ctx context.Context,args ...interface{}) {
	sugar.Panic(formatArgs(ctx,args))
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	sugar.Fatal(args)
}

func FatalWithCtx(ctx context.Context,args ...interface{}) {
	sugar.Fatal(formatArgs(ctx,args))
}