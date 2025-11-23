import type { RefCallBack } from "react-hook-form";
import { NumberInput } from "~/components/input/number";

export function ChmodInput({ value, onChange, name, disabled, ref, id }: { value?: number, onChange?: (value: number) => void, name?: string, disabled?: boolean, ref?: RefCallBack, id?: string }) {
    const octal = value?.toString(8).padStart(3, '0');
    const updateOctal = (index: number, newDigit: number) => {
        let digits = octal ? octal.split('').map(d => parseInt(d)) : [0, 0, 0];
        digits[index] = newDigit;
        const newValue = parseInt(digits.map(d => d.toString()).join(''), 8);
        onChange?.(newValue);
    }

    return (
        <div className="flex flex-row items-center gap-4">
            <NumberInput value={octal ? parseInt(octal.charAt(0)) : undefined} onChange={(value) => updateOctal(0, value)} min={0} max={7} />
            <NumberInput value={octal ? parseInt(octal.charAt(1)) : undefined} onChange={(value) => updateOctal(1, value)} min={0} max={7} />
            <NumberInput value={octal ? parseInt(octal.charAt(2)) : undefined} onChange={(value) => updateOctal(2, value)} min={0} max={7} />
        </div>
    )
}