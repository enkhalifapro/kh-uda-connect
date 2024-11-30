package cmd

import (
	"enkhalifapro/persons/build"
	"fmt"
	"net/http"

	"enkhalifapro/persons/api/rest"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbHost           string
	dbPort           string
	dbUser           string
	dbName           string
	dbPassword       string
	runRestServerCmd = &cobra.Command{
		Use:   "run-rest-server",
		Short: "run persons-service rest server",
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

			handler := rest.NewHandler()
			http.HandleFunc("/healthz", handler.Health)
			fmt.Println("Starting server on :8080")
			if err := http.ListenAndServe(":8080", nil); err != nil {
				fmt.Println("Error starting server:", err)
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
