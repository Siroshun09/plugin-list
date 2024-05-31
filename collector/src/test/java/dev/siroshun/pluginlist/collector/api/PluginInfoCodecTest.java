package dev.siroshun.pluginlist.collector.api;

import dev.siroshun.pluginlistcollector.api.PluginInfo;
import dev.siroshun.codec4j.io.gson.GsonIO;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;

import java.io.IOException;
import java.time.Instant;
import java.util.List;
import java.util.stream.Stream;

class PluginInfoCodecTest {

    private static final PluginInfo TEST_PLUGIN = new PluginInfo("TestPlugin", "test-server", "TestPlugin-1.0.jar", "1.0", "bukkit_plugin", Instant.ofEpochMilli(100));
    private static final String TEST_PLUGIN_JSON = "{\"plugin_name\":\"TestPlugin\",\"server_name\":\"test-server\",\"file_name\":\"TestPlugin-1.0.jar\",\"version\":\"1.0\",\"type\":\"bukkit_plugin\",\"last_updated\":100}";

    @ParameterizedTest
    @MethodSource("testCases")
    void testEncodeAndDecode(TestCase testCase) throws IOException {
        var encodeResult = GsonIO.toJson(testCase.plugins, PluginInfo.LIST_CODEC);
        Assertions.assertTrue(encodeResult.isSuccess(), encodeResult::toString);

        var json = encodeResult.asSuccess().result();
        Assertions.assertEquals(testCase.json, json);

        var decodeResult = GsonIO.fromJson(json, PluginInfo.LIST_CODEC);
        Assertions.assertTrue(decodeResult.isSuccess(), decodeResult::toString);
        Assertions.assertEquals(testCase.plugins, decodeResult.asSuccess().result());
    }

    private static Stream<TestCase> testCases() {
        return Stream.of(
            new TestCase(List.of(), "[]"),
            new TestCase(List.of(TEST_PLUGIN), "[" + TEST_PLUGIN_JSON + "]"),
            new TestCase(List.of(TEST_PLUGIN, TEST_PLUGIN), "[" + TEST_PLUGIN_JSON + "," + TEST_PLUGIN_JSON + "]")
        );
    }

    private record TestCase(List<PluginInfo> plugins, String json) {
    }
}
