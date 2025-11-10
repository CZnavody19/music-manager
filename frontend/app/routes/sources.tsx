import type { Route } from "./+types/sources";
import { getGQLClient } from "~/.server/apollo";
import type { Query } from "~/graphql/gen/graphql";
import { gql } from "@apollo/client";
import { Button } from "~/components/ui/button";
import { Outlet, useSubmit } from "react-router";
import { RefreshCw } from "lucide-react";
import { useFormResponse } from "~/hooks/use-form-response";
import YoutubeTable from "~/components/tables/youtube";

export function meta() {
    return [
        { title: "Sources - Music Manager" },
        { name: "description", content: "Manage your music collection with ease." },
    ];
}

export async function action({ request }: Route.ActionArgs) {
    const body = await request.json();

    if (!body.action) {
        return {}
    }

    const { client } = await getGQLClient(request);

    const { error } = await client.mutate({
        mutation: gql`
            mutation refreshPlaylist {
                refreshPlaylist
            }
        `,
    });

    return { errors: error };
}

export async function loader({ request }: Route.LoaderArgs) {
    const { client } = await getGQLClient(request);

    const { data } = await client.query<Query>({
        query: gql`
            query getVideosInPlaylist {
                getVideosInPlaylist {
                    id
                    title
                    channelTitle
                    thumbnailUrl
                    position
                    duration
                    linked
                }
            }
        `,
    });

    return { videos: data?.getVideosInPlaylist ?? [] };
}

export default function Page({ loaderData, actionData }: Route.ComponentProps) {
    const submit = useSubmit();
    useFormResponse(actionData as any);

    return (
        <div>
            <div className="flex flex-col w-full p-4 gap-4">
                <div className="flex flex-row items-center justify-between">
                    <h2 className="font-semibold text-2xl">YouTube</h2>
                    <div className="flex flex-row items-center gap-2">
                        <Button variant="outline" size="icon" onClick={() => submit({ action: "refresh" }, { method: "POST", encType: "application/json" })}><RefreshCw /></Button>
                    </div>
                </div>
                <YoutubeTable tracks={loaderData.videos} />
            </div>
            <Outlet />
        </div>
    )
}