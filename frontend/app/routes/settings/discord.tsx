import type { Route } from "./+types/discord";
import { useFormResponse } from "~/hooks/use-form-response";
import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import { DiscordConfigForm, resolver, type formType } from "~/components/forms/discord";
import { getGQLClient } from "~/.server/apollo";
import { gql } from "@apollo/client";
import type { Query } from "~/graphql/gen/graphql";

export function meta() {
    return [
        { title: "Discord - Settings - Music Manager" },
        { name: "description", content: "Manage your music collection with ease." },
    ];
}


export async function action({ request }: Route.ActionArgs) {
    const { client } = await getGQLClient(request);

    switch (request.method) {
        case "POST":
            const { errors, data } = await getValidatedFormData(request, resolver);

            if (errors) {
                return { errors };
            }

            const { error } = await client.mutate({
                mutation: gql`
                    mutation setDiscordConfig($config: DiscordConfigInput!) {
                        setDiscordConfig(config: $config)
                    }
                `,
                variables: {
                    config: data,
                },
            });

            return { errors: error };

        case "PUT":
            const body = await request.json();
            if (body.action !== "test") {
                return { errors: [{ message: "Invalid action" }] };
            }

            const { error: err } = await client.mutate({
                mutation: gql`
                    mutation sendTestDiscordMessage {
                        sendTestDiscordMessage
                    }
                `,
            });

            return { errors: err, message: "Test message sent successfully." };

        default:
            return { errors: [{ message: "Invalid request method" }] };
    }
}

export async function loader({ request }: Route.LoaderArgs) {
    const { client } = await getGQLClient(request);

    const { data, error } = await client.query<Query>({
        query: gql`
            query getDiscordConfig {
                getDiscordConfig {
                    webhookURL
                }
            }
        `,
    });

    return { errors: error, data: data?.getDiscordConfig };
}

export default function Page({ loaderData, actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
        defaultValues: loaderData?.data ?? {},
    });
    useFormResponse<typeof resolver>(actionData);

    return (
        <DiscordConfigForm form={form} />
    )
}