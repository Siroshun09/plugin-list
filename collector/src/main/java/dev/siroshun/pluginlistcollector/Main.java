package dev.siroshun.pluginlistcollector;

import dev.siroshun.pluginlistcollector.api.BackendAPI;
import dev.siroshun.pluginlistcollector.api.PluginInfo;
import dev.siroshun.pluginlistcollector.function.PluginCollector;
import dev.siroshun.pluginlistcollector.function.PluginListUploader;
import org.bukkit.Bukkit;
import org.bukkit.plugin.java.JavaPlugin;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.slf4j.Logger;
import org.slf4j.helpers.SubstituteLogger;

import java.io.IOException;
import java.nio.file.Path;
import java.util.Map;

public class Main extends JavaPlugin {

    private static final SubstituteLogger LOGGER = new SubstituteLogger("PluginList", null, true);

    public static @NotNull Logger logger() {
        return LOGGER;
    }

    @Override
    public void onLoad() {
        LOGGER.setDelegate(this.getSLF4JLogger());

        this.saveDefaultConfig();
        this.reloadConfig();
    }

    @Override
    public void onEnable() {
        var apiUrl = this.getNonEmptyValue("api-url");
        var token = this.getNonEmptyValue("token");

        if (apiUrl == null || token == null) {
            return;
        }

        Bukkit.getGlobalRegionScheduler().run(this, ignored -> this.collectPluginAndUpload(new BackendAPI(apiUrl, token)));
    }

    private void collectPluginAndUpload(@NotNull BackendAPI api) {
        var serverName = Path.of(".").toAbsolutePath().getParent().getFileName().toString();
        var plugins = PluginCollector.collect(this.getServer().getPluginManager(), serverName);

        this.getServer().getAsyncScheduler().runNow(this, ignored -> this.uploadPlugins(api, serverName, plugins));
    }

    private void uploadPlugins(@NotNull BackendAPI api, @NotNull String serverName, @NotNull Map<String, PluginInfo> plugins) {
        try {
            PluginListUploader.upload(api, serverName, plugins);
        } catch (IOException | InterruptedException e) {
            this.getSLF4JLogger().error("An exception occurred while uploading plugins.", e);
        }
    }

    private @Nullable String getNonEmptyValue(@NotNull String path) {
        var value = this.getConfig().getString(path);

        if (value == null || value.isEmpty()) {
            logger().error("{} is not set!", path);
            return null;
        }

        return value;
    }
}
