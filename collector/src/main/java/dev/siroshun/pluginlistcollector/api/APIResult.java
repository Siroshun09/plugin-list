package dev.siroshun.pluginlistcollector.api;

import dev.siroshun.codec4j.api.result.Result;
import org.jetbrains.annotations.NotNull;

public sealed interface APIResult<T> permits APIResult.BackendError, APIResult.CodecError, APIResult.Success {

    record Success<T>(T value) implements APIResult<T> {
    }

    record CodecError<T>(@NotNull Result.Failure<T> failure) implements APIResult<T> {
    }

    record BackendError<T>(int statusCode, @NotNull String message) implements APIResult<T> {
    }
}
