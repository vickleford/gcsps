package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vickleford/gcsps/internal/gcs"
	"google.golang.org/api/option"
	pubsub "google.golang.org/api/pubsub/v1"
)

var subscribeCmd = &cobra.Command{
	Use:  "subscribe [PROJECT] [TOPIC] [SUBSCRIPTION]",
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		project := args[0]
		topic := args[1]
		subscription := args[2]

		svc, err := pubsub.NewService(context.TODO(),
			option.WithEndpoint(endpoint),
			option.WithoutAuthentication(),
		)
		if err != nil {
			return err
		}

		client := gcs.New(project, svc)

		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()

		events, err := client.Subscribe(ctx, topic, subscription)
		if err != nil {
			return err
		}

		for event := range events {
			if event.Error != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", event.Message.ID, event.Message.Data)
		}

		return nil
	},
}