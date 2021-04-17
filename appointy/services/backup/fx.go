package projects

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.saastack.io/jaal/schemabuilder"
	"go.saastack.io/project/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Module is the fx module encapsulating all the providers of the package
var Module = fx.Options(
	fx.Provide(
		pb.NewPostgresProjectStore,

		NewProjectsServer,
		fx.Annotated{
			Name: "public",
			Target: func(srv pb.ProjectsServer) pb.ProjectsServer {
				return srv
			},
		},

		pb.NewLocalProjectsClient,
		fx.Annotated{
			Name: "public",
			Target: func(in struct {
				fx.In
				S pb.ProjectsServer `name:"public"`
			}) pb.ProjectsClient {
				return pb.NewLocalProjectsClient(in.S)
			},
		},

		fx.Annotated{
			Group:  "grpc-service",
			Target: RegisterGRPCService,
		},
		fx.Annotated{
			Group:  "graphql-service",
			Target: RegisterGraphQLService,
		},
		fx.Annotated{
			Group:  "http-service",
			Target: RegisterHttpService,
		},
	),
	fx.Decorate(
		pb.NewEventsProjectsServer,
	),
)

func RegisterGRPCService(in struct {
	fx.In
	Server pb.ProjectsServer `name:"public"`
}) func(s *grpc.Server) {
	return func(s *grpc.Server) {
		pb.RegisterProjectsServer(s, in.Server)
	}
}

func RegisterGraphQLService(in struct {
	fx.In
	Client pb.ProjectsClient `name:"public"`
}) func(s *schemabuilder.Schema) {
	return func(s *schemabuilder.Schema) {
		pb.RegisterProjectsOperations(s, in.Client)
	}
}

func RegisterHttpService(in struct {
	fx.In
	Client pb.ProjectsClient `name:"public"`
}) func(*runtime.ServeMux, context.Context) error {
	return func(mux *runtime.ServeMux, ctx context.Context) error {
		return pb.RegisterProjectsHandlerClient(ctx, mux, in.Client)
	}
}

