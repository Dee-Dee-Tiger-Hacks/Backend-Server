package gapi

import (
	"context"

	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRecruiter(ctx context.Context, req *pb.GetRecruiterRequest) (*pb.GetRecruiterResponse, error) {
	recruiterId := uuid.MustParse(req.GetRecruiterId())

	recruiter, err := server.store.GetRecruiterById(ctx, recruiterId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get recruiter: %s", err)
	}

	rsp := &pb.GetRecruiterResponse{
		Recruiter: convertRecruiter(recruiter),
	}

	return rsp, nil
}

func (server *Server) GetRecruiters(ctx context.Context, req *pb.GetRecruitersRequest) (*pb.GetRecruitersResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	recruiters, err := server.store.GetRecruitersByUserId(ctx, userId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get movie: %s", err)
	}

	rsp := &pb.GetRecruitersResponse{
		Recruiters: convertRecruiters(recruiters),
	}

	return rsp, nil
}
