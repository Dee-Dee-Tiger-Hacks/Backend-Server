package gapi

import (
	"context"
	"database/sql"

	db "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/sqlc"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/pb"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetResume(ctx context.Context, req *pb.GetResumeRequest) (*pb.GetResumeResponse, error) {
	userId := uuid.MustParse(req.GetUserId())
	resume, err := server.store.GetResume(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "resume not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get resume: %s", err)
	}
	rsp := &pb.GetResumeResponse{
		Resume: convertResume(resume),
	}
	return rsp, nil
}

func (server *Server) CreateResume(ctx context.Context, req *pb.CreateResumeRequest) (*pb.CreateResumeResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	resumeId, err := uuid.NewRandom()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate id: %s", err)
	}
	arg := db.CreateResumeParams{
		ID:             resumeId,
		UserID:         userId,
		ResumePublicID: req.GetResumePublicId(),
		ResumeTitle:    req.GetResumeTitle(),
		ResumePdfUrl:   req.GetResumePdfUrl(),
	}

	resume, err := server.store.CreateResume(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create resume: %s", err)
	}

	rsp := &pb.CreateResumeResponse{
		Resume: convertResume(resume),
	}
	return rsp, nil
}

func (server *Server) UpdateResume(ctx context.Context, req *pb.UpdateResumeRequest) (*pb.UpdateResumeResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	arg := db.UpdateResumeParams{
		UserID: userId,
		ResumePublicID: pgtype.Text{
			String: req.GetResumePublicId(),
			Valid:  req.ResumePublicId != nil,
		},
		ResumeTitle: pgtype.Text{
			String: req.GetResumeTitle(),
			Valid:  req.ResumeTitle != nil,
		},
		ResumePdfUrl: pgtype.Text{
			String: req.GetResumePdfUrl(),
			Valid:  req.ResumePdfUrl != nil,
		},
	}

	resume, err := server.store.UpdateResume(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "resume not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to update resume: %s", err)
	}

	rsp := &pb.UpdateResumeResponse{
		Resume: convertResume(resume),
	}
	return rsp, nil
}
