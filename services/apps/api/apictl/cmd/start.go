package cmd

import (
	"fmt"
	"strings"

	publicGRPC "github.com/jcfug8/daylier-backend/services/adapters/grpc"
	publicHTTP "github.com/jcfug8/daylier-backend/services/adapters/http"
	"github.com/jcfug8/daylier-backend/services/apps/api/adapters/grpc"
	"github.com/jcfug8/daylier-backend/services/apps/api/adapters/http"
	"github.com/jcfug8/daylier-backend/services/apps/api/app"

	"github.com/jcfug8/daylier-backend/services/apps/api/ports/api"
	"github.com/jcfug8/daylier-backend/services/apps/api/ports/users"
	"github.com/jcfug8/daylier-backend/services/ports/public"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	PUBLIC_TYPE_KEY = "public-type"
	USERS_TYPE_KEY  = "users-type"
)

// adapter types
const (
	GRPCType = "grpc://"
	HTTPType = "http://"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long: `starts the api service.
	Example: ./bin/publicctl start`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// SET_UP_P3_CLIENT
		usersClient, err := getUserClient(viper.GetString(USERS_TYPE_KEY))
		if err != nil {
			return fmt.Errorf("unable to get users client: %s", err)
		}

		app := app.NewService(usersClient)

		apiServer, err := getApiServer(viper.GetString(PUBLIC_TYPE_KEY), app)
		if err != nil {
			return fmt.Errorf("unable to create api server: %s", err)
		}

		return apiServer.Serve()
	},
}

func init() {
	RootCmd.AddCommand(StartCmd)

	StartCmd.PersistentFlags().String(PUBLIC_TYPE_KEY, "http://0.0.0.0:80", "type, address and port for service")
	viper.BindPFlag(PUBLIC_TYPE_KEY, StartCmd.PersistentFlags().Lookup(PUBLIC_TYPE_KEY))

	StartCmd.PersistentFlags().String(USERS_TYPE_KEY, "http://0.0.0.0:80", "type, address and port for users service")
	viper.BindPFlag(USERS_TYPE_KEY, StartCmd.PersistentFlags().Lookup(USERS_TYPE_KEY))
}

func getUserClient(t string) (users.Client, error) {
	var usersClient users.Client
	var err error

	switch {
	case strings.HasPrefix(t, GRPCType):
		usersClient, err = grpc.NewUsersClient(strings.TrimPrefix(t, GRPCType), &publicGRPC.CSDialer{})
		if err != nil {
			return nil, fmt.Errorf("unable to create grpc users client: %s", err)
		}
	default:
		return nil, fmt.Errorf("unknown users client type - %s", t)
	}
	return usersClient, nil
}

func getApiServer(t string, service api.Service) (public.Server, error) {
	var server public.Server

	switch {
	case strings.HasPrefix(t, HTTPType):
		log.Info("setting up http public interface for dispatcher - ", t)
		server = publicHTTP.NewCSServer(strings.TrimPrefix(t, HTTPType), http.NewApiService(service))
	default:
		return nil, fmt.Errorf("unknown dispatcher server type - %s", t)
	}

	return server, nil
}
