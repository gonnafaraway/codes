package linux

const (
	// ExitCodes
	NoError                  = 0
	WithError                = 1
	CmdError                 = 2
	PermissionDenied         = 126
	CmdNotFound              = 127
	CmdTerminated            = 128
	CmdTerminatedWithSIGINT  = 130
	CmdTerminatedWithSIGTERM = 143
	CmdTerminatedOther       = 255

	// Signals
	SIGINT  = 2
	SIGQUIT = 3
	SIGILL  = 4
	SIGTRAP = 5
	SIGABRT = 6
	SIGFPE  = 8
	SIGKILL = 9
	SIGSEGV = 11
	SIGPIPE = 13
	SIGALRM = 14
	SIGTERM = 15
)
