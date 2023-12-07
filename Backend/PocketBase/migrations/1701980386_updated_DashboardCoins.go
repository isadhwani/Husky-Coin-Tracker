package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ojln7eb289s9sm4")
		if err != nil {
			return err
		}

		collection.Name = "Coins"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ojln7eb289s9sm4")
		if err != nil {
			return err
		}

		collection.Name = "DashboardCoins"

		return dao.SaveCollection(collection)
	})
}
