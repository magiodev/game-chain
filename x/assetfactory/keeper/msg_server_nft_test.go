package keeper

import (
	"context"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	"reflect"
	"testing"
)

func Test_msgServer_MintNft(t *testing.T) {
	type fields struct {
		Keeper Keeper
	}
	type args struct {
		goCtx context.Context
		msg   *types.MsgMintNft
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.MsgMintNftResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := msgServer{
				Keeper: tt.fields.Keeper,
			}
			got, err := k.MintNft(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("MintNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MintNft() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_UpdateNft(t *testing.T) {
	type fields struct {
		Keeper Keeper
	}
	type args struct {
		goCtx context.Context
		msg   *types.MsgUpdateNft
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.MsgUpdateNftResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := msgServer{
				Keeper: tt.fields.Keeper,
			}
			got, err := k.UpdateNft(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateNft() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_BurnNft(t *testing.T) {
	type fields struct {
		Keeper Keeper
	}
	type args struct {
		goCtx context.Context
		msg   *types.MsgBurnNft
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.MsgBurnNftResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := msgServer{
				Keeper: tt.fields.Keeper,
			}
			got, err := k.BurnNft(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("BurnNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BurnNft() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_TransferNft(t *testing.T) {
	type fields struct {
		Keeper Keeper
	}
	type args struct {
		goCtx context.Context
		msg   *types.MsgTransferNft
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.MsgTransferNftResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := msgServer{
				Keeper: tt.fields.Keeper,
			}
			got, err := k.TransferNft(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferNft() got = %v, want %v", got, tt.want)
			}
		})
	}
}
