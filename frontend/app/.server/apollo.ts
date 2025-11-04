import type { Session } from "react-router";
import { getSession } from "~/.server/session";
import { ApolloClient, InMemoryCache, HttpLink } from "@apollo/client";

export const getGQLClient = async (request: Request) => {
    const session = await getSession(request.headers.get("Cookie"));

    return { client: await getGQLClientFromSession(session), session };
}

export const getGQLClientFromSession = async (session: Session) => {
    const token = session.get("token");

    return new ApolloClient({
        ssrMode: true,
        link: new HttpLink({
            uri: process.env.API_URL,
            headers: token ? {
                Authorization: token,
            } : {},
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
