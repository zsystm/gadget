package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const timestampLayout = "2006:01:02 15:04:05"

var TimestampCmd = &cobra.Command{
	Use:   "timestamp [value]",
	Short: "Convert between Unix timestamp and UTC date",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]
		if ts, err := strconv.ParseInt(input, 10, 64); err == nil {
			t := time.Unix(ts, 0).UTC()
			cmd.Printf("date: %s\n", t.Format(timestampLayout))
			return nil
		}

		t, err := time.ParseInLocation(timestampLayout, input, time.UTC)
		if err != nil {
			return fmt.Errorf("invalid input: expected unix timestamp or date in YYYY:MM:DD h:m:s format")
		}

		cmd.Printf("timestamp: %d\n", t.UTC().Unix())
		return nil
	},
}
