package projects

import (
	"context"

	"go.saastack.io/chaku/errors"
	"go.saastack.io/idutil"
	location "go.saastack.io/location/pb"
	"go.saastack.io/project/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Generic Error to be returned to client to hide possible sensitive information
	errInternal = status.Error(codes.Internal, "oops! Something went wrong")
    locationPrefix = (&location.Location{}).GetPrefix()
)

type projectsServer struct {
	projectStore pb.ProjectStore
	projectBLoC  pb.ProjectsServiceProjectServerBLoC
	*pb.ProjectsServiceProjectServerCrud
}

func NewProjectsServer(

	projectSt pb.ProjectStore,

) pb.ProjectsServer {
	r := &projectsServer{

		projectStore: projectSt,
	}

	projectSC := pb.NewProjectsServiceProjectServerCrud(projectSt, r)
	r.ProjectsServiceProjectServerCrud = projectSC

	return r
}

// These functions represent the BLoC(Business Logical Component) of the CRUDGen Server.
// These functions will contain business logic and would be called inside the generated CRUD functions

func (s *projectsServer) CreateProjectBLoC(ctx context.Context, in *pb.CreateProjectRequest) error {

    //TODO: Check and test location parent []
    //TODO: Check and test title already exist []
    if idutil.GetParent(in.GetParent()) == locationPrefix {
        return errors.ErrInvalidField
    }


    _, err := s.projectStore.GetProject(ctx, []string{}, pb.ProjectTitleEq{Title: in.GetProject().GetTitle()})
    if err!=nil {
        if err == errors.ErrNotFound {
            return nil
        }
        return err
    }

	return errors.ErrObjIdExist
}

func (s *projectsServer) GetProjectBLoC(ctx context.Context, in *pb.GetProjectRequest) error {
	return nil
}

func (s *projectsServer) UpdateProjectBLoC(ctx context.Context, in *pb.UpdateProjectRequest) error {

    for _, m := range in.GetUpdateMask().GetPaths() {
		if !pb.ValidProjectFieldMask(m) {
			return errors.ErrInvalidField
		}

		if m == "id" {
			return errors.ErrInvalidField
		}
	}

    //TODO: Check and test title already exist
    _, err := s.projectStore.GetProject(ctx, []string{}, pb.ProjectTitleEq{Title: in.GetProject().GetTitle()})
    if err!=nil {
        if err == errors.ErrNotFound {
			return nil
		}
		return err
    }

	return errors.ErrObjIdExist
}

func (s *projectsServer) DeleteProjectBLoC(ctx context.Context, in *pb.DeleteProjectRequest) error {
	return nil
}

func (s *projectsServer) BatchGetProjectBLoC(ctx context.Context, in *pb.BatchGetProjectRequest) error {
	return nil
}

func (s *projectsServer) ListProjectBLoC(ctx context.Context, in *pb.ListProjectRequest) (pb.ProjectCondition, error) {

	// Validate view masks
	for _, m := range in.GetViewMask().GetPaths() {
		if !pb.ValidProjectFieldMask(m) {
			return nil, errors.ErrInvalidField
		}
	}

	if in.GetParent() != "" {
		return pb.ProjectParentEq{Parent: in.GetParent()}, nil
	}

	return pb.ProjectParentEq{Parent: idutil.GetId(in.GetParent())}, nil
}

// These functions are not implemented by CRUDGen, needed to be implemented
type locationServiceServer struct {
	locationClient location.LocationsClient
}


func newLocationServiceServer(in struct {
    fx.In
    LocationClient location.LocationsClient `name:"location_client"`
}) *locationServiceServer {
    return &locationServiceServer {
            locationClient: in.LocationClient,
    }
}

func (s *locationServiceServer) ValidateParent(ctx context.Context, parent string) {
}
