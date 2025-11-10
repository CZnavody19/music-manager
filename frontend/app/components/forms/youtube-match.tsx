import { zodResolver } from "@hookform/resolvers/zod";
import type z from "zod";
import { YouTubeMatchSchema } from "~/schemas";
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import { Form } from "react-router";
import { Controller } from "react-hook-form";
import type { useRemixForm } from "remix-hook-form";

export const resolver = zodResolver(YouTubeMatchSchema);
export type formType = z.infer<typeof YouTubeMatchSchema>;

export function YouTubeMatchForm({ form }: { form: ReturnType<typeof useRemixForm<formType>> }) {
    return (
        <Form className="px-6 py-4" onSubmit={form.handleSubmit} method="POST">
            <FieldSet>
                <FieldLegend>Match</FieldLegend>
                <FieldDescription>Match a YouTube video with a recording.<br />Use <i>artist:"" AND recording:""</i> to search.</FieldDescription>
                <FieldGroup>
                    <Controller
                        name="trackId"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="id">Recording ID</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="id" />
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