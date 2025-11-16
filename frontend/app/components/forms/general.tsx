import { Form } from "react-router";
import { Controller } from "react-hook-form";
import { useRemixForm } from "remix-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type z from "zod";
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import { GeneralSettingsSchema } from "~/schemas";

export const resolver = zodResolver(GeneralSettingsSchema);
export type formType = z.infer<typeof GeneralSettingsSchema>;

export function GeneralConfigForm({ form }: { form: ReturnType<typeof useRemixForm<formType>> }) {
    return (
        <Form className="px-6 py-4" onSubmit={form.handleSubmit} method="POST">
            <FieldSet>
                <FieldLegend>General</FieldLegend>
                <FieldDescription>Configuration for general settings.</FieldDescription>
                <FieldGroup>
                    <Controller
                        name="downloadPath"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="downloadPath">Download Path</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="downloadPath" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="tempPath"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="tempPath">Temporary Path</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="tempPath" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Field orientation="horizontal">
                        <Button type="submit">Submit</Button>
                    </Field>
                </FieldGroup>
            </FieldSet>
        </Form>
    )
}