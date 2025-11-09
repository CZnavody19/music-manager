import { ApolloClient, HttpLink, InMemoryCache } from "@apollo/client";
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';

export function getGQLWebsocketClient(apiURL: string) {
    return new ApolloClient({
        link: typeof window !== "undefined" ? new GraphQLWsLink(createClient({
            url: apiURL,
        })) : new HttpLink({
            uri: apiURL,
        }),
        cache: new InMemoryCache(),
        defaultOptions: {
            watchQuery: {
                fetchPolicy: 'no-cache',
                errorPolicy: 'all',
            },
            query: {
                fetchPolicy: 'no-cache',
                errorPolicy: 'all',
            },
        },
    });
}