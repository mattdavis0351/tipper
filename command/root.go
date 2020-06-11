package command

import (
	"fmt"
	"tipper/bill"

	"github.com/spf13/cobra"
)

var (
	foodTotal float32
	tip       float32
)

func init() {
	RootCmd.PersistentFlags().Float32Var(&foodTotal, "food", 0, "food is total bill amount that you want a tip for (default value is 0)")
	RootCmd.PersistentFlags().Float32Var(&tip, "tip", 0, "food is total bill amount that you want a tip for (default value is 0)")

}

// RootCmd needs a comment so you leave me alone
var RootCmd = &cobra.Command{
	Use:   "tipper",
	Short: "tipper is a tip calculator",
	Long:  `Tipper is a command line tool to calculate the total bill and a tip when you eat out at a restaurant and happen to carry a laptop like a loser`,
	Run: func(cmd *cobra.Command, args []string) {
		b := &bill.FoodBill{TotalAmount: foodTotal}

		_, err := b.CalculateTip(b, tip)
		if err != nil {
			fmt.Println(err)
		}
	},
}
