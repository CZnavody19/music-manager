import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import { AuthForm, resolver, type formType } from "~/components/forms/auth";
import type { Route } from "./+types/login";
import { getGQLClient } from "~/.server/apollo";
import { gql } from "@apollo/client";
import { useFormResponse } from "~/hooks/use-form-response";
import type { Mutation } from "~/graphql/gen/graphql";
import { getSession, saveSession } from "~/.server/session";
import { redirect } from "react-router";

export function meta() {
    return [
        { title: "Login - Music Manager" },
        { name: "description", content: "Manage your music collection with ease." },
    ];
}

export async function action({ request }: Route.ActionArgs) {
    const { errors, data: formData } = await getValidatedFormData(request, resolver);

    if (errors) {
        return { errors };
    }

    const { client } = await getGQLClient(request);

    const { data, error } = await client.mutate<Mutation>({
        mutation: gql`
            mutation login($input: LoginInput!) {
                login(input: $input)
            }
        `,
        variables: {
            input: formData,
        },
    });

    if (!data?.login) {
        return { errors: { form: { message: error?.message || "Login failed" } } };
    }

    const session = await getSession(request);

    session.set("token", data.login);

    return redirect(new URL(request.url).searchParams.get("redirect") ?? "/", await saveSession(session));
}

export async function loader({ request }: Route.LoaderArgs) {
    const session = await getSession(request);

    if (session.has("token")) {
        return redirect(new URL(request.url).searchParams.get("redirect") ?? "/");
    }

    return null;
}

export default function Page({ actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
    });
    useFormResponse<typeof resolver>(actionData);

    return (
        <div className="flex justify-center my-auto">
            <div className="w-sm">
                <AuthForm form={form} login />
            </div>
        </div>
    )
}