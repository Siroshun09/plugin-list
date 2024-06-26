// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

const (
	TokenScopes = "Token.Scopes"
)

// CustomDataKey defines model for CustomDataKey.
type CustomDataKey struct {
	// Description the description of the key
	Description *string `json:"description,omitempty"`

	// DisplayName the display name of the key
	DisplayName *string `json:"display_name,omitempty"`

	// FormType the form type that is used in frontend
	FormType string `json:"form_type"`

	// Key the key
	Key string `json:"key"`
}

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// Plugin defines model for Plugin.
type Plugin struct {
	// FileName File name of the plugin
	FileName string `json:"file_name"`

	// LastUpdated Unix time when the plugin was last updated (milliseconds)
	LastUpdated int64 `json:"last_updated"`

	// PluginName Name of the plugin
	PluginName string `json:"plugin_name"`

	// ServerName Name of the server
	ServerName string `json:"server_name"`

	// Type Type of the plugin
	Type string `json:"type"`

	// Version Version of the plugin
	Version string `json:"version"`
}

// PluginInfo defines model for PluginInfo.
type PluginInfo struct {
	// CustomData User-defined information of this plugin
	CustomData *map[string]string `json:"custom_data,omitempty"`

	// InstalledServers Servers that have this plugin. This also provide data sent by its server.
	InstalledServers *[]Plugin `json:"installed_servers,omitempty"`
}

// AddCustomDataKeyInfoJSONBody defines parameters for AddCustomDataKeyInfo.
type AddCustomDataKeyInfoJSONBody struct {
	// Description the description of the key
	Description *string `json:"description,omitempty"`

	// DisplayName the display name of the key
	DisplayName *string `json:"display_name,omitempty"`

	// FormType the form type that is used in frontend (If not specified, this value will be "TEXT")
	FormType *string `json:"form_type,omitempty"`
}

// AddPluginCustomDataJSONBody defines parameters for AddPluginCustomData.
type AddPluginCustomDataJSONBody map[string]string

// AddPluginsJSONBody defines parameters for AddPlugins.
type AddPluginsJSONBody = []struct {
	// FileName File name of the plugin
	FileName string `json:"file_name"`

	// LastUpdated Unix time when the plugin was last updated (milliseconds)
	LastUpdated int64 `json:"last_updated"`

	// PluginName Name of the plugin
	PluginName string `json:"plugin_name"`

	// Type Type of the plugin
	Type string `json:"type"`

	// Version Version of the plugin
	Version string `json:"version"`
}

// AddPluginJSONBody defines parameters for AddPlugin.
type AddPluginJSONBody struct {
	// FileName File name of the plugin
	FileName string `json:"file_name"`

	// LastUpdated Unix time when the plugin was last updated (milliseconds)
	LastUpdated int64 `json:"last_updated"`

	// Type Type of the plugin
	Type string `json:"type"`

	// Version Version of the plugin
	Version string `json:"version"`
}

// AddCustomDataKeyInfoJSONRequestBody defines body for AddCustomDataKeyInfo for application/json ContentType.
type AddCustomDataKeyInfoJSONRequestBody AddCustomDataKeyInfoJSONBody

// AddPluginCustomDataJSONRequestBody defines body for AddPluginCustomData for application/json ContentType.
type AddPluginCustomDataJSONRequestBody AddPluginCustomDataJSONBody

// AddPluginsJSONRequestBody defines body for AddPlugins for application/json ContentType.
type AddPluginsJSONRequestBody = AddPluginsJSONBody

