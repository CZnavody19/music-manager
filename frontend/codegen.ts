import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
    schema: "http://localhost:8080/",
    documents: ['src/**/*.tsx'],
    generates: {
        './app/graphql/gen/': {
            preset: 'client',
            plugins: [],
            presetConfig: {
                gqlTagName: 'gql',
            },
            config: {
                scalars: {
                    Time: 'string',
                },
            }
        }
    },
    ignoreNoDocuments: true,
};

export default config;