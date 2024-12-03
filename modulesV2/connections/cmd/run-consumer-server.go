package cmd

import (
	"enkhalifapro/connections/build"
	g "enkhalifapro/connections/grpc"
	"enkhalifapro/connections/internal"
	"fmt"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbHost                   string
	dbPort                   string
	dbUser                   string
	dbName                   string
	dbPassword               string
	kafkaAddress             string
	kafkaPartition           int
	locationsAddedTopicName  string
	locationsServiceGRPCAdrs string
	runConsumerServerCmd     = &cobra.Command{
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
			kafKaConn := kafka.NewReader(kafka.ReaderConfig{
				Brokers:   []string{kafkaAddress},
				Topic:     locationsAddedTopicName,
				Partition: kafkaPartition,
			})

			// Set up a connection to the locations grpc server.
			grpcConn, err := grpc.NewClient(locationsServiceGRPCAdrs, grpc.WithInsecure())
			if err != nil {
				logger.Fatalf("did not connect: %v", err)
			}
			defer grpcConn.Close()

			locationsSrv := g.NewLocationsClient(grpcConn)
			service := internal.NewService(db, kafKaConn, locationsSrv)
			fmt.Println(service)

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
	locationsAddedTopicName = viper.GetString("LOCATIONS_ADDED_TOPIC_NAME")
	locationsServiceGRPCAdrs = viper.GetString("LOCATIONS_SERVICE_GRPC_ADRS")
}
