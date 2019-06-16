package pool

import "github.com/cosmos/cosmos-sdk/codec"

// generic sealed codec to be used throughout module
var moduleCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	moduleCdc = cdc.Seal()
}

// RegisterCodec registers concrete types on the codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgDepositFund{}, "pool/MsgDepositFund", nil)
	cdc.RegisterConcrete(MsgWithdrawFund{}, "pool/MsgWithdrawFund", nil)
}
