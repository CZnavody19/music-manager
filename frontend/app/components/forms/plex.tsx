import { Form } from "react-router";
import { Controller } from "react-hook-form";
import { useRemixForm } from "remix-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type z from "zod";
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import { PlexSettingsSchema } from "~/schemas";
import { NumberInput } from "~/components/input/number";
import { TabSelect } from "~/components/input/tab-select";

export const resolver = zodResolver(PlexSettingsSchema);
export type formType = z.infer<typeof PlexSettingsSchema>;

export function PlexConfigForm({ form }: { form: ReturnType<typeof useRemixForm<formType>> }) {
    return (
        <Form className="px-6 py-4" onSubmit={form.handleSubmit} method="POST">
            <FieldSet>
                <FieldLegend>Integration</FieldLegend>
                <FieldDescription>Configuration for Plex integration.</FieldDescription>
                <FieldGroup>
                    <Controller
                        name="protocol"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="protocol">Protocol</FieldLabel>
                                <TabSelect {...field} aria-invalid={fieldState.invalid} id="protocol" options={[
                                    { value: "http", label: "HTTP" },
                                    { value: "https", label: "HTTPS" }
                                ]} />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="host"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="host">Host</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="host" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="port"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="port">Port</FieldLabel>
                                <NumberInput {...field} aria-invalid={fieldState.invalid} id="port" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="token"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="token">Token</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="token" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="libraryID"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="libraryID">Library ID</FieldLabel>
                                <NumberInput {...field} aria-invalid={fieldState.invalid} id="libraryID" />
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