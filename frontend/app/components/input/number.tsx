import { Minus, Plus } from "lucide-react";
import type { RefCallBack } from "react-hook-form";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

export function NumberInput({ value, onChange, name, disabled, ref, id }: { value?: number, onChange?: (value: number) => void, name?: string, disabled?: boolean, ref?: RefCallBack, id?: string }) {
    return (
        <div className="flex flex-row items-center">
            <Input ref={ref} disabled={disabled} name={name} id={id} type="number" value={value} onChange={(e) => !!onChange && onChange(e.target.valueAsNumber)} className="flex flex-row items-center gap-2 rounded-r-none [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none" />
            <Button type="button" size="icon" variant="outline" className="rounded-none" onClick={() => !!value && !!onChange && onChange(value + 1)}><Plus /></Button>
            <Button type="button" size="icon" variant="outline" className="rounded-l-none" onClick={() => !!value && !!onChange && onChange(value - 1)}><Minus /></Button>
        </div>
    )
}