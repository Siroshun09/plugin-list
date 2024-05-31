package dev.siroshun.pluginlistcollector.api;

import dev.siroshun.codec4j.api.result.Result;
import dev.siroshun.codec4j.io.gson.GsonIO;
import org.jetbrains.annotations.NotNull;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import java.util.Collection;
import java.util.List;

public class BackendAPI {

    private static final HttpClient HTTP_CLIENT = HttpClient.newHttpClient();

    private final String baseUrl;
    private final String token;

    public BackendAPI(@NotNull String baseUrl, @NotNull String token) {
        this.baseUrl = baseUrl;
        this.token = token;
    }

    public @NotNull APIResult<List<PluginInfo>> getPlugins(@NotNull String serverName) throws IOException, InterruptedException {
        var request = HttpRequest.newBuilder().uri(URI.create(this.baseUrl + "/servers/" + serverName + "/plugins")).GET().build();
        var response = HTTP_CLIENT.send(request, HttpResponse.BodyHandlers.ofInputStream());

        if (response.statusCode() != 200) {
            return toErrorResponse(response);
        }

        try (var reader = GsonIO.newIn(new InputStreamReader(response.body(), StandardCharsets.UTF_8))) {
            return toResponse(PluginInfo.LIST_CODEC.decode(reader));
        }
    }

    public @NotNull APIResult<Void> uploadPlugins(@NotNull String serverName, @NotNull Collection<PluginInfo> plugins) throws IOException, InterruptedException {
        var body = GsonIO.toBytes(plugins, PluginInfo.LIST_CODEC.asIteratorEncoder());

        if (body.isFailure()) {
            return toResponse(body.asFailure().cast());
        }

        return processResponse(
            this.createRequest("/servers/" + serverName + "/plugins")
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofInputStream(() -> new ByteArrayInputStream(body.asSuccess().result())))
                .build()
        );
    }

    public @NotNull APIResult<Void> deletePlugin(@NotNull String serverName, @NotNull String pluginName) throws IOException, InterruptedException {
        return processResponse(this.createRequest("/servers/" + serverName + "/plugins/" + pluginName).DELETE().build());
    }

    private static @NotNull APIResult<Void> processResponse(HttpRequest request) throws IOException, InterruptedException {
        var response = HTTP_CLIENT.send(request, HttpResponse.BodyHandlers.ofInputStream());
        return switch (response.statusCode()) {
            case 201, 204 -> new APIResult.Success<>(null);
            case 401 -> new APIResult.BackendError<>(401, "Unauthorized");
            default -> toErrorResponse(response);
        };
    }

    private @NotNull HttpRequest.Builder createRequest(@NotNull String api) {
        return HttpRequest.newBuilder().header("X-API-KEY", this.token).uri(URI.create(this.baseUrl + api));
    }

    private static <T> @NotNull APIResult<T> toErrorResponse(@NotNull HttpResponse<InputStream> response) {
        try (var input = response.body()) {
            return new APIResult.BackendError<>(response.statusCode(), new String(input.readAllBytes(), StandardCharsets.UTF_8));
        } catch (IOException ignored) {
            return new APIResult.BackendError<>(response.statusCode(), "Invalid error response");
        }
    }

    private static <T> @NotNull APIResult<T> toResponse(@NotNull Result<T> result) {
        return result.isSuccess() ? new APIResult.Success<>(result.asSuccess().result()) : new APIResult.CodecError<>(result.asFailure());
    }
}
