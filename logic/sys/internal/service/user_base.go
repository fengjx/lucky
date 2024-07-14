package service

import (
	"context"
	"strings"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"
	"github.com/fengjx/go-halo/json"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen/log"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/connom/kit"
	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var UserBaseSvc = &userBaseService{}

type userBaseService struct {
}

// Query 分页查询
func (svc userBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[entity.SysUser], error) {
	readDB := dao.SysUserDao.GetReadDB()
	query.TableName = meta.SysUserMeta.TableName()
	list, page, err := daox.Find[entity.SysUser](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query user err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &types.PageVO[entity.SysUser]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc userBaseService) Add(ctx context.Context, sysUser *entity.SysUser) (int64, error) {
	md5Pwd, salt := svc.genPwd(sysUser.Pwd)
	sysUser.Salt = salt
	sysUser.Pwd = md5Pwd
	return dao.SysUserDao.SaveContext(ctx, sysUser)
}

// Update 更新记录
func (svc userBaseService) Update(ctx context.Context, sysUser *entity.SysUser) (bool, error) {
	return dao.SysUserDao.UpdateContext(ctx, sysUser,
		meta.SysUserMeta.PrimaryKey(),
		meta.SysUserMeta.Pwd,
		meta.SysUserMeta.Salt,
		meta.SysUserMeta.Ctime,
		meta.SysUserMeta.Utime,
	)
}

// BatchUpdate 批量更新
func (svc userBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.SysUserMeta.PrimaryKey() {
				id = v
				continue
			}
			if k == meta.SysUserMeta.Pwd {
				md5Pwd, salt := svc.genPwd(v.(string))
				v = md5Pwd
				attr[meta.SysUserMeta.Salt] = salt
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.SysUserDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc userBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.SysUserDao.Deleter().Where(
		ql.C(meta.SysUserMeta.IdIn(ids...)),
	).ExecContext(ctx)
	if err != nil {
		l.Error("delete user err", zap.Error(err))
		return err
	}
	l.Info("delete user success")
	return nil
}

// UpdatePwd 修改用户密码
func (svc userBaseService) UpdatePwd(ctx context.Context, id int64, newPwd string) error {
	l := log.GetLogger(ctx).With(zap.Any("id", id))
	pwd, salt := svc.genPwd(newPwd)
	_, err := dao.SysUserDao.UpdateFieldContext(ctx, id, map[string]any{
		meta.SysUserMeta.Pwd:  pwd,
		meta.SysUserMeta.Salt: salt,
	})
	if err != nil {
		l.Error("update sys_user pwd err", zap.Error(err))
	}
	return err
}

func (svc userBaseService) genPwd(pwd string) (md5Pwd, salt string) {
	if pwd == "" {
		return
	}
	salt = utils.RandomString(6)
	sb := strings.Builder{}
	sb.WriteString(pwd)
	sb.WriteString(salt)
	md5Pwd = kit.MD5Hash(sb.String())
	return
}

func (svc userBaseService) Get(ctx context.Context, uid int64) (*entity.SysUser, error) {
	if uid == 0 {
		return nil, nil
	}
	user := &entity.SysUser{}
	exist, err := dao.SysUserDao.GetByIDContext(ctx, uid, user)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, err
	}
	return user, nil
}
