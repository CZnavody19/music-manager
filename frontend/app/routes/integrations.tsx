import { gql } from "@apollo/client";
import { RefreshCw } from "lucide-react";
import { useSubmit } from "react-router";
import { getGQLClient } from "~/.server/apollo";
import PlexTable from "~/components/tables/plex";
import { Button } from "~/components/ui/button";
import type { Query } from "~/graphql/gen/graphql";
import type { Route } from "./+types/integrations";

export function meta() {
    return [
        { title: "Integrations - Music Manager" },
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
            mutation refreshPlexTracks {
                refreshPlexTracks
            }
        `,
    });

    return { errors: error };
}

export async function loader({ request }: Route.LoaderArgs) {
    const { client } = await getGQLClient(request);

    const { data } = await client.query<Query>({
        query: gql`
            query {
                getPlexTracks {
                    id
                    title
                    artist
                    duration
                    mbid
                    trackID
                }
            }
        `,
    });

    return { tracks: data?.getPlexTracks ?? [] };
}

export default function Page({ loaderData }: Route.ComponentProps) {
    const submit = useSubmit();

    return (
        <div>
            <div className="flex flex-col w-full p-4 gap-4">
                <div className="flex flex-row items-center justify-between">
                    <h2 className="font-semibold text-2xl">Plex</h2>
                    <div className="flex flex-row items-center gap-2">
                        <Button variant="outline" size="icon" onClick={() => submit({ action: "refresh" }, { method: "POST", encType: "application/json" })}><RefreshCw /></Button>
                    </div>
                </div>
                <PlexTable tracks={loaderData.tracks} />
            </div>
        </div>
    )
}