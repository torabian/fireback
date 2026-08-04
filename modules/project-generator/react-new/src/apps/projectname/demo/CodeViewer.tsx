import { useState } from "react";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import {
  oneDark,
  duotoneLight,
} from "react-syntax-highlighter/dist/esm/styles/prism";

export const CodeViewer = ({ codeString }: { codeString: string }) => {
  const isDark = document?.body?.classList.contains("dark-theme");
  const [copied, setCopied] = useState(false);

  const handleCopy = () => {
    navigator.clipboard.writeText(codeString).then(() => {
      setCopied(true);
      setTimeout(() => setCopied(false), 1500);
    });
  };

  return (
    <div className="code-viewer-container">
      <button className="copy-button" onClick={handleCopy}>
        {copied ? "Copied!" : "Copy"}
      </button>
      <SyntaxHighlighter language="tsx" style={isDark ? oneDark : duotoneLight}>
        {codeString}
      </SyntaxHighlighter>
    </div>
  );
};