// AddPluginJSONRequestBody defines body for AddPlugin for application/json ContentType.
type AddPluginJSONRequestBody AddPluginJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get the custom data keys
	// (GET /custom_data/keys/)
	GetCustomDataKeys(w http.ResponseWriter, r *http.Request)
	// Get information of the custom data key
	// (GET /custom_data/keys/{key}/)
	GetCustomDataKeyInfo(w http.ResponseWriter, r *http.Request, key string)
	// Add or update information of the custom data key
	// (POST /custom_data/keys/{key}/)
	AddCustomDataKeyInfo(w http.ResponseWriter, r *http.Request, key string)
	// Get the list of known plugins
	// (GET /plugins/)
	GetPluginNames(w http.ResponseWriter, r *http.Request)
	// Get the detailed information of the plugin
	// (GET /plugins/{plugin_name})
	GetPluginInfo(w http.ResponseWriter, r *http.Request, pluginName string)
	// Get the custom data of the plugin
	// (GET /plugins/{plugin_name}/custom-data/)
	GetPluginCustomData(w http.ResponseWriter, r *http.Request, pluginName string)
	// Add or update custom data of the plugin
	// (POST /plugins/{plugin_name}/custom-data/)
	AddPluginCustomData(w http.ResponseWriter, r *http.Request, pluginName string)
	// Get the list of servers
	// (GET /servers/)
	GetServerNames(w http.ResponseWriter, r *http.Request)
	// Get the list of installed plugins
	// (GET /servers/{server_name}/plugins)
	GetPluginsByServer(w http.ResponseWriter, r *http.Request, serverName string)
	// Add or update the plugins
	// (POST /servers/{server_name}/plugins)
	AddPlugins(w http.ResponseWriter, r *http.Request, serverName string)
	// Delete the specified plugin from the list
	// (DELETE /servers/{server_name}/plugins/{plugin_name})
	DeletePlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string)
	// Add or update the plugin
	// (POST /servers/{server_name}/plugins/{plugin_name})
	AddPlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get the custom data keys
