import type { ColumnDef } from "@tanstack/react-table"
import { Trash2 } from "lucide-react"
import { useSubmit } from "react-router"
import { BaseTable, type Filter } from "~/components/tables/base"
import { SortableHeader } from "~/components/tables/sorable-header"
import { Badge } from "~/components/ui/badge"
import { Button } from "~/components/ui/button"
import type { Track } from "~/graphql/gen/graphql"
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

export default function TrackTable({ tracks }: { tracks: Track[] }) {
    const submit = useSubmit();

    const columns: ColumnDef<Track>[] = [
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
            accessorKey: "length",
            header: ({ column }) => <SortableHeader className="flex mx-auto" title="Length" column={column} />,
            cell: ({ row }) => {
                return <div className="text-center">{formatSeconds(row.original.length / 1000)}</div>
            },
        },
        {
            accessorKey: "id",
            header: ({ column }) => <SortableHeader title="MBID" column={column} />,
        },
        {
            accessorKey: "isrcs",
            header: () => <div className="text-center">ISRCs</div>,
            cell: ({ row }) => {
                return <Badge variant={row.original.isrcs.length > 0 ? "success" : "destructive"} className="flex mx-auto">{row.original.isrcs.length}</Badge>
            }
        },
        {
            accessorKey: "linkedYoutube",
            header: ({ column }) => <SortableHeader title="YouTube" column={column} />,
            cell: ({ row }) => {
                if (row.original.linkedYoutube) {
                    return <Badge variant="success">Linked</Badge>
                } else {
                    return <Badge variant="destructive">Not linked</Badge>
                }
            },
        },
        {
            accessorKey: "linkedPlex",
            header: ({ column }) => <SortableHeader title="Plex" column={column} />,
            cell: ({ row }) => {
                if (row.original.linkedPlex) {
                    return <Badge variant="success">Linked</Badge>
                } else {
                    return <Badge variant="destructive">Not linked</Badge>
                }
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