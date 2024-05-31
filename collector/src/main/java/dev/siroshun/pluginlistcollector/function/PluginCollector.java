package dev.siroshun.pluginlistcollector.function;

import dev.siroshun.pluginlistcollector.api.PluginInfo;
import dev.siroshun.pluginlistcollector.Main;
import org.bukkit.plugin.Plugin;
import org.bukkit.plugin.PluginDescriptionFile;
import org.bukkit.plugin.PluginManager;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.attribute.FileTime;
import java.util.Arrays;
import java.util.Map;
import java.util.Objects;
import java.util.function.Function;
import java.util.stream.Collectors;

public final class PluginCollector {

    public static @NotNull Map<String, PluginInfo> collect(@NotNull PluginManager manager, @NotNull String serverName) {
        return Arrays.stream(manager.getPlugins())
            .map(plugin -> createPluginInfo(plugin, serverName))
            .filter(Objects::nonNull)
            .collect(Collectors.toMap(PluginInfo::name, Function.identity()));
    }

    @SuppressWarnings("UnstableApiUsage")
    private static @Nullable PluginInfo createPluginInfo(@NotNull Plugin plugin, String serverName) {
        var jarFile = jarFile(plugin);

        if (jarFile == null) {
            return null;
        }

        FileTime lastModified;

        try {
            lastModified = Files.getLastModifiedTime(jarFile);
        } catch (IOException e) {
            Main.logger().warn("Cannot get the last modified time of the file {}", jarFile, e);
            return null;
        }

        return new PluginInfo(
            plugin.getName(),
            serverName,
            jarFile.getFileName().toString(),
            plugin.getPluginMeta().getVersion(),
            plugin.getPluginMeta() instanceof PluginDescriptionFile ? "bukkit_plugin" : "paper_plugin",
            lastModified.toInstant()
        );
    }

    private static @Nullable Path jarFile(@NotNull Plugin plugin) {
        var location = plugin.getClass().getProtectionDomain().getCodeSource().getLocation();
        if (location != null) {
            try {
                return Path.of(location.toURI());
            } catch (URISyntaxException ignored) {
            }
        }
        return null;
    }
}
