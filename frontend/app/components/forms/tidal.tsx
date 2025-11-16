import { Form } from "react-router";
import { Controller } from "react-hook-form";
import { useRemixForm } from "remix-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type z from "zod";
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import { TidalSettingsSchema } from "~/schemas";
import { NumberInput } from "~/components/input/number";
import { TabSelect } from "~/components/input/tab-select";

export const resolver = zodResolver(TidalSettingsSchema);
export type formType = z.infer<typeof TidalSettingsSchema>;

export function TidalConfigForm({ form }: { form: ReturnType<typeof useRemixForm<formType>> }) {
    return (
        <Form className="px-6 py-4" onSubmit={form.handleSubmit} method="POST">
            <FieldSet>
                <FieldLegend>Tidal</FieldLegend>
                <FieldDescription>Configuration for Tidal integration.</FieldDescription>
                <FieldGroup>
                    <Controller
                        name="authTokenType"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="authTokenType">Token type</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="authTokenType" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="authAccessToken"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="authAccessToken">Access token</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="authAccessToken" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="authRefreshToken"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="authRefreshToken">Refresh token</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="authRefreshToken" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="authExpiresAt"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="authExpiresAt">Expires at</FieldLabel>
                                <NumberInput {...field} value={new Date(field.value).getTime() / 1000} onChange={(value) => field.onChange(new Date(value * 1000))} aria-invalid={fieldState.invalid} id="authExpiresAt" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="authClientID"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="authClientID">Client ID</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="authClientID" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="authClientSecret"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="authClientSecret">Client Secret</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="authClientSecret" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="downloadTimeout"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="downloadTimeout">Download Timeout</FieldLabel>
                                <NumberInput {...field} aria-invalid={fieldState.invalid} id="downloadTimeout" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="downloadRetries"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="downloadRetries">Download Retries</FieldLabel>
                                <NumberInput {...field} aria-invalid={fieldState.invalid} id="downloadRetries" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="downloadThreads"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="downloadThreads">Download Threads</FieldLabel>
                                <NumberInput {...field} aria-invalid={fieldState.invalid} id="downloadThreads" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="audioQuality"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="audioQuality">Audio Quality</FieldLabel>
                                <TabSelect {...field} aria-invalid={fieldState.invalid} id="audioQuality" options={[
                                    { value: "LOW", label: "AAC 96kbps" },
                                    { value: "HIGH", label: "AAC 320kbps" },
                                    { value: "LOSSLESS", label: "FLAC" },
                                    { value: "HI_RES_LOSSLESS", label: "HiRes FLAC" }
                                ]} />
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