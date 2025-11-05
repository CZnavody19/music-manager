import { useSubmit } from "react-router";
import { Badge } from "~/components/ui/badge";
import { Button } from "~/components/ui/button";

export function Service({ name, imageUrl, enabled, id }: { name: string, imageUrl: string, enabled: boolean, id: string }) {
    const submit = useSubmit();

    const makeReq = async (enable: boolean) => {
        submit({ enable, id }, { method: "POST", encType: "application/json" });
    }

    return (
        <div className="flex flex-row gap-4 items-center justify-between w-full max-w-sm h-min border rounded-md px-4 py-3">
            <div className="flex flex-row gap-4 items-center">
                <img src={imageUrl} height={52} width={52} />
                <div className="flex flex-col justify-between">
                    <h2 className="text-lg font-semibold">{name}</h2>
                    <Badge variant={enabled ? "success" : "destructive"}>{enabled ? "Enabled" : "Disabled"}</Badge>
                </div>
            </div>
            {enabled ? (
                <Button variant="outline" onClick={() => makeReq(false)}>Disable</Button>
            ) : (
                <Button variant="outline" onClick={() => makeReq(true)}>Enable</Button>
            )}
        </div>
    )
}