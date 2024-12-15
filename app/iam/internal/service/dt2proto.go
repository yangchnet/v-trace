package service

import (
	"database/sql"
	"time"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func User2Proto(user *db.User, role string) *v1.User {
	if user == nil {
		return &v1.User{}
	}

	return &v1.User{
		Id:       pbtools.ToProtoInt64(user.ID),
		Username: pbtools.ToProtoString(user.Username),
		Nickname: pbtools.ToProtoString(user.Nickname.String),
		Role:     pbtools.ToProtoString(role),
		Phone:    pbtools.ToProtoString(user.Phone),
		Email:    pbtools.ToProtoString(user.Email.String),
		CreateAt: pbtools.ToProtoTimestamp(user.CreatedAt.Time),
		Realname: pbtools.ToProtoString(user.Realname.String),
		Idcard:   pbtools.ToProtoString(user.Idcard.String),
	}
}

func Org2Proto(org *db.Org) *v1.Org {
	return &v1.Org{
		Id:               pbtools.ToProtoInt64(org.ID),
		OrgName:          pbtools.ToProtoString(org.OrgName),
		OrgCode:          pbtools.ToProtoString(org.OrgCode),
		LegalPersonName:  pbtools.ToProtoString(org.LegalPersonName),
		LegalPersonPhone: pbtools.ToProtoString(org.LegalPersonPhone),
		Owner:            pbtools.ToProtoString(org.Owner),
		OrgInfo:          pbtools.ToProtoBytes(org.Info),
	}
}

func Proto2Org(org *v1.Org) *db.Org {
	return &db.Org{
		ID: int32(org.Id.Value),
		CreatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: true,
		},
		Owner:            org.Owner.Value,
		Info:             org.OrgInfo.Value,
		OrgName:          org.OrgName.Value,
		OrgCode:          org.OrgCode.Value,
		LegalPersonName:  org.LegalPersonName.Value,
		LegalPersonPhone: org.LegalPersonPhone.Value,
	}
}

func Proto2User(user *v1.User) *db.User {
	return &db.User{
		ID:       int32(user.Id.Value),
		Username: user.Username.Value,
		Nickname: sql.NullString{
			String: user.Username.Value,
			Valid:  true,
		},
		Passwd: "",
		Email: sql.NullString{
			String: user.Email.Value,
			Valid:  true,
		},
		CreatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: true,
		},
		Realname: sql.NullString{
			String: user.Realname.Value,
			Valid:  true,
		},
		Idcard: sql.NullString{
			String: user.Idcard.Value,
			Valid:  true,
		},
		Phone: user.Phone.Value,
	}
}
