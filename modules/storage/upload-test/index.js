#!/usr/bin/env node

import fs from "node:fs";
import path from "node:path";
import tus from "tus-js-client";

const filePath = process.argv[2];
const endpoint =
  process.env.TUS_ENDPOINT || "http://localhost:4500/storage/files";

if (!filePath) {
  console.error("Usage: node index.js <file>");
  console.error(
    "       TUS_ENDPOINT=http://host:port/storage/files node index.js <file>",
  );
  process.exit(1);
}

const stat = fs.statSync(filePath);

// Persists in-progress upload URLs to disk (keyed by file fingerprint), so
// that if this process is killed mid-upload (Ctrl+C, crash, network drop)
// the next run resumes from where it left off instead of starting over.
const urlStorage = new tus.FileUrlStorage(
  path.join(
    path.dirname(new URL(import.meta.url).pathname),
    ".tus-uploads.json",
  ),
);

const upload = new tus.Upload(fs.createReadStream(filePath), {
  endpoint,
  uploadSize: stat.size, // required since a stream can't report its own size
  retryDelays: [0, 1000, 3000, 5000],
  urlStorage,
  metadata: {
    filename: path.basename(filePath),
    filetype: "application/octet-stream",
  },

  onError(error) {
    console.error("\nUpload failed:", error.message ?? error);
    process.exit(1);
  },

  onProgress(uploaded, total) {
    const pct = ((uploaded / total) * 100).toFixed(1);
    process.stdout.write(`\r${pct}% (${uploaded}/${total} bytes)   `);
  },

  onSuccess() {
    console.log("\nUpload complete!");
    console.log("URL:", upload.url);
  },
});

// Ctrl+C aborts the in-flight request cleanly; the upload's progress is
// already persisted server-side (and its URL cached in .tus-uploads.json),
// so simply re-running this script resumes instead of re-uploading.
process.on("SIGINT", async () => {
  console.log(
    "\nInterrupted - aborting current request (progress is saved, re-run to resume)...",
  );
  await upload.abort();
  process.exit(130);
});

const previousUploads = await upload.findPreviousUploads();
if (previousUploads.length > 0) {
  console.log(
    `Found a previous upload for this file, resuming: ${previousUploads[0].uploadUrl}`,
  );
  upload.resumeFromPreviousUpload(previousUploads[0]);
} else {
  console.log(
    `Starting new upload of ${filePath} (${stat.size} bytes) to ${endpoint}`,
  );
}

upload.start();
