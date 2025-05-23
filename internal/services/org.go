package services

import (
	"context"
	"fmt"

	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/pkg/utils"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type OrgService struct{}

func NewOrgService() *OrgService {
	return &OrgService{}
}

func (s *OrgService) GetOrg(id string) (*db.OrganizationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get org",
		"params": id,
	}).Msg("DB Query")

	user, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).With(db.Organization.Owner.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *OrgService) GetUserOrgs(id string) ([]db.OrganizationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get user orgs",
		"params": id,
	}).Msg("DB Query")
	user, err := prisma.Client.Organization.FindMany(
		db.Organization.OwnerID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *OrgService) SearchOrg(name string) ([]db.OrganizationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "search org",
		"params": name,
	}).Msg("DB Query")
	users, err := prisma.Client.Organization.FindMany(
		db.Organization.Name.Contains(name),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *OrgService) UpdateOrg(id string, payload types.OrgPayload) (*db.OrganizationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "update org",
		"id":     id,
		"params": payload,
	}).Msg("DB Query")
	users, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Update(
		db.Organization.Name.Set(payload.Name),
		db.Organization.Description.Set(payload.Description),
			).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *OrgService) CreateOrg(id string, payload types.OrgPayload) (*db.OrganizationModel, error) {
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":   "create org",
		"ownerId": id,
		"params":  payload,
	}).Msg("DB Query")
	ctx := context.Background()
	users, err := prisma.Client.Organization.CreateOne(
		db.Organization.Name.Set(payload.Name),
		db.Organization.Description.Set(payload.Description),
		db.Organization.Image.Set("uplodas/profiles/default.jpg"),
		db.Organization.BgImage.Set("uplodas/bgs/default.jpg"),
		db.Organization.Owner.Link(db.User.ID.Equals(id)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *OrgService) UpdateOrgImage(id string, path string) (string, error) {
	fmt.Println(id)
	ctx := context.Background()
	user, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Update(
		db.Organization.Image.Set(path),
	).Exec(ctx)
	if err != nil {
		return "", err
	}
	return user.Image, nil
}

func (s *OrgService) UpdateOrgBg(id string, path string) (string, error) {
	fmt.Println(id)
	ctx := context.Background()
	user, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Update(
		db.Organization.BgImage.Set(path),
	).Exec(ctx)
	if err != nil {
		return "", err
	}
	return user.Image, nil
}

func (s *OrgService) DeleteOrg(id string) (string, error) {
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "delete org",
		"params": id,
	}).Msg("DB Query")
	ctx := context.Background()
	deleted, err := prisma.Client.Organization.FindUnique(
		db.Organization.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return "", nil
	}
	fmt.Println(deleted.ID)
	return deleted.ID, nil
}
