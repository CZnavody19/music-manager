import type { Route } from "./+types/plex";
import { useFormResponse } from "~/hooks/use-form-response";
import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import { PlexConfigForm, resolver, type formType } from "~/components/forms/plex";
import { getGQLClient } from "~/.server/apollo";
import { gql } from "@apollo/client";
import type { Query } from "~/graphql/gen/graphql";

export function meta() {
    return [
        { title: "Plex - Settings - Music Manager" },
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
            mutation setPlexConfig($config: PlexConfigInput!) {
                setPlexConfig(config: $config)
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
            query getPlexConfig {
                getPlexConfig {
                    protocol
                    host
                    port
                    token
                    libraryID
                }
            }
        `,
    });

    return { errors: error, data: data?.getPlexConfig };
}

export default function Page({ loaderData, actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
        defaultValues: loaderData?.data ?? {},
    });
    useFormResponse<typeof resolver>(actionData);

    return (
        <PlexConfigForm form={form} />
    )
}