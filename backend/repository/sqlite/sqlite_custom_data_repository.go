package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
)

type customDataRepository struct {
	conn *sqliteConnection
}

const (
	customDataKeyTableSchema = `
		CREATE TABLE IF NOT EXISTS custom_data_keys (
			key VARCHAR(32) NOT NULL PRIMARY KEY,
			description VARCHAR(255) NOT NULL,
			display_name VARCHAR(255) NOT NULL,
			form_type VARCHAR(255) NOT NULL
		)
	`

	pluginCustomDataTableSchema = `
		CREATE TABLE IF NOT EXISTS plugin_custom_data (
			plugin_name VARCHAR(32) NOT NULL,
			key VARCHAR(32) NOT NULL,
			data VARCHAR(255) NOT NULL,
			PRIMARY KEY (plugin_name, key),
			FOREIGN KEY (key) REFERENCES custom_data_keys(key)
		)
	`

	selectAllKeys = `SELECT key, description, display_name, form_type FROM custom_data_keys`

	selectSpecifiedKey = `SELECT description, display_name, form_type FROM custom_data_keys WHERE key=?`

	insertOrUpdateKey = `
		INSERT INTO custom_data_keys (key, description, display_name, form_type) VALUES (?, ?, ?, ?) 
		ON CONFLICT (key) DO UPDATE SET description = excluded.description, display_name = excluded.display_name, form_type = excluded.form_type
	`

	checkExistingKey = "SELECT EXISTS(SELECT TRUE FROM custom_data_keys WHERE key=?) AS customer_check;"

	selectPluginCustomData = `SELECT key, data FROM plugin_custom_data WHERE plugin_name=?`

	insertOrUpdatePluginCustomData = `
		INSERT INTO plugin_custom_data (plugin_name, key, data) VALUES (?, ?, ?) 
		ON CONFLICT (plugin_name, key) DO UPDATE SET data = excluded.data
	`
)

func (c *sqliteConnection) NewCustomDataRepository(ctx context.Context) (repository.CustomDataRepository, error) {
	if _, err := c.db.ExecContext(ctx, customDataKeyTableSchema); err != nil {
		return nil, err
	}

	if _, err := c.db.ExecContext(ctx, pluginCustomDataTableSchema); err != nil {
		return nil, err
	}

	return customDataRepository{c}, nil
}

func (c customDataRepository) GetKeys(ctx context.Context) (_ []domain.CustomDataKey, returnErr error) {
	rows, err := c.conn.db.QueryContext(ctx, selectAllKeys)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	var result []domain.CustomDataKey
	for rows.Next() {
		var key domain.CustomDataKey
		if err = rows.Scan(&key.Key, &key.Description, &key.DisplayName, &key.FormType); err != nil {
			return nil, err
		}
		result = append(result, key)
	}

	return result, nil
}

func (c customDataRepository) SearchForKey(ctx context.Context, key string) (_ *domain.CustomDataKey, returnErr error) {
	rows, err := c.conn.db.QueryContext(ctx, selectSpecifiedKey, key)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	if rows.Next() {
		result := &domain.CustomDataKey{Key: key}
		if err = rows.Scan(&result.Description, &result.DisplayName, &result.FormType); err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return nil, nil
	}
}

func (c customDataRepository) AddOrUpdateKey(ctx context.Context, key domain.CustomDataKey) (returnErr error) {
	rows, err := c.conn.db.QueryContext(ctx, insertOrUpdateKey, key.Key, key.Description, key.DisplayName, key.FormType)
	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	return nil
}

func (c customDataRepository) ExistsKey(ctx context.Context, key string) (result bool, returnErr error) {
	rows, err := c.conn.db.QueryContext(ctx, checkExistingKey, key)
	if err != nil {
		return false, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	if rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return false, err
		}
		return result, nil
	}
	return false, nil
}

func (c customDataRepository) GetPluginInfo(ctx context.Context, pluginName string) (_ []domain.PluginCustomData, returnErr error) {
	rows, err := c.conn.db.QueryContext(ctx, selectPluginCustomData, pluginName)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	var result []domain.PluginCustomData
	for rows.Next() {
		var data domain.PluginCustomData
		if err = rows.Scan(&data.Key, &data.Data); err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (c customDataRepository) AddOrUpdatePluginInfo(ctx context.Context, pluginName string, data domain.PluginCustomData) (returnErr error) {
	rows, err := c.conn.db.QueryContext(ctx, insertOrUpdatePluginCustomData, pluginName, data.Key, data.Data)
	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	return nil
}
