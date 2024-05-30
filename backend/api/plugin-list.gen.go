// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
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

// AddPluginsJSONBody defines parameters for AddPlugins.
type AddPluginsJSONBody = []struct {
	// FileName File name of the plugin
	FileName *string `json:"file_name,omitempty"`

	// LastUpdated Unix time when the plugin was last updated (milliseconds)
	LastUpdated *int64 `json:"last_updated,omitempty"`

	// PluginName Name of the plugin
	PluginName *string `json:"plugin_name,omitempty"`

	// Type Type of the plugin
	Type *string `json:"type,omitempty"`

	// Version Version of the plugin
	Version *string `json:"version,omitempty"`
}

// AddPluginJSONBody defines parameters for AddPlugin.
type AddPluginJSONBody struct {
	// FileName File name of the plugin
	FileName *string `json:"file_name,omitempty"`

	// LastUpdated Unix time when the plugin was last updated (milliseconds)
	LastUpdated *int64 `json:"last_updated,omitempty"`

	// Type Type of the plugin
	Type *string `json:"type,omitempty"`

	// Version Version of the plugin
	Version *string `json:"version,omitempty"`
}

// AddPluginsJSONRequestBody defines body for AddPlugins for application/json ContentType.
type AddPluginsJSONRequestBody = AddPluginsJSONBody

// AddPluginJSONRequestBody defines body for AddPlugin for application/json ContentType.
type AddPluginJSONRequestBody AddPluginJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
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

	"H4sIAAAAAAAC/+xXUW/bNhD+K8RtDxugRHaStYPfXKwbggJbgKTDhiAIWPFkM5FIlTwl8QL994GkLMmS",
	"nDrJui1F/WSJ5N3x7rvvPt1DovNCK1RkYXYPBm2hlUX/8F7xkpbayL9QvDVGG/dSoE2MLEhqBTPgSYLW",
	"MtLXqJi0LJfWSrVg2jCpbngmBVRVBDZZYs690cZQYXSBhmTwlWiBQ/N+M/NrEaTa5JxgBlLR4QFEQKsC",
	"wyMu0EAVQY7W8sVWQ+vl5qglI9XCh2jwYykNCpidQ+1wvf2iiuAkKxdSOcM8y35LYXbev0EqM7xUPB/x",
	"/rPMkLklplNGS2RFsDaII4KMW7osC8HJhdI39F7JO0YyR3a7RNUxxW65Ze4sq8+y73KZZdJiopWw3/fS",
	"9+poNH3B1pZL/LpT/BbNDZodTISNYybCi/7Zs1Wxg/sbNNbv7x//PSx8ykIPCN2EbN4t6tS7dVvb65Xx",
	"orqIAO94XmTYQwqcoaWArb3p/mT/ipv+6dn09fSHH6evX03cr1ekzvlegDMgtHRZv5q2V/1QXl9Lumwy",
	"0KQMpvsTqHy7YlIaSatT17YB3Weuw4dpndetT5olBjlhjEJSLDBDWmfZshvJ2fzkGCJwLQRL5MKXvg70",
	"j735yfHeu7d/tkHyQr7DVYhGqlQPPQfbe5m0VNsmSS6/IysbV3SXjEAXqHghYQaH+5P9Q4ig4LT0V41D",
	"zmzsHhZIQ9+/IHkUeRc6rbFswZs13O06FmHfqV9ywHfLG/R6MJkE5lOEynvhRZHJxJ+Pr2yAcaBO969B",
	"0PmwtJ3nA7iIQBLm3smW7gJuDK/zu3m3sXv5TSkvM3pUwN8aTGEG38TtjInrSRCHMTDiv1R4V2DiKAzr",
	"PRHYMs+5WT2QererKdx9pxGquEbhztVco5aWnBg3yKSyxLMMBZOBc22BiUwlipbFBoUPTWnfrE7XWwpu",
	"eI7konXTY5j2YCwMCtJsgbSOBTr8MSi+7ymH3rajNpmqZTQyJUadCvXZ7+KZEG1g91Dpa7p6JBafUJX/",
	"MWzb4NcVduNX2xF4zoVwciqMg87wGtLNXIiTZumloO1jiZbeaLF6GtD+AQ3W3nV0GP+LEq2JpDfyP4t0",
	"G7v2M+VYa7IvMp4v1FrbTqaMCbf6jf5whQntQi8na0rRjHebbADiakCL02HgtvTfQs7xUVgfI5LGTjz8",
	"uvrv+arWfZ4xasV37tVry2Tb6ejTIzi+78C0Cil0QnGYzJ+CgNyk9bqPUqPzhkwHJBhONnB+Cg06weka",
	"7mEYfkZijMYCrW//+EA32nskys2PnOeIhaNtMp2FOosvvTl2h+2jx/32af8yhv0LwvTTJMlXJRKUyEvV",
	"C0MamDOFt+ucCU78qzB4gjBwTqq/AwAA//+UFLM8aBYAAA==",
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