package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vickleford/gcsps/internal/gcs"
	"google.golang.org/api/option"
	pubsub "google.golang.org/api/pubsub/v1"
)

var publishCmd = &cobra.Command{
	Use:  "publish [PROJECT] [TOPIC] [DATA]",
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		project := args[0]
		topic := args[1]
		data := args[2]

		svc, err := pubsub.NewService(context.TODO(),
			option.WithEndpoint(endpoint),
			option.WithoutAuthentication(),
		)
		if err != nil {
			return err
		}

		client := gcs.New(project, svc)

		fmt.Fprintf(cmd.OutOrStdout(), "publishing to project %s on topic %s: %s\n", project, topic, data)

		id, err := client.Publish(context.TODO(), topic, data)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "published message successfully with id %s\n", id)

		return nil
	},
}
