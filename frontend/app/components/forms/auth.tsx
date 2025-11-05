import { Form } from "react-router";
import { Controller } from "react-hook-form";
import { useRemixForm } from "remix-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type z from "zod";
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import { AuthSchema } from "~/schemas";

export const resolver = zodResolver(AuthSchema);
export type formType = z.infer<typeof AuthSchema>;

export function AuthForm({ form, login }: { form: ReturnType<typeof useRemixForm<formType>>, login?: boolean }) {
    return (
        <Form className="px-6 py-4" onSubmit={form.handleSubmit} method="POST">
            <FieldSet>
                <FieldLegend>{login ? "Login" : "Authentication"}</FieldLegend>
                <FieldDescription>{login ? "Login to the application." : "Configuration for authentication."}</FieldDescription>
                <FieldGroup>
                    <Controller
                        name="username"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="username">Username</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="username" />
                                {fieldState.invalid && (
                                    <FieldError errors={[fieldState.error]} />
                                )}
                            </Field>
                        )} />
                    <Controller
                        name="password"
                        control={form.control}
                        render={({ field, fieldState }) => (
                            <Field data-invalid={fieldState.invalid}>
                                <FieldLabel htmlFor="password">Password</FieldLabel>
                                <Input {...field} aria-invalid={fieldState.invalid} id="password" type="password" />
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