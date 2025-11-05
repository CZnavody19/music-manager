import type { RefCallBack } from "react-hook-form"
import { Tabs, TabsList, TabsTrigger } from "~/components/ui/tabs"

export type TabSelectOption = {
    value: string,
    label: string
};

export function TabSelect({ value, onChange, options, disabled, ref, id }: { value?: string, onChange?: (value: string) => void, options: TabSelectOption[], disabled?: boolean, ref?: RefCallBack, id?: string }) {
    return (
        <Tabs value={value} onValueChange={onChange} id={id} ref={ref} className="w-[400px]">
            <TabsList>
                {options.map((option) => (
                    <TabsTrigger disabled={disabled} value={option.value} key={option.value}>{option.label}</TabsTrigger>
                ))}
            </TabsList>
        </Tabs>
    )
}