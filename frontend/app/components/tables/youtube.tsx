import type { ColumnDef } from "@tanstack/react-table"
import { BaseTable, type Filter } from "~/components/tables/base"
import { SortableHeader } from "~/components/tables/sorable-header"
import { Badge } from "~/components/ui/badge"
import type { YouTubeVideo } from "~/graphql/gen/graphql"
import { formatSeconds } from "~/lib/utils"
import { Button } from "~/components/ui/button";
import { ExternalLink, Plus } from "lucide-react";
import { Link } from "react-router";

const columns: ColumnDef<YouTubeVideo>[] = [
    {
        accessorKey: "position",
        header: ({ column }) => <SortableHeader title="#" column={column} />,
    },
    {
        accessorKey: "thumbnailUrl",
        header: () => <div>Thumbnail</div>,
        cell: ({ row }) => {
            return <img src={row.original.thumbnailUrl} alt={row.original.title} className="w-[120px] h-[90px] object-cover" />
        }
    },
    {
        accessorKey: "title",
        header: ({ column }) => <SortableHeader title="Title" column={column} />,
        cell: ({ row }) => {
            return row.original.title.slice(0, 50) + (row.original.title.length > 50 ? "…" : "")
        }
    },
    {
        accessorKey: "channelTitle",
        header: ({ column }) => <SortableHeader title="Channel" column={column} />,
        cell: ({ row }) => {
            return row.original.channelTitle.slice(0, 35) + (row.original.channelTitle.length > 35 ? "…" : "")
        }
    },
    {
        accessorKey: "duration",
        header: ({ column }) => <SortableHeader className="flex mx-auto" title="Duration" column={column} />,
        cell: ({ row }) => {
            return <div className="text-center">{formatSeconds(row.original.duration)}</div>
        },
    },
    {
        accessorKey: "linked",
        header: () => <div className="text-center">Matched</div>,
        cell: ({ row }) => {
            return <Badge variant={row.original.linked ? "success" : "destructive"} className="flex mx-auto">{row.original.linked ? "Yes" : "No"}</Badge>
        }
    },
    {
        accessorKey: "action",
        header: () => <div className="text-center">Action</div>,
        cell: ({ row }) => {
            return (
                <div className="flex flex-row items-center justify-center gap-2">
                    {!row.original.linked && <Button variant="outline" size="icon" asChild>
                        <Link to={`match/${row.original.id}`} prefetch="intent">
                            <Plus />
                        </Link>
                    </Button>}
                    <Button variant="outline" size="icon" asChild>
                        <Link to={`https://www.youtube.com/watch?v=${row.original.id}`} target="_blank" rel="noopener noreferrer">
                            <ExternalLink />
                        </Link>
                    </Button>
                </div>
            )
        }
    },
]

const filters: Filter[] = [
    {
        column: "title",
        label: "Filter by title",
    },
    {
        column: "channelTitle",
        label: "Filter by channel",
    },
];

export default function YoutubeTable({ tracks }: { tracks: YouTubeVideo[] }) {
    return <BaseTable columns={columns} filters={filters} data={tracks} />
}