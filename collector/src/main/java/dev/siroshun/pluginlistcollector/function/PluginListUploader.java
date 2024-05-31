package dev.siroshun.pluginlistcollector.function;

import dev.siroshun.pluginlistcollector.api.BackendAPI;
import dev.siroshun.pluginlistcollector.api.PluginInfo;
import dev.siroshun.pluginlistcollector.api.APIResult;
import org.jetbrains.annotations.NotNull;

import java.io.IOException;
import java.util.List;
import java.util.Map;

import static dev.siroshun.pluginlistcollector.Main.logger;

public final class PluginListUploader {

    public static void upload(@NotNull BackendAPI api, @NotNull String serverName, @NotNull Map<String, PluginInfo> plugins) throws IOException, InterruptedException {
        var previouslyUploadedPlugins = switch (api.getPlugins(serverName)) {
            case APIResult.Success<List<PluginInfo>> success -> success.value();
            case APIResult.CodecError<List<PluginInfo>> error -> {
                logger().warn("Failed to decode the list of previously uploaded plugins. Detailed result: {}", error.failure());
                yield null;
            }
            case APIResult.BackendError<List<PluginInfo>> error -> {
                logger().warn("Failed to get the list of previously uploaded plugins. Detailed response: {}", error);
                yield null;
            }
        };

        if (previouslyUploadedPlugins == null) {
            return;
        }

        var uploadResult = api.uploadPlugins(serverName, List.copyOf(plugins.values()));

        if (uploadResult instanceof APIResult.CodecError<Void> error) {
            logger().warn("Failed to encode the list of plugins. Detailed result: {}", error);
            return;
        } else if (uploadResult instanceof APIResult.BackendError<Void> error) {
            logger().warn("Failed to upload the list of plugins. Detailed response: {}", error);
            return;
        }

        if (previouslyUploadedPlugins.isEmpty()) {
            return;
        }

        for (var plugin : previouslyUploadedPlugins) {
            if (plugins.containsKey(plugin.name())) {
                continue;
            }

            if (api.deletePlugin(serverName, plugin.name()) instanceof APIResult.BackendError<Void> error) {
                logger().error("Failed to delete the plugin '{}'. Detailed response: {}", plugin.name(), error);
                return;
            }
        }

        logger().info("{} plugin(s) uploaded!", plugins.size());
    }
}
