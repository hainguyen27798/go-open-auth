package repos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/query"
	"github.com/open-auth/pkg/utils"
	"go.uber.org/zap"
)

type IPermissionRepo interface {
	CreateNewPermission(payload models.InsertNewPermissionParams) error
	GetAllPermissions() []models.Permission
	SearchPermissions(search string, by string, skip int, limit int) ([]models.Permission, int64)
	GetPermission(id string) *models.Permission
	UpdatePermission(permission models.UpdatePermissionParams) (bool, error)
	DeletePermission(id string) bool
	GetPermissionOptions(roleId string) []models.Permission
}

type permissionRepo struct {
	sqlX *sqlx.DB
}

func NewPermissionRepo() IPermissionRepo {
	return &permissionRepo{
		sqlX: global.Mdb,
	}
}

func (pr *permissionRepo) CreateNewPermission(payload models.InsertNewPermissionParams) error {
	session, err := utils.NewTransaction(pr.sqlX)
	if err != nil {
		return err
	}

	if _, err := session.NamedExecCommit(query.InsertNewPermission, payload); err != nil {
		return err
	}

	return nil
}

func (pr *permissionRepo) GetPermission(id string) *models.Permission {
	var permission models.Permission
	err := pr.sqlX.Get(&permission, query.GetPermissionById, id)
	if err != nil {
		return nil
	}
	return &permission
}

func (pr *permissionRepo) GetAllPermissions() []models.Permission {
	var permissions []models.Permission
	if err := pr.sqlX.Select(&permissions, query.GetAllPermissions); err != nil {
		global.Logger.Error("GetAllPermission: ", zap.Error(err))
		return []models.Permission{}
	}
	return permissions
}

func (pr *permissionRepo) SearchPermissions(search string, by string, skip int, limit int) ([]models.Permission, int64) {
	var permissions []models.Permission
	var total int64
	queryString := query.SearchPermissionsBy[by]
	queryCount := query.CountPermissionSearchBy[by]
	search = "%" + search + "%"

	if queryString == "" {
		queryString = query.SearchPermissionsBy["service_name"]
		queryCount = query.CountPermissionSearchBy["service_name"]
	}

	if err := pr.sqlX.Select(&permissions, queryString, search, limit, skip); err != nil {
		global.Logger.Error("GetAllPermission: ", zap.Error(err))
		return []models.Permission{}, 0
	}

	if err := pr.sqlX.Get(&total, queryCount, search); err != nil {
		global.Logger.Error("CountPermission: ", zap.Error(err))
		return []models.Permission{}, 0
	}

	return permissions, total
}

func (pr *permissionRepo) UpdatePermission(payload models.UpdatePermissionParams) (bool, error) {
	permission := pr.GetPermission(*payload.ID)

	if permission == nil {
		return false, nil
	}

	querySet := utils.PartialUpdate(payload)

	queryString := fmt.Sprintf(query.UpdatePermission, "SET "+querySet)

	session, err := utils.NewTransaction(pr.sqlX)
	if err != nil {
		return false, err
	}

	_, err = session.NamedExecCommit(queryString, payload)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (pr *permissionRepo) DeletePermission(id string) bool {
	session, err := utils.NewTransaction(pr.sqlX)
	if err != nil {
		return false
	}

	count, err := session.ExecCommit(query.DeletePermission, id)
	if err != nil {
		return false
	}
	return count > 0
}

func (pr *permissionRepo) GetPermissionOptions(roleId string) []models.Permission {
	var permissions []models.Permission
	err := pr.sqlX.Select(&permissions, query.SelectPermissionOptions, roleId)
	if err != nil {
		global.Logger.Error("GetPermissionOptions: ", zap.Error(err))
		return []models.Permission{}
	}
	return permissions
}
