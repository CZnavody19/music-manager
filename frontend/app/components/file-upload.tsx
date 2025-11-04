import { Upload } from "lucide-react";
import { useRef, useState } from "react";
import { Button } from "~/components/ui/button";
import { Card } from "~/components/ui/card";

export function FileUpload({ name, accept, acceptText, multiple }: { name: string, accept: string, acceptText: string, multiple?: boolean }) {
    const filePickerRef = useRef<HTMLInputElement>(null);
    const [text, setText] = useState(acceptText);

    const openFilePicker = () => {
        filePickerRef.current?.click();
    };

    const onDropFiles = (event: React.DragEvent) => {
        event.preventDefault();
        const droppedFiles = event.dataTransfer.files;
        if (filePickerRef.current && (droppedFiles.length == 1 || multiple)) {
            filePickerRef.current.files = droppedFiles;
        }
    };

    const onFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const files = event.target.files;
        if (!files) return;
        if (files.length === 0) {
            setText(acceptText);
            return;
        }

        if (multiple) {
            setText(`${files.length} file(s) selected`);
            return;
        }

        setText(`${files[0].name} selected`);
    };

    return (
        <Card
            className="group flex max-h-[200px] w-full max-w-sm flex-col items-center justify-center gap-4 py-8 border-dashed text-sm cursor-pointer hover:bg-muted/50 transition-colors"
            onDragOver={(e) => e.preventDefault()}
            onDrop={onDropFiles}
            onClick={openFilePicker}
        >
            <div className="grid space-y-3">
                <div className="flex items-center gap-x-2 text-muted-foreground">
                    <Upload className="size-5" />
                    <div>
                        Drop {multiple ? "files" : "a file"} here or{" "}
                        <Button
                            variant="link"
                            className="text-primary p-0 h-auto font-normal"
                            onClick={openFilePicker}
                        >
                            browse files
                        </Button>{" "}
                        to add
                    </div>
                </div>
            </div>
            <input
                ref={filePickerRef}
                type="file"
                className="hidden"
                accept={accept}
                multiple={multiple}
                name={name}
                onChange={onFileChange}
            />
            <span className="text-base/6 text-muted-foreground group-disabled:opacity-50 mt-2 block sm:text-xs">{text}</span>
        </Card>
    )
}