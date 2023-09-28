/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net"

	pb "github.com/KavishkaRajakaruna/go-tutorial/grpc/pkg/grpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Provides server",
	Long:  `Provides server`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		//register services
		pb.RegisterGopherServer(grpcServer, &Server{})
		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve %v", err)
		}
	},
}

const (
	port         = ":9000"
	KubeGoAPIURL = "https://kubego-api-xxxxx-ew.a.run.app"
)

type Server struct {
	pb.UnimplementedGopherServer
}

type Gopher struct {
	URL string `json: "url"`
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
