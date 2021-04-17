package projects

import (
	"context"

	"go.saastack.io/chaku/errors"
	"go.saastack.io/project/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Generic Error to be returned to client to hide possible sensitive information
	errInternal = status.Error(codes.Internal, "oops! Something went wrong")
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
	return nil
}

func (s *projectsServer) GetProjectBLoC(ctx context.Context, in *pb.GetProjectRequest) error {
	return nil
}

func (s *projectsServer) UpdateProjectBLoC(ctx context.Context, in *pb.UpdateProjectRequest) error {
	//TODO: Check if the title is already given
	return nil
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

	////if idutil.GetPrefix(in.GetParent()) == locationPrefix {
	////	return pb.DepartmentFullParentEq{Parent: idutil.GetParent(in.GetParent())}, nil
	////}

	//return pb.ProjectParentEq{Parent: idutil.GetId(in.GetParent())}, nil

	if in.GetParent() != "" {
		return pb.ProjectParentEq{Parent: in.GetParent()}, nil
	}

	return pb.TrueCondition{}, nil
}

// These functions are not implemented by CRUDGen, needed to be implemented
