import { useEffect } from "react";
import type { FieldErrors, FieldValues } from "react-hook-form";
import { toast } from "sonner";

export type FormActionResponse<T extends FieldValues> = {
    errors?: FieldErrors<T> | null;
}

export function useFormResponse<T extends FieldValues>(actionData?: FormActionResponse<T> | null) {
    useEffect(() => {
        if (!actionData) return;

        if (actionData.errors) {
            toast.error("There was an error submitting the form. Please check your input.", { richColors: true, description: actionData.errors.name?.message?.toString() });
        } else {
            toast.success("Form submitted successfully.", { richColors: true });
        }
    }, [actionData]);
}