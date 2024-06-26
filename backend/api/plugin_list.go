//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../../schemas/openapi.yaml

package api

import (
	"context"
	"encoding/json"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/usecase"
	"log/slog"
	"net/http"
	"time"
)

// https://github.com/oapi-codegen/oapi-codegen/blob/master/examples/petstore-expanded/chi/api/petstore.go

type PluginList struct {
	mcPluginUseCase   usecase.MCPluginUseCase
	customDataUseCase usecase.CustomDataUseCase
}

// Make sure we conform to ServerInterface
var _ ServerInterface = (*PluginList)(nil)

// NewPluginList は usecase.MCPluginUseCase と usecase.CustomDataUseCase を使用して OpenAPI Schema に定義された API を実装した PluginList を作成します。
func NewPluginList(mcPluginUseCase usecase.MCPluginUseCase, customDataUseCase usecase.CustomDataUseCase) *PluginList {
	return &PluginList{mcPluginUseCase, customDataUseCase}
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
	plugins, err := p.mcPluginUseCase.GetMCPluginsByServerName(r.Context(), serverName)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get plugins by a server name", slog.String("serverName", serverName), err)
		return
	}

	result := make([]Plugin, len(plugins))
	for i, plugin := range plugins {
		result[i] = toPlugin(plugin)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (p *PluginList) AddPlugins(w http.ResponseWriter, r *http.Request, serverName string) {
	var requests AddPluginsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for Plugin array")
		return
	}

	for _, request := range requests {
		mcPlugin := domain.MCPlugin{
			PluginName:  request.PluginName,
			ServerName:  serverName,
			FileName:    request.FileName,
			Version:     request.Version,
			Type:        request.Type,
			LastUpdated: time.UnixMilli(request.LastUpdated),
		}

		if err := p.mcPluginUseCase.SubmitMCPlugin(r.Context(), mcPlugin); err != nil {
			sendError(w, http.StatusInternalServerError, "Internal server error")
			slog.Error("Failed to process a plugin", slog.Any("request", request), err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *PluginList) AddPlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	var request AddPluginJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for Plugin")
		return
	}

	mcPlugin := domain.MCPlugin{
		PluginName:  pluginName,
		ServerName:  serverName,
		FileName:    request.FileName,
		Version:     request.Version,
		Type:        request.Type,
		LastUpdated: time.UnixMilli(request.LastUpdated),
	}

	if err := p.mcPluginUseCase.SubmitMCPlugin(r.Context(), mcPlugin); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to process a plugin", slog.Any("request", request), err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *PluginList) DeletePlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	if err := p.mcPluginUseCase.DeleteMCPlugin(r.Context(), serverName, pluginName); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to delete a plugin", slog.String("server_name", serverName), slog.String("plugin_name", pluginName), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p *PluginList) GetServerNames(w http.ResponseWriter, r *http.Request) {
	serverNames, err := p.mcPluginUseCase.GetServerNames(r.Context())
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get the list of servers", err)
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
		PluginName:  p.PluginName,
		ServerName:  p.ServerName,
		FileName:    p.FileName,
		Version:     p.Version,
		Type:        p.Type,
		LastUpdated: p.LastUpdated.UnixMilli(),
	}
}

func (p *PluginList) GetCustomDataKeys(w http.ResponseWriter, r *http.Request) {
	keys, err := p.customDataUseCase.GetKeys(r.Context())
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get the custom data keys", err)
		return
	}

	result := make([]CustomDataKey, len(keys))
	for i, key := range keys {
		result[i] = toAPICustomDataKey(key)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (p *PluginList) GetCustomDataKeyInfo(w http.ResponseWriter, r *http.Request, key string) {
	keyInfo, err := p.customDataUseCase.SearchForKey(r.Context(), key)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to search for the custom data key", slog.String("key", key), err)
		return
	}

	if keyInfo == nil {
		sendError(w, http.StatusNotFound, "Key not found")
	} else {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(toAPICustomDataKey(*keyInfo))
	}
}

