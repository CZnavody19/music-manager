import { toast } from "sonner";
import { FileUpload } from "~/components/file-upload";
import { Button } from "~/components/ui/button";
import { Field, FieldDescription, FieldGroup, FieldLabel, FieldLegend, FieldSet } from "~/components/ui/field"
import type { Route } from "./+types/youtube";

export function meta() {
    return [
        { title: "YouTube - Settings - Music Manager" },
        { name: "description", content: "Manage your music collection with ease." },
    ];
}

export async function loader() {
    return { apiUrl: "http://localhost:8080/" }
}

export default function Page({ loaderData }: Route.ComponentProps) {
    return (
        <form className="px-6 py-4" action={(formData) => {
            fetch(new URL("/upload/youtube", loaderData.apiUrl), {
                method: "POST",
                body: formData,
            }).then((res) => {
                if (res.ok) {
                    toast.success("Files uploaded successfully", { richColors: true });
                } else {
                    toast.error("Failed to upload files", { richColors: true });
                }
            });
        }}>
            <FieldSet>
                <FieldLegend>OAuth</FieldLegend>
                <FieldDescription>Configuration for YouTube authentication.</FieldDescription>
                <FieldGroup>
                    <Field>
                        <FieldLabel htmlFor="oauth">OAuth file</FieldLabel>
                        <FileUpload name="oauth" accept="application/json" acceptText="Supported: JSON files only (max. 10MB)" />
                    </Field>
                    <Field>
                        <FieldLabel htmlFor="token">Token file</FieldLabel>
                        <FileUpload name="token" accept="application/json" acceptText="Supported: JSON files only (max. 10MB)" />
                    </Field>
                    <Field orientation="horizontal">
                        <Button type="submit">Submit</Button>
                    </Field>
                </FieldGroup>
            </FieldSet>
        </form>
    )
}