package safeguard

const signalPrefix = "signal "

// Signal table.
var signals = map[string]int{ //nolint:gochecknoglobals // global signal table.
	"hangup":                   1,
	"interrupt":                2,
	"quit":                     3,
	"illegal instruction":      4,
	"trace/BPT trap":           5,
	"abort trap":               6,
	"EMT trap":                 7,
	"floating point exception": 8,
	"killed":                   9,
	"bus error":                10,
	"segmentation fault":       11,
	"bad system call":          12,
	"broken pipe":              13,
	"alarm clock":              14,
	"terminated":               15,
	"urgent I/O condition":     16,
	"suspended (signal)":       17,
	"suspended":                18,
	"continued":                19,
	"child exited":             20,
	"stopped (tty input)":      21,
	"stopped (tty output)":     22,
	"I/O possible":             23,
	"cputime limit exceeded":   24,
	"filesize limit exceeded":  25,
	"virtual timer expired":    26,
	"profiling timer expired":  27,
	"window size changes":      28,
	"information request":      29,
	"user defined signal 1":    30,
	"user defined signal 2":    31,
}
