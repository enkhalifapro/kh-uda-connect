package cmd

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	rootCmd = &cobra.Command{
		Use:   "connections-service",
		Short: "connections service manages connections data",
		Long:  `connections service manages connections data`,
		Run: func(cmd *cobra.Command, args []string) {
			topic := "test1"
			partition := 0

			conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
			if err != nil {
				log.Fatal("failed to dial leader:", err)
			}

			//conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
			for {
				id, _ := uuid.NewUUID()
				msg := fmt.Sprintf("ayman two! %v", time.Now())
				_, err = conn.WriteMessages(
					kafka.Message{
						Key:   []byte(id.String()),
						Value: []byte(msg),
					},
				)
				if err != nil {
					log.Fatal("failed to write messages:", err)
				}
				fmt.Printf("sent message id: %s\n", id.String())
				time.Sleep(time.Second * 3)
			}

			if err := conn.Close(); err != nil {
				log.Fatal("failed to close writer:", err)
			}

			//partition := 0

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
