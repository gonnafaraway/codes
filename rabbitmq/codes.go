package rabbitmq

const (
	// QUEUE_PENDING The queue used by the API to insert new tasks into queue.
	QUEUE_PENDING = "pending"
	// QUEUE_STARTED The queue used by workers to notify the coordinator that a task has begun processing.
	QUEUE_STARTED = "started"
	// QUEUE_COMPLETED The queue used by workers to send tasks to when a task completes successfully.
	QUEUE_COMPLETED = "completed"
	// QUEUE_ERROR The queue used by workers to send tasks to when an error occurs in processing.
	QUEUE_ERROR = "error"
	// QUEUE_DEFAULT The default queue for tasks.
	QUEUE_DEFAULT = "default"
	// QUEUE_HEARBEAT The queue used by workers to periodically notify the coordinator about their aliveness.
	QUEUE_HEARBEAT = "hearbeat"
	// QUEUE_JOBS The queue used by for job creation and job-related state changes (e.g. cancellation).
	QUEUE_JOBS = "jobs"
)
