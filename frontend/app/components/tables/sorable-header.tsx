import type { Column } from "@tanstack/react-table"
import { ArrowUpDown } from "lucide-react"
import { Button } from "~/components/ui/button"

export function SortableHeader({ column, title, className }: { column: Column<any, unknown>, title: string, className?: string }) {
    return (
        <Button variant="ghost" className={className} onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}>
            {title}
            <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
    )
}