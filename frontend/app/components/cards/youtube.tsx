import { ExternalLink } from "lucide-react";
import { Link } from "react-router";
import { Button } from "~/components/ui/button";
import type { YouTubeVideo } from "~/graphql/gen/graphql";

export function YouTubeCard({ video }: { video: YouTubeVideo }) {
    return (
        <div className="flex flex-row border rounded-md overflow-hidden max-w-lg">
            <img src={video.thumbnailUrl} alt={video.title} className="w-[120px] h-[90px] object-cover" />
            <div className="flex flex-col p-2">
                <h3 className="font-semibold line-clamp-1">{video.title}</h3>
                <p className="text-sm text-muted-foreground">{video.channelTitle}</p>
                <p className="text-sm text-muted-foreground">Duration: {video.duration} seconds</p>
            </div>
            <div className="ml-auto p-2 flex flex-row items-center gap-2">
                <Button variant="outline" size="icon" asChild>
                    <Link to={`https://www.youtube.com/watch?v=${video.id}`} target="_blank" rel="noopener noreferrer">
                        <ExternalLink />
                    </Link>
                </Button>
            </div>
        </div>
    )
}