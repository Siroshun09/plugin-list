//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../../schemas/openapi.yaml

package api

import (
	"encoding/json"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/usecase"
	"log/slog"
	"net/http"
	"time"
)

// https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/chi/api/petstore.go

type PluginList struct {
	useCase usecase.MCPluginUseCase
}

// Make sure we conform to ServerInterface
var _ ServerInterface = (*PluginList)(nil)

func NewPluginList(useCase usecase.MCPluginUseCase) *PluginList {
	return &PluginList{useCase}
}

// sendError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendError(w http.ResponseWriter, code int, message string) {
	err := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}

func (p *PluginList) GetPluginsByServer(w http.ResponseWriter, r *http.Request, serverName string) {
	plugins, err := p.useCase.GetMCPluginsByServerName(r.Context(), serverName)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get plugins by serverName:", slog.String("serverName", serverName), err)
		return
	}

	var result []Plugin
	for _, p := range plugins {
		result = append(result, toPlugin(p))
	}

	w.WriteHeader(http.StatusOK)
	if result != nil {
		_ = json.NewEncoder(w).Encode(result)
	} else {
		_ = json.NewEncoder(w).Encode(make([]Plugin, 0))
	}

}

func (p *PluginList) AddPlugin(w http.ResponseWriter, r *http.Request, serverName string) {
	var plugin Plugin

	if err := json.NewDecoder(r.Body).Decode(&plugin); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for Plugin")
		return
	}

	if plugin.ServerName != serverName {
		sendError(w, http.StatusBadRequest, "ServerName mismatch")
		return
	}

	mcPlugin := toMCPlugin(&plugin)

	if err := p.useCase.SubmitMCPlugin(r.Context(), &mcPlugin); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get plugins by serverName:", "request", plugin, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(plugin)
}

func (p *PluginList) DeletePlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	if err := p.useCase.DeleteMCPlugin(r.Context(), serverName, pluginName); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get plugins by serverName:", slog.String("server_name", serverName), slog.String("plugin_name", pluginName), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

/* Helper methods to convert MCPlugin <-> Plugin */

func toMCPlugin(plugin *Plugin) domain.MCPlugin {
	return domain.MCPlugin{
		PluginName:  plugin.PluginName,
		ServerName:  plugin.ServerName,
		FileName:    plugin.FileName,
		Version:     plugin.Version,
		Type:        plugin.Type,
		LastUpdated: time.UnixMilli(plugin.LastUpdated),
	}
}

func toPlugin(p *domain.MCPlugin) Plugin {
	return Plugin{
		FileName:    p.FileName,
		LastUpdated: p.LastUpdated.UnixMilli(),
		PluginName:  p.PluginName,
		ServerName:  p.ServerName,
		Type:        p.Type,
		Version:     p.Version,
	}
}
