package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/kava-labs/kava-devnet/blockchain/x/auction"
	"github.com/spf13/cobra"
)


// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group nameservice queries under a subcommand
	auctionQueryCmd := &cobra.Command{
		Use:   "auction",
		Short: "Querying commands for the auction module",
	}

	auctionQueryCmd.AddCommand(client.GetCommands(
		getCmdGetAuctions(queryRoute, cdc),
	)...)

	return auctionQueryCmd
}

// getCmdGetAuctions queries the auctions in the store
func getCmdGetAuctions(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getauctions",
		Short: "get a list of active auctions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getauctions", queryRoute), nil)
			if err != nil {
				fmt.Printf("error when getting auctions - %s", err)
				return nil
			}
			var out auction.QueryResAuctions
			cdc.MustUnmarshalJSON(res, &out)
			if len(out) == 0 {
				out = append(out, "There are currently no auctions")
			}
			return cliCtx.PrintOutput(out)
		},
	}
}
