import { defineConfig } from 'orval';

export default defineConfig({
    backend: {
        input: {
            target: "../schemas/openapi.yaml",
        },
        output: {
            target: "./src/api/backend.ts",
            clean: true,
            client: "react-query",
        },
    },
});
