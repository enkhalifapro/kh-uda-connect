package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "persons-service",
		Short: "persons service manages persons data",
		Long:  `persons service manages persons data`,
		Run: func(cmd *cobra.Command, args []string) {
			topic := "test1"
			partition := 0

			r := kafka.NewReader(kafka.ReaderConfig{
				Brokers:   []string{"kafka:9092"},
				Topic:     topic,
				Partition: partition,
			})

			for {
				m, err := r.ReadMessage(context.Background())
				if err != nil {
					log.Fatal("failed to write messages:", err)
				}
				// Print the message
				fmt.Printf("Message received: %s\n", string(m.Value))

				// Simulate processing time
				time.Sleep(1 * time.Second)
			}
		},
	}
)

// Execute ...
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
