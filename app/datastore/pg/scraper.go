package pg

import (
	"context"
	"database/sql"
	"polx/app/datastore"
	"polx/app/domain/bo"
	"polx/app/domain/definition"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	scraperOnce sync.Once
	scraperInst *scraperRepo
)

type scraperRepo struct {
	db *sql.DB
}

func GetScraperRepo() definition.ScraperRepo {
	scraperOnce.Do(func() {
		scraperInst = &scraperRepo{
			db: datastore.RwInstance(),
		}
	})
	return scraperInst
}

func (s *scraperRepo) GetShills(ctx context.Context, query string) ([]bo.Shill, error) {
	log := logrus.New()
	statement := "SELECT shill_name FROM shills WHERE shill_name LIKE '$1%' LIMIT 5"
	res, err := s.db.QueryContext(ctx, statement, query)

	if err != nil {
		log.Errorf("%s", err.Error())
		return nil, err
	}

	var shills []bo.Shill
	for res.Next() {
		var shill bo.Shill
		if err := res.Scan(
			&shill.Name,
		); err != nil {
			log.Errorf("%s", err.Error())
			return nil, err
		}

		shills = append(shills, shill)
	}
	return shills, nil
}
