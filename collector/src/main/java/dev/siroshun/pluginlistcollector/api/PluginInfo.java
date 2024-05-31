package dev.siroshun.pluginlistcollector.api;

import dev.siroshun.codec4j.api.codec.collection.ListCodec;
import dev.siroshun.codec4j.api.codec.object.FieldCodec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodecFactory;

import java.time.Instant;

import static dev.siroshun.codec4j.api.codec.ValueCodecs.LONG_CODEC;
import static dev.siroshun.codec4j.api.codec.ValueCodecs.STRING_CODEC;

public record PluginInfo(String name, String server, String file, String version, String type, Instant lastUpdated) {

    public static final ListCodec<PluginInfo> LIST_CODEC = ListCodec.unmodifiableList(
        ObjectCodecFactory.create(
            FieldCodec.getter(PluginInfo::name).name("plugin_name").codec(STRING_CODEC).build(),
            FieldCodec.getter(PluginInfo::server).name("server_name").codec(STRING_CODEC).build(),
            FieldCodec.getter(PluginInfo::file).name("file_name").codec(STRING_CODEC).build(),
            FieldCodec.getter(PluginInfo::version).name("version").codec(STRING_CODEC).build(),
            FieldCodec.getter(PluginInfo::type).name("type").codec(STRING_CODEC).build(),
            FieldCodec.getter(PluginInfo::lastUpdated).name("last_updated").codec(LONG_CODEC.map(Instant::toEpochMilli, Instant::ofEpochMilli)).build(),
            PluginInfo::new
        )
    );

}
