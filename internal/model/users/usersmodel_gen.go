// Code generated by goctl. DO NOT EDIT!

package users

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheUsersIdPrefix     = "cache:users:id:"
	cacheUsersMobilePrefix = "cache:users:mobile:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Users, error)
		FindOneByMobile(ctx context.Context, mobile string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id          uint64         `db:"id"`          // 自增ID
		Name        string         `db:"name"`        // 用户名
		Mobile      string         `db:"mobile"`      // 手机号
		Password    sql.NullString `db:"password"`    // 密码
		Avatar      string         `db:"avatar"`      // 头像
		Status      int64          `db:"status"`      // 状态
		Extra       sql.NullString `db:"extra"`       // 额外信息
		CreatedTime int64          `db:"createdTime"` // 创建时间
		UpdatedTime int64          `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullInt64  `db:"deletedTime"` // 删除时间
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	usersMobileKey := fmt.Sprintf("%s%v", cacheUsersMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usersIdKey, usersMobileKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id uint64) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByMobile(ctx context.Context, mobile string) (*Users, error) {
	usersMobileKey := fmt.Sprintf("%s%v", cacheUsersMobilePrefix, mobile)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersMobileKey := fmt.Sprintf("%s%v", cacheUsersMobilePrefix, data.Mobile)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Mobile, data.Password, data.Avatar, data.Status, data.Extra, data.CreatedTime, data.UpdatedTime, data.DeletedTime)
	}, usersIdKey, usersMobileKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersMobileKey := fmt.Sprintf("%s%v", cacheUsersMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Mobile, newData.Password, newData.Avatar, newData.Status, newData.Extra, newData.CreatedTime, newData.UpdatedTime, newData.DeletedTime, newData.Id)
	}, usersIdKey, usersMobileKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
