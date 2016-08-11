package htmlerror

import (
	"bytes"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

type Frame struct {
	Filename string `json:"filename"`
	Function string `json:"function"`
	Module   string `json:"module"`

	Line int    `json:"line"`
	Path string `json:"path"`

	Context     string   `json:"content"`
	PreContext  []string `json:"pre_context,omitempty"`
	PostContext []string `json:"post_context,omitempty"`
}

func NewStacktrace(skip int) []*Frame {
	var frames []*Frame

	for i := 1 + skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		frame := NewStacktraceFrame(pc, file, line)
		if frame != nil {
			frames = append(frames, frame)
		}
	}

	if len(frames) < 1 {
		return nil
	}

	return frames
}

func NewStacktraceFrame(pc uintptr, file string, line int) *Frame {
	frame := &Frame{
		Filename: file,
		Line:     line,
	}
	frame.Module, frame.Function = funcName(pc)

	// runtime.goexit is effectively a placeholder that comes from
	// runtime/asm_amd64.s and is meaningless.
	if frame.Module == "runtime" && frame.Function == "goexit" {
		return nil
	}

	frame.Context, _ = fileLine(file, line)

	return frame
}

// return the name of function and package containing given pc
func funcName(pc uintptr) (pack string, name string) {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return
	}
	name = fn.Name()

	// split full package.method name into components
	if idx := strings.LastIndex(name, "."); idx != -1 {
		pack = name[:idx]
		name = name[idx+1:]
	}
	name = strings.Replace(name, "·", ".", -1)

	return
}

// return line from file
func fileLine(filename string, line int) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	lines := bytes.Split(data, []byte{'\n'})

	line--
	if line <= len(lines) {
		return string(lines[line]), nil
	}

	return "", nil
}
