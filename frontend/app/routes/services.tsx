import { getGQLClient } from "~/.server/apollo"
import type { Route } from "./+types/services"
import { gql } from "@apollo/client";
import type { Query } from "~/graphql/gen/graphql";
import { Service } from "~/components/service";

type ServiceActionBody = {
    enable: boolean;
    id: string;
}

export function links() {
    return [
        { rel: "preload", href: "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/youtube.svg", as: "image" },
        { rel: "preload", href: "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/tidal-dark.svg", as: "image" },
        { rel: "preload", href: "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/discord.svg", as: "image" },
        { rel: "preload", href: "https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/plex.svg", as: "image" },
    ]
}

export async function action({ request }: Route.ActionArgs) {
    const body: ServiceActionBody = await request.json();
    const { client } = await getGQLClient(request);

    switch (body.id) {
        case "discord":
            await client.mutate({
                mutation: gql`
                    mutation enableDiscord($enable: Boolean!) {
                        enableDiscord(enable: $enable)
                    }
                `,
                variables: { enable: body.enable },
            });
            return {};
        case "plex":
            await client.mutate({
                mutation: gql`
                    mutation enablePlex($enable: Boolean!) {
                        enablePlex(enable: $enable)
                    }
                `,
                variables: { enable: body.enable },
            });
            return {};
        case "youtube":
            await client.mutate({
                mutation: gql`
                    mutation enableYoutube($enable: Boolean!) {
                        enableYoutube(enable: $enable)
                    }
                `,
                variables: { enable: body.enable },
            });
            return {};
        default:
            return {};
    }
}

export async function loader({ request }: Route.LoaderArgs) {
    const { client } = await getGQLClient(request);

    const { data } = await client.query<Query>({
        query: gql`
            query getServiceStatus {
                getServiceStatus {
                    youtube,
                    discord,
                    plex
                }
            }
        `,
    });

    return { status: data?.getServiceStatus }
}

export default function Page({ loaderData }: Route.ComponentProps) {
    return (
        <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 md:p-8 gap-20 xl:gap-4">
            <div className="flex flex-col w-full max-w-sm gap-4 mx-auto">
                <h2 className="font-semibold text-2xl text-center">Sources</h2>
                <Service id="youtube" name="YouTube" imageUrl="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/youtube.svg" enabled={loaderData.status?.youtube ?? false} />
            </div>

            <div className="flex flex-col w-full max-w-sm gap-4 mx-auto">
                <h2 className="font-semibold text-2xl text-center">Downloaders</h2>
                <Service id="tidal" name="Tidal" imageUrl="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/tidal-dark.svg" enabled={false} />
            </div>

            <div className="flex flex-col w-full max-w-sm gap-4 mx-auto">
                <h2 className="font-semibold text-2xl text-center">Notifications</h2>
                <Service id="discord" name="Discord" imageUrl="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/discord.svg" enabled={loaderData.status?.discord ?? false} />
            </div>

            <div className="flex flex-col w-full max-w-sm gap-4 mx-auto">
                <h2 className="font-semibold text-2xl text-center">Integrations</h2>
                <Service id="plex" name="Plex" imageUrl="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/svg/plex.svg" enabled={loaderData.status?.plex ?? false} />
            </div>
        </div>
    )
}