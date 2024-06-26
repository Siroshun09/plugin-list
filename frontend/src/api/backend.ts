/**
 * Generated by orval v6.28.2 🍺
 * Do not edit manually.
 * plugin-list API
 * plugin-list API
 * OpenAPI spec version: 1.0.0
 */
import axios from 'axios'
import type {
  AxiosRequestConfig,
  AxiosResponse
} from 'axios'
export type AddCustomDataKeyInfoBody = {
  /** the description of the key */
  description?: string;
  /** the display name of the key */
  display_name?: string;
  /** the form type that is used in frontend (If not specified, this value will be "TEXT") */
  form_type?: string;
};

export type AddPluginCustomDataBody = {[key: string]: string};

/**
 * User-defined information of this plugin
 */
export type GetPluginCustomData200 = {[key: string]: string};

export type AddPluginBody = {
  /** File name of the plugin */
  file_name: string;
  /** Unix time when the plugin was last updated (milliseconds) */
  last_updated: number;
  /** Type of the plugin */
  type: string;
  /** Version of the plugin */
  version: string;
};

export type AddPluginsBodyItem = {
  /** File name of the plugin */
  file_name: string;
  /** Unix time when the plugin was last updated (milliseconds) */
  last_updated: number;
  /** Name of the plugin */
  plugin_name: string;
  /** Type of the plugin */
  type: string;
  /** Version of the plugin */
  version: string;
};

/**
 * access token is missing or invalid
 */
export type UnauthorizedErrorResponse = void;

export interface Error {
  /** Error code */
  code: number;
  /** Error message */
  message: string;
}

export type CustomDataKeyAllOf = {
  /** the description of the key */
  description?: string;
  /** the display name of the key */
  display_name?: string;
  /** the form type that is used in frontend */
  form_type: string;
  /** the key */
  key: string;
};

export type CustomDataKey = CustomDataKeyAllOf;

/**
 * User-defined information of this plugin
 */
export type PluginInfoAllOfCustomData = {[key: string]: string};

export type PluginInfoAllOf = {
  /** User-defined information of this plugin */
  custom_data?: PluginInfoAllOfCustomData;
  /** Servers that have this plugin. This also provide data sent by its server. */
  installed_servers?: Plugin[];
};

export type PluginInfo = PluginInfoAllOf;

export type PluginAllOf = {
  /** File name of the plugin */
  file_name: string;
  /** Unix time when the plugin was last updated (milliseconds) */
  last_updated: number;
  /** Name of the plugin */
  plugin_name: string;
  /** Name of the server */
  server_name: string;
  /** Type of the plugin */
  type: string;
  /** Version of the plugin */
  version: string;
};

export type Plugin = PluginAllOf;





  /**
 * Get the list of servers
 * @summary Get the list of servers
 */
export const getServerNames = <TData = AxiosResponse<string[]>>(
     options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/servers/`,options
    );
  }

/**
 * Get the list of plugins that are installed in the specified server
 * @summary Get the list of installed plugins
 */
export const getPluginsByServer = <TData = AxiosResponse<Plugin[]>>(
    serverName: string, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/servers/${serverName}/plugins`,options
    );
  }

/**
 * Add or update the plugins
 * @summary Add or update the plugins
 */
export const addPlugins = <TData = AxiosResponse<void>>(
    serverName: string,
    addPluginsBodyItem: AddPluginsBodyItem[], options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.post(
      `/servers/${serverName}/plugins`,
      addPluginsBodyItem,options
    );
  }

/**
 * Add or update the plugin
 * @summary Add or update the plugin
 */
export const addPlugin = <TData = AxiosResponse<void>>(
    serverName: string,
    pluginName: string,
    addPluginBody: AddPluginBody, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.post(
      `/servers/${serverName}/plugins/${pluginName}`,
      addPluginBody,options
    );
  }

/**
 * Delete the specified plugin from the list
 * @summary Delete the specified plugin from the list
 */
export const deletePlugin = <TData = AxiosResponse<void>>(
    serverName: string,
    pluginName: string, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.delete(
      `/servers/${serverName}/plugins/${pluginName}`,options
    );
  }

/**
 * Get the list of known plugins that are recoded on the database
 * @summary Get the list of known plugins
 */
export const getPluginNames = <TData = AxiosResponse<string[]>>(
     options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/plugins/`,options
    );
  }

/**
 * Get the detailed information of the plugin
 * @summary Get the detailed information of the plugin
 */
export const getPluginInfo = <TData = AxiosResponse<PluginInfo>>(
    pluginName: string, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/plugins/${pluginName}`,options
    );
  }

/**
 * Get the custom data of the plugin
 * @summary Get the custom data of the plugin
 */
export const getPluginCustomData = <TData = AxiosResponse<GetPluginCustomData200>>(
    pluginName: string, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/plugins/${pluginName}/custom-data/`,options
    );
  }

/**
 * Add or update custom data of the plugin
 * @summary Add or update custom data of the plugin
 */
export const addPluginCustomData = <TData = AxiosResponse<void>>(
    pluginName: string,
    addPluginCustomDataBody: AddPluginCustomDataBody, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.post(
      `/plugins/${pluginName}/custom-data/`,
      addPluginCustomDataBody,options
    );
  }

/**
 * Get the custom data keys
 * @summary Get the custom data keys
 */
export const getCustomDataKeys = <TData = AxiosResponse<CustomDataKey[]>>(
     options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/custom_data/keys/`,options
    );
  }

/**
 * Get information of the custom data key
 * @summary Get information of the custom data key
 */
export const getCustomDataKeyInfo = <TData = AxiosResponse<CustomDataKey>>(
    key: string, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/custom_data/keys/${key}/`,options
    );
  }

/**
 * Add or update information of the custom data key
 * @summary Add or update information of the custom data key
 */
export const addCustomDataKeyInfo = <TData = AxiosResponse<void>>(
    key: string,
    addCustomDataKeyInfoBody: AddCustomDataKeyInfoBody, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.post(
      `/custom_data/keys/${key}/`,
      addCustomDataKeyInfoBody,options
    );
  }

export type GetServerNamesResult = AxiosResponse<string[]>
export type GetPluginsByServerResult = AxiosResponse<Plugin[]>
export type AddPluginsResult = AxiosResponse<void>
export type AddPluginResult = AxiosResponse<void>
export type DeletePluginResult = AxiosResponse<void>
export type GetPluginNamesResult = AxiosResponse<string[]>
export type GetPluginInfoResult = AxiosResponse<PluginInfo>
export type GetPluginCustomDataResult = AxiosResponse<GetPluginCustomData200>
export type AddPluginCustomDataResult = AxiosResponse<void>
export type GetCustomDataKeysResult = AxiosResponse<CustomDataKey[]>
export type GetCustomDataKeyInfoResult = AxiosResponse<CustomDataKey>
export type AddCustomDataKeyInfoResult = AxiosResponse<void>
