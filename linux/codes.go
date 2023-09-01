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

	// SIGHUP Hangup detected on controlling terminal or death of controlling process.
	SIGHUP = 1
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
	// SIGBUS Bus error (bad memory access).
	SIGBUS = 7
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
	// SIGSTKFLT Stack fault on coprocessor (unused).
	SIGSTKFLT = 16
	// SIGCHLD Child stopped or terminated.
	SIGCHLD = 17
	// SIGCONT Continue if stopped.
	SIGCONT = 18
	// SIGSTOP Stop process.
	SIGSTOP = 19
	// SIGTSTP Stop typed at terminal.
	SIGTSTP = 20
	// SIGTTIN Terminal input for background process.
	SIGTTIN = 21
	// SIGTTOU Terminal output for background process.
	SIGTTOU = 22
	// SIGURG Urgent condition on socket (4.2BSD).
	SIGURG = 23
	// SIGXCPU CPU time limit exceeded (4.2BSD).
	SIGXCPU = 24
	// SIGXFSZ File size limit exceeded (4.2BSD).
	SIGXFSZ = 25
	// SIGVTALRM Virtual alarm clock (4.2BSD).
	SIGVTALRM = 26
	// SIGPROF Profiling timer expired.
	SIGPROF = 27
	// SIGWINCH Window resize signal.
	SIGWINCH = 28
	// SIGIO I/O now possible (4.2BSD).
	SIGIO = 29
	// SIGPWR Power failure (System V).
	SIGPWR = 30
	// SIGSYS Bad system call (SVr4).
	SIGSYS = 31
)
