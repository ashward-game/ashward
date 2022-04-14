package stakingrewards

import "context"

func HandleLogExited(ctx context.Context, evt *LogExited) error {
	return nil
}

func HandleLogOwnershipTransferred(ctx context.Context, evt *LogOwnershipTransferred) error {
	return nil
}

func HandleLogPaused(ctx context.Context, evt *LogPaused) error {
	return nil
}

func HandleLogRewardsPaid(ctx context.Context, evt *LogRewardsPaid) error {
	return nil
}

func HandleLogStaked(ctx context.Context, evt *LogStaked) error {
	return nil
}

func HandleLogStopped(ctx context.Context, evt *LogStopped) error {
	return nil
}

func HandleLogTokensApproved(ctx context.Context, evt *LogTokensApproved) error {
	return nil
}

func HandleLogTokensReceived(ctx context.Context, evt *LogTokensReceived) error {
	return nil
}

func HandleLogUnpaused(ctx context.Context, evt *LogUnpaused) error {
	return nil
}

func HandleLogWithdrawn(ctx context.Context, evt *LogWithdrawn) error {
	return nil
}
