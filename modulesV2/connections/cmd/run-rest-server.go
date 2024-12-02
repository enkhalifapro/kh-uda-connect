package cmd

import (
	"enkhalifapro/connections/build"
	"enkhalifapro/connections/internal"
	"fmt"
	//_ "google.golang.org/grpc"
	//"net"
	"net/http"

	"enkhalifapro/connections/api/rest"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//pb "path/to/your/proto/package"
)

var (
	dbHost           string
	dbPort           string
	dbUser           string
	dbName           string
	dbPassword       string
	runRestServerCmd = &cobra.Command{
		Use:   "run-rest-server",
		Short: "run connections-service rest server",
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

			/*// init kafka client
			partition := 0
			// Set up SASL authentication
			dialer := &kafka.Dialer{
				Timeout:   10 * time.Second,
				DualStack: true,
				TLS: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			}
			if kafkaUser != "" && kafkaPassword != "" {
				dialer.SASLMechanism = plain.Mechanism{
					Username: kafkaUser,
					Password: kafkaPassword,
				}
			}

			conn, err := dialer.DialLeader(context.Background(), "tcp", kafkaBroker, kafkaTopic, partition)
			if err != nil {
				log.Fatal("failed to dial leader:", err)
			}
			defer conn.Close()
			conn.SetWriteDeadline(time.Now().Add(30 * time.Second))*/

			service := internal.NewService(db)
			handler := rest.NewHandler(service)
			router := httprouter.New()
			router.GET("/", handler.GetAll)
			router.GET("/id/:id", handler.GetByID)
			router.POST("/", handler.Create)
			router.GET("/healthz", handler.Health)

			ch := make(chan error, 0)
			go func() {
				defer close(ch)
				fmt.Println("Starting server on :8080")
				if err := http.ListenAndServe(":8080", router); err != nil {
					ch <- err
				}
			}()

			/*// Start gRPC server
			go func() {
				lis, err := net.Listen("tcp", ":50051")
				if err != nil {
					log.Fatalf("Failed to listen: %v", err)
				}
				s := grpc.NewServer()
				pb.RegisterGreeterServer(s, &server{})
				log.Println("Starting gRPC server on :50051")
				if err := s.Serve(lis); err != nil {
					log.Fatalf("Failed to serve: %v", err)
				}
			}()*/

			select {
			case err := <-ch:
				fmt.Println(err)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	getRestEnvVars()
	rootCmd.AddCommand(runRestServerCmd)
}

func getRestEnvVars() {
	dbHost = viper.GetString("DB_HOST")
	dbPort = viper.GetString("DB_PORT")
	dbUser = viper.GetString("DB_USERNAME")
	dbPassword = viper.GetString("DB_PASSWORD")
	dbName = viper.GetString("DB_NAME")
}
