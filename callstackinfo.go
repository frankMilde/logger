package logger

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	MAX_FILENAME_LENGTH int
)

func init() {
	sourceFiles := getSourceFiles()
	MAX_FILENAME_LENGTH = getMaxFileNameLength(sourceFiles)
}

func getCallerInfo() string {
	pc, file, line, _ := getCaller()

	baseName := filepath.Base(file)
	pad := MAX_FILENAME_LENGTH - len(baseName)

	fileName := strings.Repeat(" ", pad) + baseName
	method := filepath.Ext(runtime.FuncForPC(pc).Name())
	lineNr := strconv.Itoa(line)

	return fileName + method + "[" + lineNr + "]    "
}

func getCaller() (pc uintptr, file string, line int, ok bool) {

	stackCounter := 0
	_, callingFile, _, _ := runtime.Caller(stackCounter)

	for strings.Contains(filepath.Dir(callingFile), "logger") {
		stackCounter++
		_, callingFile, _, _ = runtime.Caller(stackCounter)
	}

	return runtime.Caller(stackCounter)
}

func getSourceFiles() []os.FileInfo {

	allFiles, _ := ioutil.ReadDir("./")

	var sourceFiles []os.FileInfo

	for _, f := range allFiles {
		if isGoSourceFile(f.Name()) {
			sourceFiles = append(sourceFiles, f)
		}
	}

	return sourceFiles
}

func isGoSourceFile(name string) bool {
	return filepath.Ext(name) == ".go" && !strings.Contains(name, "_test")
}

func getMaxFileNameLength(sourceFiles []os.FileInfo) int {

	maxLength := -1

	for _, f := range sourceFiles {
		length := len(f.Name())
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
