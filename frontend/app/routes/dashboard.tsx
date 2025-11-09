import { getGQLClient } from "~/.server/apollo";
import type { Route } from "./+types/dashboard";
import { gql } from "@apollo/client";
import type { Query } from "~/graphql/gen/graphql";
import TrackTable from "~/components/tables/track";

export function meta() {
	return [
		{ title: "Dashboard - Music Manager" },
		{ name: "description", content: "Manage your music collection with ease." },
	];
}

export async function loader({ request }: Route.LoaderArgs) {
	const { client } = await getGQLClient(request);

	const { data } = await client.query<Query>({
		query: gql`
			query getTracks {
				getTracks {
					id
					title
					artist
					length
					isrcs
					linkedYoutube
					linkedPlex
				}
			}
		`,
	});

	return { tracks: data?.getTracks ?? [] };
}

export default function Page({ loaderData }: Route.ComponentProps) {
	return (
		<div>
			<div className="flex flex-col w-full p-4 gap-4">
				<div className="flex flex-row items-center justify-between">
					<h2 className="font-semibold text-2xl">Tracks</h2>
					<div className="flex flex-row items-center gap-2">

					</div>
				</div>
				<TrackTable tracks={loaderData.tracks} />
				{/* <ScrollArea className="h-[calc(100vh-8.75rem)] w-full">
					<div className="flex flex-col gap-4">
						{loaderData.tracks?.map((track) => (
							<TrackCard key={track.id} track={track} />
						))}
					</div>
				</ScrollArea> */}
			</div>
		</div>
	)
}
