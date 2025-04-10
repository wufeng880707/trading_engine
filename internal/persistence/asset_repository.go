package persistence

import (
	"context"

	"github.com/duolacloud/crud-core/repositories"
	models "github.com/yzimhao/trading_engine/v2/internal/models/asset"
	"github.com/yzimhao/trading_engine/v2/internal/models/types"
	"github.com/yzimhao/trading_engine/v2/internal/persistence/database/entities"
	"gorm.io/gorm"
)

type AssetRepository interface {
	repositories.CrudRepository[models.Asset, models.CreateAsset, models.UpdateAsset]
	Despoit(ctx context.Context, transId, userId, symbol string, amount types.Numeric) error
	Withdraw(ctx context.Context, transId, userId, symbol string, amount types.Numeric) error
	Transfer(ctx context.Context, transId, from, to, symbol string, amount types.Numeric) error
	TransferWithTx(ctx context.Context, tx *gorm.DB, transId, from, to, symbol string, amount types.Numeric) error
	Freeze(ctx context.Context, tx *gorm.DB, transId, userId, symbol string, amount types.Numeric) (*entities.AssetFreeze, error)
	UnFreeze(ctx context.Context, tx *gorm.DB, transId, userId, symbol string, amount types.Numeric) error
	QueryFreeze(ctx context.Context, filter map[string]any) (assetFreezes []*models.AssetFreeze, err error)
}
