package pg

import "polx/app/datastore"

func ClearAll() error {
	_, err := datastore.RwInstance().Exec("TRUNCATE trades, users, user_notifications")
	return err
}
