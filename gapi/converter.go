package gapi

import (
	"time"

	db "github.com/Dee-Dee-Tiger-Hacks/Backend-Server/db/sqlc"
	"github.com/Dee-Dee-Tiger-Hacks/Backend-Server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:                user.ID.String(),
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		LinkedinUrl:       user.LinkedinUrl,
		Phone:             user.Phone,
		AvatarUrl:         user.AvatarUrl,
		JobTitle:          user.JobTitle,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreateAt),
	}
}

func convertResume(resume db.Resume) *pb.Resume {
	return &pb.Resume{
		Id:             resume.ID.String(),
		UserId:         resume.UserID.String(),
		ResumePublicId: resume.ResumePublicID,
		ResumeTitle:    resume.ResumeTitle,
		ResumePdfUrl:   resume.ResumePdfUrl,
	}
}

func convertRecruiter(recruiter db.Recruiter) *pb.Recruiter {
	return &pb.Recruiter{
		Id:              recruiter.ID.String(),
		UserId:          recruiter.UserID.String(),
		LinkedinUrl:     recruiter.LinkedinUrl,
		Company:         recruiter.Company,
		Email:           recruiter.Email,
		Overview:        recruiter.Overview,
		SuggestedEmail:  recruiter.SuggestedEmail,
		PotentialTopics: recruiter.PotentialTopics,
	}
}

func convertRecruiters(recruiters []db.Recruiter) []*pb.Recruiter {
	res := []*pb.Recruiter{}
	for i := 0; i < len(recruiters); i++ {
		res = append(res, &pb.Recruiter{
			Id:              recruiters[i].ID.String(),
			UserId:          recruiters[i].UserID.String(),
			LinkedinUrl:     recruiters[i].LinkedinUrl,
			Company:         recruiters[i].Company,
			Email:           recruiters[i].Email,
			Overview:        recruiters[i].Overview,
			SuggestedEmail:  recruiters[i].SuggestedEmail,
			PotentialTopics: recruiters[i].PotentialTopics,
		})
	}
	return res
}

func convertEmail(email db.Email) *pb.Email {
	return &pb.Email{
		Id:           email.ID.String(),
		UserId:       email.UserID.String(),
		Subject:      email.Subject,
		Content:      email.Content,
		Title:        email.Title,
		EmailAddress: email.EmailAddress,
	}
}

func convertEmails(emails []db.Email) []*pb.Email {
	res := []*pb.Email{}
	for i := 0; i < len(emails); i++ {
		res = append(res, &pb.Email{
			Id:           emails[i].ID.String(),
			UserId:       emails[i].UserID.String(),
			Subject:      emails[i].Subject,
			Content:      emails[i].Content,
			Title:        emails[i].Title,
			EmailAddress: emails[i].EmailAddress,
		})
	}
	return res
}

func convertToken(accessToken string, expiresAt time.Time) *pb.Token {
	return &pb.Token{
		Token:     accessToken,
		ExpiresAt: timestamppb.New(expiresAt),
	}
}
