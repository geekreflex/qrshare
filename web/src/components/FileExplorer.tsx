import { useEffect, useState } from "react";

type FileItem = {
  name: string;
  isDir: boolean;
  path: string;
};

export default function FileExplorer() {
  const [files, setFiles] = useState<FileItem[]>([]);

  useEffect(() => {
    fetch("/api/files/list") // or just '/' depending on your Go backend
      .then((res) => res.json())
      .then((data) => setFiles(data))
      .catch((err) => console.error("Error loading files:", err));
  }, []);

  return (
    <div className="grid gap-3">
      {files.map((file) => (
        <div
          key={file.path}
          className="p-3 border rounded bg-white shadow flex justify-between items-center"
        >
          <span>{file.name}</span>
          <a
            href={`/api/files/raw/${file.path}`}
            className="text-blue-500 hover:underline"
            download={!file.isDir}
          >
            {file.isDir ? "Open" : "Download"}
          </a>
        </div>
      ))}
    </div>
  );
}
