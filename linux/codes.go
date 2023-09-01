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

	// SIGINT Interrupt from keyboard.
	SIGINT = 2
	// SIGQUIT Quit from keyboard.
	SIGQUIT = 3
	// SIGILL Illegal Instruction.
	SIGILL = 4
	// SIGTRAP Trace/breakpoint trap.
	SIGTRAP = 5
	// SIGABRT Abort signal from abort().
	SIGABRT = 6
	// SIGFPE Floating-point exception.
	SIGFPE = 8
	// SIGKILL Kill signal.
	SIGKILL = 9
	// SIGSEGV Invalid memory reference.
	SIGSEGV = 11
	// SIGPIPE Broken pipe: write to pipe with no readers.
	SIGPIPE = 13
	// SIGALRM Timer signal from alarm().
	SIGALRM = 14
	// SIGTERM Termination signal.
	SIGTERM = 15
)
