package token

import "context"

func HandleLogApproval(ctx context.Context, evt *LogApproval) error {
	return nil
}

func HandleLogPaused(ctx context.Context, evt *LogPaused) error {
	return nil
}

func HandleLogRoleAdminChanged(ctx context.Context, evt *LogRoleAdminChanged) error {
	return nil
}

func HandleLogRoleGranted(ctx context.Context, evt *LogRoleGranted) error {
	return nil
}

func HandleLogRoleRevoked(ctx context.Context, evt *LogRoleRevoked) error {
	return nil
}

func HandleLogTaxUpdated(ctx context.Context, evt *LogTaxUpdated) error {
	return nil
}

func HandleLogTransfer(ctx context.Context, evt *LogTransfer) error {
	return nil
}

func HandleLogUnpaused(ctx context.Context, evt *LogUnpaused) error {
	return nil
}