// (GET /custom_data/keys/)
func (_ Unimplemented) GetCustomDataKeys(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get information of the custom data key
// (GET /custom_data/keys/{key}/)
func (_ Unimplemented) GetCustomDataKeyInfo(w http.ResponseWriter, r *http.Request, key string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add or update information of the custom data key
// (POST /custom_data/keys/{key}/)
func (_ Unimplemented) AddCustomDataKeyInfo(w http.ResponseWriter, r *http.Request, key string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get the list of known plugins
// (GET /plugins/)
func (_ Unimplemented) GetPluginNames(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get the detailed information of the plugin
// (GET /plugins/{plugin_name})
func (_ Unimplemented) GetPluginInfo(w http.ResponseWriter, r *http.Request, pluginName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get the custom data of the plugin
// (GET /plugins/{plugin_name}/custom-data/)
func (_ Unimplemented) GetPluginCustomData(w http.ResponseWriter, r *http.Request, pluginName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add or update custom data of the plugin
// (POST /plugins/{plugin_name}/custom-data/)
func (_ Unimplemented) AddPluginCustomData(w http.ResponseWriter, r *http.Request, pluginName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get the list of servers
// (GET /servers/)
func (_ Unimplemented) GetServerNames(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get the list of installed plugins
// (GET /servers/{server_name}/plugins)
func (_ Unimplemented) GetPluginsByServer(w http.ResponseWriter, r *http.Request, serverName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add or update the plugins
// (POST /servers/{server_name}/plugins)
func (_ Unimplemented) AddPlugins(w http.ResponseWriter, r *http.Request, serverName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete the specified plugin from the list
// (DELETE /servers/{server_name}/plugins/{plugin_name})
func (_ Unimplemented) DeletePlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add or update the plugin
// (POST /servers/{server_name}/plugins/{plugin_name})
func (_ Unimplemented) AddPlugin(w http.ResponseWriter, r *http.Request, serverName string, pluginName string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetCustomDataKeys operation middleware
func (siw *ServerInterfaceWrapper) GetCustomDataKeys(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCustomDataKeys(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCustomDataKeyInfo operation middleware
func (siw *ServerInterfaceWrapper) GetCustomDataKeyInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", chi.URLParam(r, "key"), &key, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCustomDataKeyInfo(w, r, key)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddCustomDataKeyInfo operation middleware
func (siw *ServerInterfaceWrapper) AddCustomDataKeyInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", chi.URLParam(r, "key"), &key, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddCustomDataKeyInfo(w, r, key)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPluginNames operation middleware
func (siw *ServerInterfaceWrapper) GetPluginNames(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPluginNames(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPluginInfo operation middleware
func (siw *ServerInterfaceWrapper) GetPluginInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "plugin_name" -------------
	var pluginName string

	err = runtime.BindStyledParameterWithOptions("simple", "plugin_name", chi.URLParam(r, "plugin_name"), &pluginName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "plugin_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPluginInfo(w, r, pluginName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPluginCustomData operation middleware
func (siw *ServerInterfaceWrapper) GetPluginCustomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "plugin_name" -------------
	var pluginName string

	err = runtime.BindStyledParameterWithOptions("simple", "plugin_name", chi.URLParam(r, "plugin_name"), &pluginName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "plugin_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPluginCustomData(w, r, pluginName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddPluginCustomData operation middleware
func (siw *ServerInterfaceWrapper) AddPluginCustomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "plugin_name" -------------
	var pluginName string

	err = runtime.BindStyledParameterWithOptions("simple", "plugin_name", chi.URLParam(r, "plugin_name"), &pluginName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "plugin_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPluginCustomData(w, r, pluginName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetServerNames operation middleware
func (siw *ServerInterfaceWrapper) GetServerNames(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetServerNames(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPluginsByServer operation middleware
func (siw *ServerInterfaceWrapper) GetPluginsByServer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "server_name" -------------
	var serverName string

	err = runtime.BindStyledParameterWithOptions("simple", "server_name", chi.URLParam(r, "server_name"), &serverName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "server_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPluginsByServer(w, r, serverName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddPlugins operation middleware
func (siw *ServerInterfaceWrapper) AddPlugins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "server_name" -------------
	var serverName string

	err = runtime.BindStyledParameterWithOptions("simple", "server_name", chi.URLParam(r, "server_name"), &serverName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "server_name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, TokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPlugins(w, r, serverName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeletePlugin operation middleware
func (siw *ServerInterfaceWrapper) DeletePlugin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "server_name" -------------
	var serverName string

	err = runtime.BindStyledParameterWithOptions("simple", "server_name", chi.URLParam(r, "server_name"), &serverName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "server_name", Err: err})
		return
	}

	// ------------- Path parameter "plugin_name" -------------
	var pluginName string

	err = runtime.BindStyledParameterWithOptions("simple", "plugin_name", chi.URLParam(r, "plugin_name"), &pluginName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "plugin_name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, TokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeletePlugin(w, r, serverName, pluginName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddPlugin operation middleware
func (siw *ServerInterfaceWrapper) AddPlugin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "server_name" -------------
	var serverName string

	err = runtime.BindStyledParameterWithOptions("simple", "server_name", chi.URLParam(r, "server_name"), &serverName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "server_name", Err: err})
		return
	}

	// ------------- Path parameter "plugin_name" -------------
	var pluginName string

	err = runtime.BindStyledParameterWithOptions("simple", "plugin_name", chi.URLParam(r, "plugin_name"), &pluginName, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "plugin_name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, TokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPlugin(w, r, serverName, pluginName)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/custom_data/keys/", wrapper.GetCustomDataKeys)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/custom_data/keys/{key}/", wrapper.GetCustomDataKeyInfo)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/custom_data/keys/{key}/", wrapper.AddCustomDataKeyInfo)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/plugins/", wrapper.GetPluginNames)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/plugins/{plugin_name}", wrapper.GetPluginInfo)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/plugins/{plugin_name}/custom-data/", wrapper.GetPluginCustomData)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/plugins/{plugin_name}/custom-data/", wrapper.AddPluginCustomData)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/servers/", wrapper.GetServerNames)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/servers/{server_name}/plugins", wrapper.GetPluginsByServer)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/servers/{server_name}/plugins", wrapper.AddPlugins)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/servers/{server_name}/plugins/{plugin_name}", wrapper.DeletePlugin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/servers/{server_name}/plugins/{plugin_name}", wrapper.AddPlugin)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xab0/cOBP/Kpaf50UrBRYoT/to39Fr74Qq9ZBKTz1RhEw8Yd1N7NR2oDmU736ynf9x",
	"dgNL6XK9vuomzng885uZ33i4xaFIUsGBa4Xnt1iCSgVXYH985CTTCyHZX0DfSimkeUhBhZKlmgmO55iE",
	"ISiFtFgCR0yhhCnF+BUSEjF+TWJGcVEEWIULSIgV+kumtEjeEE3eQW4ekDj+PcLzs1ucSpGC1Mzt3tmo",
	"v69eAGo9QSJC5tESchxgnaeA51hpyfgVLgJMmUpjkl9wksCILLcCmRVrhEVCJhfuqU+SeY3Ma6QXRBub",
	"ZAooYhxFUnANnPqELp0thuK8ShQBlvA1YxIonp9ht6bR67w4DzB8I0kae5Q8XQDKZIxuFiDBHjSNsyvm",
	"3EcY14RxMEp2rYY/yhiXiuJMxlaLGhVd14WCeja2i5F957QlGs8x4/rFQXNExjVcgTQ2SUApcjUqqHq9",
	"zjrlhtXy8yLAJ/bAq8AXsRhG4PIri6GDE2c+n1djovRFllKijSp9QR85+4Y0S8C4grc9cUMUMt+i8lv0",
	"LGFxzBSEglP1vGe+l4de8zlZI4d4P0l/BfIa5AQRbqFPhD9QTk18rN3+GqTyRv8f7sU6CT0gtA3SPVvQ",
	"8nezbSmv58Z+cLWQgk9BaYetnf3dvd0vRPa/nu+/2v/f//dfvdwz/3pOan3fU3CONSh9UT7ab456mS2X",
	"TF/UFqhNhvd396wFnLxjHolVeA9tVr6gRBO7jFJmbE3ik86yYWbtIlqB3KEQmQSCGHcgrR3F1MBT4vIL",
	"hNoIYlxpEsdAy0MOSwD+4F64vLog19CWuYtOzQ8SK4FSKa4ZBWROgxRwjS5zxLQqYbqLA8w0JHaL/0qI",
	"8Bz/Z9bUwVlZrWalK2oUYyIlyXFR9CDQM14/3TJlUiupgjsSEhlvGgsGNpPO8ULrVM1ns1LobiiSmfV4",
	"WqvgMdDZlqPvvLC1H8JMMp1/MFZ1MDo1dGFoqqOSR2iBQglEwwwo0zMKMegqyhW6ZgQdnRwbJ5qPFkCo",
	"TT2lqp92jk6Od969/bNRk6TMcA2rDSsDobuzk70TM6VL2Zpp41zPm84hzTEDLFLgJGV4jl/smpMHOCV6",
	"YY86a4FjtoRczczTK9BDJX4DbdOZ+8Kh13yB7QbSRtIxdQs7LMqs6BC3g709V4YN37A7kTSNWWhFzL4o",
	"l1MdzM3/JkVDl7h5giLw0JfBWeyqiGSxvpOGqxRzJMSjQMbhWwqhKaBQrgmwypKEyHyVvc2yod9ul5AX",
	"q703SHkD4Wt9aRO1wY8kCegqzNcaFrcSks0pZXQYHDax4RY2NVHLDIKWlfv183xDXN0BTtPhc7h3OLT+",
	"EnLEhUaRyDjdNpBNgIXha0J5YHVEqWmmXAa/D8COKN1ygH3NQOnXguZ3ctZP0iyiZ8eRRbZKIWQRAxo4",
	"2nNN4gzQDYtjdAnoMz59++n0M37uZcE9tjUE8Xu4mYatrmuLQX7YH55TZfaOYJuC8s5BZStCSUDWF3BL",
	"FkSEllzc8Jq3WA8TCUiCaUkpEq7nM5tcEgW+0uDomOm0Nq7xdfSetYie4XLNrwN83mLGI63c6nrvPfo2",
	"Fv0RRdt+vm3R42Kt0ylowmJf39NqUEccPDUnlx2EzTtaoCt/bal3axJ2h9p78na3Nd4OgtAyzQjYJph8",
	"C5E3RetRGJa0dMfS0jt1EhOh2HCFewJy3b7bD8rvfvmx4mr2Ye4K1tb71U7a8h5toOwU5jw9FI4ofYBQ",
	"IJN3f6SAuB/LnhwLPx7ThsOusvNm1PXw4OB+dAtn3HKMC/8kZ3CMehZQugxdCpqjBVGoFFT1i9tJplcE",
	"qilq5d3ldA5dXXZ6Cpa7D35ocjy46mz9fkCCXJ1ri6lxrWLbcbeti+Gi4iiTvTnohepbbdPz2klS1eY2",
	"M6URpqJe5x+qJWuzsxPWISqV7u0M3He+Jwl350aPx0o2nFiswuI9vLLFsG2Ub7WfUyhCk6vUOClQTwht",
	"9yv5NdAeYCLu4zat4dQjDsxrTXojsO8ySB+hdJsMxxuR/aHb5mPzRrYbXj3KGN0zA16Xs06qPNWj15uT",
	"O/fel51qObPhX0P9+CRYjlZtGiqHqmf2DxTG6Fk7x62v68MrMDeLHRrzjZvRdmtFRfilSOoMPcis7ss6",
	"Ru6TW6fdgX3HbBs80cu6w7FJOHJ+pv/04JgO2ztziHEK8TQYxBPC9EMMEH9eevMkScgmRGOYP44Qh5vK",
	"2NRd+P3LKO7KKMwmxd8BAAD//45lzcBRLgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
