import { useNavigate } from "react-router";
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "~/components/ui/dialog";
import type { Route } from "./+types/youtube";
import { resolver, YouTubeMatchForm, type formType } from "~/components/forms/youtube-match";
import { getValidatedFormData, useRemixForm } from "remix-hook-form";
import { getGQLClient } from "~/.server/apollo";
import type { Query } from "~/graphql/gen/graphql";
import { gql } from "@apollo/client";
import { formatSeconds } from "~/lib/utils";
import { useFormResponse } from "~/hooks/use-form-response";

export function meta() {
    return [
        { title: "Match video - Music Manager" },
        { name: "description", content: "Manage your music collection with ease." },
    ];
}

export async function action({ request, params }: Route.ActionArgs) {
    const { errors, data } = await getValidatedFormData(request, resolver);

    if (errors) {
        return { errors };
    }

    const { client } = await getGQLClient(request);

    const { error } = await client.mutate({
        mutation: gql`
            mutation matchVideo($videoID: String!, $trackID: UUID!) {
                matchVideo(videoID: $videoID, trackID: $trackID)
            }
        `,
        variables: {
            videoID: params.videoId,
            trackID: data.trackId,
        },
    });

    return { errors: error };
}

export async function loader({ request, params }: Route.LoaderArgs) {
    const { client } = await getGQLClient(request);

    const { data } = await client.query<Query>({
        query: gql`
            query getVideoByID($videoID: String!) {
                getVideoByID(videoID: $videoID) {
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
        variables: {
            videoID: params.videoId,
        }
    });

    return { video: data?.getVideoByID ?? null };
}

export default function Page({ params, loaderData, actionData }: Route.ComponentProps) {
    const form = useRemixForm<formType>({
        resolver,
    });
    const navigate = useNavigate();
    useFormResponse<typeof resolver>(actionData);

    return (
        <Dialog open={true} onOpenChange={() => navigate(-1)}>
            <DialogContent className="sm:max-w-full w-min">
                <DialogHeader>
                    <DialogTitle>Match video</DialogTitle>
                </DialogHeader>
                <div className="flex flex-row gap-4 h-[75vh] w-[75vw]">
                    <div className="w-2/5 h-full">
                        {loaderData.video && (
                            <div className="mb-4">
                                <h2 className="text-xl font-bold">{loaderData.video?.title}</h2>
                                <p className="text-sm text-muted-foreground">{loaderData.video?.channelTitle}</p>
                                <p className="text-sm text-muted-foreground">{formatSeconds(loaderData.video.duration)}</p>
                            </div>
                        )}
                        <iframe id="ytplayer" src={`https://www.youtube.com/embed/${params.videoId}?fs=0&rel=0`} className="w-full h-1/2" />
                        <YouTubeMatchForm form={form} />
                    </div>
                    <iframe id="musicbrainz" src="https://musicbrainz.org/search?type=recording&method=advanced" className="w-3/5" />
                </div>
            </DialogContent>
        </Dialog>
    )
}