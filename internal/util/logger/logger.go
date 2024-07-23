package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sfaizh/ticket-management-system/internal/globals"
	"github.com/sfaizh/ticket-management-system/internal/structs"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

const timeFormat string = "16/07/2024 00:00:00"

func logln(level structs.LogLevel, v ...interface{}) {
	logger.Println(level.String(), joinLine(v...), getSuffix(4))
}

func joinLine(v ...interface{}) string {
	str := fmt.Sprintln(v...)
	return str[:len(str)-1]
}

func getSuffix(calldepth int) string {
	fnName, file, line := getCallerFn(calldepth)

	if !globals.LogConfig.FullPaths {
		fnName = filepath.Base(file)
	}

	if globals.LogConfig.Verbose {
		return fmt.Sprintf("[%s in %s:%d]", fnName, file, line)
	}

	return fmt.Sprintf("[%s:%d]", fnName, line)
}

func getCallerFn(skipFrames int) (function string, file string, line int) {
	frame := getFrame(skipFrames)
	return frame.Function, frame.File, frame.Line
}

func getFrame(skipFrames int) runtime.Frame {

	// We need the frame at index skipFrames+2, since we never want
	// runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for
	// one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
