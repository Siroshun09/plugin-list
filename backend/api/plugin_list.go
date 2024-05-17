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

// NewPluginList は usecase.MCPluginUseCase を使用して OpenAPI Schema に定義された API を実装した PluginList を作成します。
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

	result := make([]Plugin, len(plugins))
	for i, plugin := range plugins {
		converted := toPlugin(plugin)
		result[i] = converted
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (p *PluginList) AddPlugins(w http.ResponseWriter, r *http.Request, serverName string) {
	var plugins []Plugin

	if err := json.NewDecoder(r.Body).Decode(&plugins); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for Plugin array")
		return
	}

	for _, plugin := range plugins {
		plugin.ServerName = serverName

		mcPlugin := toMCPlugin(plugin)

		if err := p.useCase.SubmitMCPlugin(r.Context(), mcPlugin); err != nil {
			sendError(w, http.StatusInternalServerError, "Internal server error")
			slog.Error("Failed to get plugins by serverName:", "request", &plugin, err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *PluginList) AddPlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	var plugin Plugin

	if err := json.NewDecoder(r.Body).Decode(&plugin); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for Plugin")
		return
	}

	plugin.PluginName = pluginName
	plugin.ServerName = serverName

	mcPlugin := toMCPlugin(plugin)

	if err := p.useCase.SubmitMCPlugin(r.Context(), mcPlugin); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get plugins by serverName:", "request", plugin, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *PluginList) DeletePlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	if err := p.useCase.DeleteMCPlugin(r.Context(), serverName, pluginName); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get plugins by serverName:", slog.String("server_name", serverName), slog.String("plugin_name", pluginName), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p *PluginList) GetServerNames(w http.ResponseWriter, r *http.Request) {
	serverNames, err := p.useCase.GetServerNames(r.Context())
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to the list of servers:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(serverNames)
}

/* Helper methods to convert MCPlugin <-> Plugin */

func toMCPlugin(plugin Plugin) domain.MCPlugin {
	return domain.MCPlugin{
		PluginName:  plugin.PluginName,
		ServerName:  plugin.ServerName,
		FileName:    plugin.FileName,
		Version:     plugin.Version,
		Type:        plugin.Type,
		LastUpdated: time.UnixMilli(plugin.LastUpdated),
	}
}

func toPlugin(p domain.MCPlugin) Plugin {
	return Plugin{
		FileName:    p.FileName,
		LastUpdated: p.LastUpdated.UnixMilli(),
		PluginName:  p.PluginName,
		ServerName:  p.ServerName,
		Type:        p.Type,
		Version:     p.Version,
	}
}
