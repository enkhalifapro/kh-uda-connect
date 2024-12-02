package cmd

import (
	"fmt"

	"enkhalifapro/connections/build"
	"enkhalifapro/connections/internal"

	"github.com/segmentio/kafka-go"

	//_ "google.golang.org/grpc"
	//"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//pb "path/to/your/proto/package"
)

var (
	dbHost                 string
	dbPort                 string
	dbUser                 string
	dbName                 string
	dbPassword             string
	kafkaAddress           string
	kafkaPartition         int
	locationAddedTopicName string
	runConsumerServerCmd   = &cobra.Command{
		Use:   "run-consumer-server",
		Short: "run connections-service consumer server",
		Long:  `run command will start gathering confluence data`,
		Run: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			// todo: add version to logger
			logger := logrus.WithFields(
				logrus.Fields{
					"service": build.AppName,
					"version": build.Version,
				},
			)
			// init db connection
			connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
			db, err := sqlx.Connect("postgres", connStr)
			if err != nil {
				logger.Fatalln("DB Connection Error:", err)
			}
			db.SetMaxOpenConns(1)
			db.SetMaxIdleConns(1)
			db.SetConnMaxLifetime(0) // 0, connections are reused forever.

			// init kafka consumer connection
			conn := kafka.NewReader(kafka.ReaderConfig{
				Brokers:   []string{kafkaAddress},
				Topic:     locationAddedTopicName,
				Partition: kafkaPartition,
			})

			service := internal.NewService(db, conn)

		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	getServerEnvVars()
	rootCmd.AddCommand(runConsumerServerCmd)
}

func getServerEnvVars() {
	dbHost = viper.GetString("DB_HOST")
	dbPort = viper.GetString("DB_PORT")
	dbUser = viper.GetString("DB_USERNAME")
	dbPassword = viper.GetString("DB_PASSWORD")
	dbName = viper.GetString("DB_NAME")

	kafkaAddress = viper.GetString("KAFKA_ADDRESS")
	kafkaPartition = viper.GetInt("KAFKA_PARTITION")
	locationAddedTopicName = "locationAddedTopic"
}
