import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import type { Route } from "./+types/tidal";
import { TidalConfigForm, resolver, type formType } from "~/components/forms/tidal";
import { useFormResponse } from "~/hooks/use-form-response";
import { getGQLClient } from "~/.server/apollo";
import type { Query } from "~/graphql/gen/graphql";
import { gql } from "@apollo/client";

export function meta() {
    return [
        { title: "Tidal - Settings - Music Manager" },
        { name: "description", content: "Manage your music collection with ease." },
    ];
}

export async function action({ request }: Route.ActionArgs) {
    const { errors, data } = await getValidatedFormData(request, resolver);

    if (errors) {
        return { errors };
    }

    const { client } = await getGQLClient(request);

    const { error } = await client.mutate({
        mutation: gql`
            mutation setTidalConfig($config: TidalConfigInput!) {
                setTidalConfig(config: $config)
            }
        `,
        variables: {
            config: data,
        },
    });

    return { errors: error };
}

export async function loader({ request }: Route.LoaderArgs) {
    const { client } = await getGQLClient(request);

    const { data, error } = await client.query<Query>({
        query: gql`
            query getTidalConfig {
                getTidalConfig {
                    authTokenType
                    authAccessToken
                    authRefreshToken
                    authExpiresAt
                    authClientID
                    authClientSecret
                    downloadTimeout
                    downloadRetries
                    downloadThreads
                    audioQuality
                }
            }
        `,
    });

    return { errors: error, data: data?.getTidalConfig };
}

export default function Page({ loaderData, actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
        defaultValues: loaderData?.data ?? {},
    });
    useFormResponse<typeof resolver>(actionData);

    return (
        <TidalConfigForm form={form} />
    )
}