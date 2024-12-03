package gapi

import (
	"context"

	db "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/sqlc"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/pb"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.CreateEmailResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	arg := db.CreateEmailParams{
		ID:           id,
		UserID:       userId,
		Subject:      req.GetSubject(),
		Content:      req.GetContent(),
		Title:        req.GetTitle(),
		EmailAddress: req.GetEmailAddress(),
	}

	email, err := server.store.CreateEmail(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create email: %s", err)
	}

	rsp := &pb.CreateEmailResponse{
		Email: convertEmail(email),
	}
	return rsp, nil
}

func (server *Server) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.GetEmailResponse, error) {
	id := uuid.MustParse(req.GetId())

	email, err := server.store.GetEmailById(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get email: %s", err)
	}

	rsp := &pb.GetEmailResponse{
		Email: convertEmail(email),
	}

	return rsp, nil
}

func (server *Server) GetEmails(ctx context.Context, req *pb.GetEmailsRequest) (*pb.GetEmailsResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	numOfEmails, err := server.store.CountEmailsByUserId(ctx, userId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to count emails: %s", err)
	}

	arg := db.GetEmailsByUserIdParams{
		UserID: userId,
		Offset: max(req.GetPage()-1, 0) * req.GetLimit(),
		Limit:  req.GetLimit(),
	}

	emails, err := server.store.GetEmailsByUserId(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get emails: %s", err)
	}

	rsp := &pb.GetEmailsResponse{
		Emails: convertEmails(emails),
		Total:  int32(numOfEmails),
	}

	return rsp, nil
}

func (server *Server) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateEmailResponse, error) {
	userId := uuid.MustParse(req.GetUserId())
	id := uuid.MustParse(req.GetId())

	email, err := server.store.GetEmailById(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get email: %s", err)
	}

	if email.UserID != userId {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other user's email")
	}

	arg := db.UpdateEmailParams{
		ID:      id,
		Subject: pgtype.Text{String: req.GetSubject(), Valid: req.Subject != nil},
		Content: pgtype.Text{String: req.GetContent(), Valid: req.Content != nil},
		Title:   pgtype.Text{String: req.GetTitle(), Valid: req.Title != nil},
	}

	updatedEmail, err := server.store.UpdateEmail(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update email: %s", err)
	}

	rsp := &pb.UpdateEmailResponse{
		Email: convertEmail(updatedEmail),
	}

	return rsp, nil
}
