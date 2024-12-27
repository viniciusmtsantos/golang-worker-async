package process

var (
	ProcessCreditUserPoints = taskNameQueueName{
		"task:process_credit_user_points",
		"ProcessCreditUserPoints"}
)

type taskNameQueueName struct {
	TaskName  string
	QueueName string
}