func (p *PluginList) AddCustomDataKeyInfo(w http.ResponseWriter, r *http.Request, key string) {
	var request AddCustomDataKeyInfoJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for CustomDataKey")
		return
	}

	customDataKey := domain.CustomDataKey{
		Key:         key,
		DisplayName: nonNilOrElse(request.DisplayName, ""),
		Description: nonNilOrElse(request.Description, ""),
		FormType:    nonNilOrElse(request.FormType, "TEXT"),
	}

	if err := p.customDataUseCase.AddOrUpdateKey(r.Context(), customDataKey); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to add or update the custom data key", slog.String("key", key), slog.Any("request", request), err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (p *PluginList) GetPluginNames(w http.ResponseWriter, r *http.Request) {
	pluginNames, err := p.mcPluginUseCase.GetPluginNames(r.Context())
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get the list of plugins", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(pluginNames)
}

func (p *PluginList) GetPluginInfo(w http.ResponseWriter, r *http.Request, pluginName string) {
	installInfo, err := p.mcPluginUseCase.GetInstalledPluginInfo(r.Context(), pluginName)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get the list of installed plugins", slog.String("pluginName", pluginName), err)
		return
	}

	installedServers := make([]Plugin, len(installInfo))
	for i, plugin := range installInfo {
		installedServers[i] = toPlugin(plugin)
	}

	customDataMap, err := p.GetPluginCustomDataMap(r.Context(), pluginName)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get the custom data of the plugin", slog.String("pluginName", pluginName), err)
		return
	}

	result := PluginInfo{
		InstalledServers: &installedServers,
		CustomData:       &customDataMap,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (p *PluginList) GetPluginCustomData(w http.ResponseWriter, r *http.Request, pluginName string) {
	customDataMap, err := p.GetPluginCustomDataMap(r.Context(), pluginName)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to get the custom data of the plugin", slog.String("pluginName", pluginName), err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(customDataMap)
}

func (p *PluginList) GetPluginCustomDataMap(ctx context.Context, pluginName string) (map[string]string, error) {
	pluginInfo, err := p.customDataUseCase.GetPluginInfo(ctx, pluginName)
	if err != nil {
		return nil, err
	}

	customDataMap := make(map[string]string, len(pluginInfo))
	for _, info := range pluginInfo {
		customDataMap[info.Key] = info.Data
	}

	return customDataMap, nil
}

func (p *PluginList) AddPluginCustomData(w http.ResponseWriter, r *http.Request, pluginName string) {
	var request AddPluginCustomDataJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for AddPluginCustomData")
		return
	}

	customData := make([]domain.PluginCustomData, len(request))
	i := 0
	for key, data := range request {
		if exists, err := p.customDataUseCase.ExistsKey(r.Context(), key); err != nil {
			sendError(w, http.StatusInternalServerError, "Internal server error")
			slog.Error("Failed to check the custom data key", slog.String("pluginName", pluginName), slog.String("key", key), slog.Any("data", data), err)
			return
		} else if !exists {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(key)
			return
		}
		customData[i] = domain.PluginCustomData{Key: key, Data: data}
		i++
	}

	if err := p.customDataUseCase.AddOrUpdatePluginInfo(r.Context(), pluginName, customData); err != nil {
		sendError(w, http.StatusInternalServerError, "Internal server error")
		slog.Error("Failed to add or update the custom data of the plugin:", slog.String("pluginName", pluginName), slog.Any("customData", customData), err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

/* Helper methods to convert domain.CustomDataKey <-> api.CustomDataKey */
func toAPICustomDataKey(key domain.CustomDataKey) CustomDataKey {
	return CustomDataKey{
		Key:         key.Key,
		DisplayName: &key.DisplayName,
		Description: &key.Description,
		FormType:    key.FormType,
	}
}

// A function inspired by Java's Objects.requireNonNullElse(T, T)
func nonNilOrElse[T any](val *T, def T) T {
	if val == nil {
		return def
	} else {
		return *val
	}
}
