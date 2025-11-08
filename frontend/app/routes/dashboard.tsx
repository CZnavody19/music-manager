import { getGQLClient } from "~/.server/apollo";
import type { Route } from "./+types/dashboard";
import { gql } from "@apollo/client";
import type { Query } from "~/graphql/gen/graphql";
import { ScrollArea } from "~/components/ui/scroll-area";
import { YouTubeCard } from "~/components/cards/youtube";

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
			query getVideosInPlaylist {
				getVideosInPlaylist {
					id
					title
					channelTitle
					thumbnailUrl
					duration
					position
				}
			}
		`,
	});

	const videos = data?.getVideosInPlaylist;

	videos?.sort((a, b) => (a.position ?? 0) - (b.position ?? 0));

	return { videos };
}

export default function Page({ loaderData }: Route.ComponentProps) {
	return (
		<div>
			<div className="flex flex-col w-full max-w-lg p-4 gap-4">
				<div className="flex flex-row items-center justify-between">
					<h2 className="font-semibold text-2xl">Playlist items</h2>
					<div className="flex flex-row items-center gap-2">

					</div>
				</div>
				<ScrollArea className="h-[calc(100vh-8.75rem)] w-full">
					<div className="flex flex-col gap-4">
						{loaderData.videos?.map((video) => (
							<YouTubeCard key={video.id} video={video} />
						))}
					</div>
				</ScrollArea>
			</div>
		</div>
	)
}
