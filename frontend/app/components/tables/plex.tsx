import type { ColumnDef } from "@tanstack/react-table"
import { Trash2 } from "lucide-react"
import { useSubmit } from "react-router"
import { BaseTable, type Filter } from "~/components/tables/base"
import { SortableHeader } from "~/components/tables/sorable-header"
import { Badge } from "~/components/ui/badge"
import { Button } from "~/components/ui/button"
import type { PlexTrack } from "~/graphql/gen/graphql"
import { formatSeconds } from "~/lib/utils"

const filters: Filter[] = [
    {
        column: "title",
        label: "Filter by title",
    },
    {
        column: "artist",
        label: "Filter by artist",
    },
];

export default function PlexTable({ tracks }: { tracks: PlexTrack[] }) {
    const submit = useSubmit()

    const columns: ColumnDef<PlexTrack>[] = [
        {
            accessorKey: "title",
            header: ({ column }) => <SortableHeader title="Title" column={column} />,
            cell: ({ row }) => {
                return row.original.title.slice(0, 35) + (row.original.title.length > 35 ? "…" : "")
            }
        },
        {
            accessorKey: "artist",
            header: ({ column }) => <SortableHeader title="Artist" column={column} />,
            cell: ({ row }) => {
                return row.original.artist.slice(0, 50) + (row.original.artist.length > 50 ? "…" : "")
            }
        },
        {
            accessorKey: "duration",
            header: ({ column }) => <SortableHeader className="flex mx-auto" title="Duration" column={column} />,
            cell: ({ row }) => {
                return <div className="text-center">{formatSeconds(row.original.duration / 1000)}</div>
            },
        },
        {
            accessorKey: "mbid",
            header: () => <div className="text-center">Plex matched</div>,
            cell: ({ row }) => {
                return <Badge variant={row.original.mbid ? "success" : "destructive"} className="flex mx-auto">{row.original.mbid ? "Yes" : "No"}</Badge>
            }
        },
        {
            accessorKey: "trackID",
            header: () => <div className="text-center">Matched</div>,
            cell: ({ row }) => {
                return <Badge variant={row.original.trackID ? "success" : "destructive"} className="flex mx-auto">{row.original.trackID ? "Yes" : "No"}</Badge>
            },
        },
        {
            accessorKey: "action",
            header: () => <div className="text-center">Action</div>,
            cell: ({ row }) => {
                return (
                    <div className="flex flex-row items-center justify-center gap-2">
                        <Button variant="destructive" size="icon" onClick={() => submit({ action: "delete", id: row.original.id }, { method: "POST", encType: "application/json" })}><Trash2 /></Button>
                    </div>
                )
            }
        },
    ]

    return <BaseTable columns={columns} filters={filters} data={tracks} />
}