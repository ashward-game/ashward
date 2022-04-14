package openboxgenesis

import (
	"context"
	"net/rpc"
	"orbit_nft/cmd/openbox"
	orbitContext "orbit_nft/contract/context"
)

func HandleLogBoxBought(ctx context.Context, evt *LogBoxBought) error {
	obRPCAddress := ctx.Value(orbitContext.KeyOpenboxRPCAddress).(string)

	worker, err := rpc.DialHTTP("tcp", obRPCAddress)
	if err != nil {
		return err
	}

	params := &openbox.HashWithClientRandom{
		Hash:         evt.ServerHash,
		ClientRandom: evt.ClientRandom.Bytes(),
		BoxGrade:     int(evt.BoxGrade),
	}

	if err := worker.Call("OpenboxService.OpenBox", params, nil); err != nil {
		return err
	}

	return nil
}

func HandleLogBoxOpened(ctx context.Context, evt *LogBoxOpened) error {
	return nil
}

func HandleLogPublicSaleOpened(ctx context.Context, evt *LogPublicSaleOpened) error {
	return nil
}

func HandleLogOwnershipTransferred(ctx context.Context, evt *LogOwnershipTransferred) error {
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

func HandleLogSubscriberRegistered(ctx context.Context, evt *LogSubscriberRegistered) error {
	return nil
}
