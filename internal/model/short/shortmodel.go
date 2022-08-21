package short

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShortModel = (*customShortModel)(nil)

type (
	// ShortModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortModel.
	ShortModel interface {
		shortModel
	}

	customShortModel struct {
		*defaultShortModel
	}
)

// NewShortModel returns a model for the database table.
func NewShortModel(conn sqlx.SqlConn, c cache.CacheConf) ShortModel {
	return &customShortModel{
		defaultShortModel: newShortModel(conn, c),
	}
}
