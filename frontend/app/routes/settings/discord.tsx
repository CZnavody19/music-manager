import type { Route } from "./+types/discord";
import { useFormResponse } from "~/hooks/use-form-response";
import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import { DiscordConfigForm, resolver, type formType } from "~/components/forms/discord";
import { getGQLClient } from "~/.server/apollo";
import { gql } from "@apollo/client";

export function meta() {
    return [
        { title: "Discord - Settings - Music Manager" },
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
            mutation setDiscordConfig($config: DiscordConfig!) {
                setDiscordConfig(config: $config)
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
}

export default function Page({ actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
    });
    useFormResponse<typeof resolver>(actionData);

    return (
        <DiscordConfigForm form={form} />
    )
}