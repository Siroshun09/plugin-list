package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
	"time"
)

type mcPluginRepository struct {
	conn *sqliteConnection
}

const (
	mcPluginTableSchema = `
		CREATE TABLE IF NOT EXISTS mc_plugins (
			plugin_name VARCHAR(32) NOT NULL,
			server_name VARCHAR(16) NOT NULL,
			filename VARCHAR(64) NOT NULL,
			version VARCHAR(32) NOT NULL,
			type VARCHAR(16) NOT NULL,
			last_updated INTEGER NOT NULL,
			PRIMARY KEY (plugin_name, server_name)
		)
		`

	insertOrUpdateMcPluginQuery = `
		INSERT INTO mc_plugins (plugin_name, server_name, filename, version, type, last_updated) VALUES (?, ?, ?, ?, ?, ?) 
		ON CONFLICT (plugin_name, server_name) DO UPDATE SET filename = excluded.filename, version = excluded.version,
		type = excluded.type, last_updated = excluded.last_updated
		`

	deleteMcPluginQuery = `DELETE FROM mc_plugins WHERE plugin_name = ? AND server_name = ?`

	selectMCPluginsByServerNameQuery = "SELECT * FROM mc_plugins WHERE server_name=?"

	selectServerNamesQuery = "SELECT DISTINCT server_name FROM mc_plugins"

	selectPluginNamesQuery = "SELECT DISTINCT plugin_name FROM mc_plugins"

	selectMCPluginsByPluginNameQuery = "SELECT * FROM mc_plugins WHERE plugin_name = ?"
)

func (c *sqliteConnection) NewMCPluginRepository(ctx context.Context) (repository.MCPluginRepository, error) {
	if _, err := c.db.ExecContext(ctx, mcPluginTableSchema); err != nil {
		return nil, err
	}

	return mcPluginRepository{c}, nil
}

func (m mcPluginRepository) CreateOrUpdateMCPlugin(ctx context.Context, plugin domain.MCPlugin) (returnErr error) {
	rows, err := m.conn.db.QueryContext(ctx, insertOrUpdateMcPluginQuery, plugin.PluginName, plugin.ServerName, plugin.FileName, plugin.Version, plugin.Type, plugin.LastUpdated.UnixMilli())
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

func (m mcPluginRepository) DeleteMCPlugin(ctx context.Context, serverName string, pluginName string) (returnErr error) {
	rows, err := m.conn.db.QueryContext(ctx, deleteMcPluginQuery, pluginName, serverName)
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

func (m mcPluginRepository) GetMCPluginsByServerName(ctx context.Context, serverName string) (plugins []domain.MCPlugin, returnErr error) {
	rows, err := m.conn.db.QueryContext(ctx, selectMCPluginsByServerNameQuery, serverName)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	var result []domain.MCPlugin
	for rows.Next() {
		var plugin domain.MCPlugin
		var unixTime int64
		if err := rows.Scan(&plugin.PluginName, &plugin.ServerName, &plugin.FileName, &plugin.Version, &plugin.Type, &unixTime); err != nil {
			return nil, err
		}
		plugin.LastUpdated = time.UnixMilli(unixTime)
		result = append(result, plugin)
	}

	return result, nil
}

func (m mcPluginRepository) GetServerNames(ctx context.Context) (serverNames []string, returnErr error) {
	rows, err := m.conn.db.QueryContext(ctx, selectServerNamesQuery)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	var result []string
	for rows.Next() {
		var serverName string
		if err := rows.Scan(&serverName); err != nil {
			return nil, err
		}
		result = append(result, serverName)
	}

	return result, nil
}

func (m mcPluginRepository) GetPluginNames(ctx context.Context) (pluginNames []string, returnErr error) {
	rows, err := m.conn.db.QueryContext(ctx, selectPluginNamesQuery)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	var result []string

	for rows.Next() {
		var serverName string
		if err := rows.Scan(&serverName); err != nil {
			return nil, err
		}
		result = append(result, serverName)
	}

	return result, nil
}

func (m mcPluginRepository) GetInstalledPluginInfo(ctx context.Context, pluginName string) (plugins []domain.MCPlugin, returnErr error) {
	rows, err := m.conn.db.QueryContext(ctx, selectMCPluginsByPluginNameQuery, pluginName)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	var result []domain.MCPlugin

	for rows.Next() {
		var plugin domain.MCPlugin
		var unixTime int64
		if err := rows.Scan(&plugin.PluginName, &plugin.ServerName, &plugin.FileName, &plugin.Version, &plugin.Type, &unixTime); err != nil {
			return nil, err
		}
		plugin.LastUpdated = time.UnixMilli(unixTime)
		result = append(result, plugin)
	}

	return result, nil
}
