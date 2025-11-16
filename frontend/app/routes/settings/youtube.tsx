import { toast } from "sonner";
import { FileUpload } from "~/components/file-upload";
import { Button } from "~/components/ui/button";
import { Field, FieldDescription, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import type { Route } from "./+types/youtube";
import { getGQLClient } from "~/.server/apollo";
import type { Query } from "~/graphql/gen/graphql";
import { gql } from "@apollo/client";
import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import { resolver, YouTubeConfigForm, type formType } from "~/components/forms/youtube";
import { useFormResponse } from "~/hooks/use-form-response";

export function meta() {
    return [
        { title: "YouTube - Settings - Music Manager" },
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
            mutation setYoutubeConfig($config: YoutubeConfigInput!) {
                setYoutubeConfig(config: $config)
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

    const { data } = await client.query<Query>({
        query: gql`
            query getYoutubeConfig {
                getYoutubeConfig{
                    playlistID
                }
            }
        `,
    })

    return { apiUrl: "http://" + process.env.PUBLIC_API_URL, data: data?.getYoutubeConfig }
}

export default function Page({ loaderData, actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
        defaultValues: loaderData?.data ?? {},
    });
    useFormResponse<typeof resolver>(actionData);

    return (
        <div className="grid grid-cols-2 gap-20">
            <YouTubeConfigForm form={form} />

            <form className="px-6 py-4" action={(formData) => {
                fetch(new URL("/upload/youtube", loaderData.apiUrl), {
                    method: "POST",
                    body: formData,
                }).then((res) => {
                    if (res.ok) {
                        toast.success("Files uploaded successfully", { richColors: true });
                    } else {
                        toast.error("Failed to upload files", { richColors: true });
                    }
                });
            }}>
                <FieldSet>
                    <FieldLegend>OAuth</FieldLegend>
                    <FieldDescription>Configuration for YouTube authentication.</FieldDescription>
                    <FieldGroup>
                        <Field>
                            <FieldLabel htmlFor="oauth">OAuth file</FieldLabel>
                            <FileUpload name="oauth" accept="application/json" acceptText="Supported: JSON files only (max. 10MB)" />
                        </Field>
                        <Field>
                            <FieldLabel htmlFor="token">Token file</FieldLabel>
                            <FileUpload name="token" accept="application/json" acceptText="Supported: JSON files only (max. 10MB)" />
                        </Field>
                        <Field orientation="horizontal">
                            <Button type="submit">Submit</Button>
                        </Field>
                    </FieldGroup>
                </FieldSet>
            </form>
        </div>
    )
}