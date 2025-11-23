import { Minus, Plus } from "lucide-react";
import type { RefCallBack } from "react-hook-form";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

export function NumberInput({ value, onChange, name, disabled, ref, id, min, max }: { value?: number, onChange?: (value: number) => void, name?: string, disabled?: boolean, ref?: RefCallBack, id?: string, min?: number, max?: number }) {
    return (
        <div className="flex flex-row items-center w-full">
            <Input ref={ref} disabled={disabled} name={name} id={id} type="number" min={min} max={max} value={value} onChange={(e) => !!onChange && (min !== undefined ? e.target.valueAsNumber >= min : true) && (max !== undefined ? e.target.valueAsNumber <= max : true) && onChange(e.target.valueAsNumber)} className="flex flex-row items-center gap-2 rounded-r-none [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none" />
            <Button type="button" size="icon" variant="outline" className="rounded-none" onClick={() => value !== undefined && !!onChange && (max !== undefined ? value < max : true) && onChange(value + 1)}><Plus /></Button>
            <Button type="button" size="icon" variant="outline" className="rounded-l-none" onClick={() => value !== undefined && !!onChange && (min !== undefined ? value > min : true) && onChange(value - 1)}><Minus /></Button>
        </div>
    )
}