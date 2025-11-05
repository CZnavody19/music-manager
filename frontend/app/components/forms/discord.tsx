import { Form } from "react-router";
import { Controller } from "react-hook-form";
import { useRemixForm } from "remix-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type z from "zod";
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import { DiscordSettingsSchema } from "~/schemas";

export const resolver = zodResolver(DiscordSettingsSchema);
export type formType = z.infer<typeof DiscordSettingsSchema>;

export function DiscordConfigForm({ form }: { form: ReturnType<typeof useRemixForm<formType>> }) {
    return (
        <Form className="px-6 py-4" onSubmit={form.handleSubmit} method="POST">
            <FieldSet>
                <FieldLegend>Webhooks</FieldLegend>
                <FieldDescription>Configuration for Discord webhooks.</FieldDescription>
                <FieldGroup>
                    <Controller
                        name="webhookURL"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="url">Webhook URL</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="url" />
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