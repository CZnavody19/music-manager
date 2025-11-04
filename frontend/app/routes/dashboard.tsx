import { getGQLClient } from "~/.server/apollo";
import type { Route } from "./+types/dashboard";
import { Card, CardAction, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "~/components/ui/card"
import { gql } from "@apollo/client";
import type { Query } from "~/graphql/gen/graphql";

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
			query getStatus {
				getStatus {
					youtubeActive
				}
			}
		`,
	});

	return data?.getStatus;
}

export default function Page({ loaderData }: Route.ComponentProps) {
	return (
		<div>
			<Card className="w-full max-w-sm">
				<CardHeader>
					<CardTitle>Card Title</CardTitle>
					<CardDescription>Card Description</CardDescription>
					<CardAction>Card Action</CardAction>
				</CardHeader>
				<CardContent>
					<p>{loaderData?.youtubeActive ? "YouTube is active" : "YouTube is inactive"}</p>
				</CardContent>
				<CardFooter>
					<p>Card Footer</p>
				</CardFooter>
			</Card>
		</div>
	)
}
