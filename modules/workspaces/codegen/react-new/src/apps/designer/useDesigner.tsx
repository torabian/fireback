import { useEffect, useRef, useState } from "react";
import { Module2 } from "./defs";

import { mockYaml } from "./mock";
const yaml = require("js-yaml");

const vscode = (window as any).acquireVsCodeApi
  ? (window as any).acquireVsCodeApi()
  : null;

export function useDesigner() {
  const [content, setContent] = useState<Module2>({});

  const parseYaml = (yamlString: string) => {
    try {
      const parsed = yaml.load(yamlString);
      setContent(parsed);
    } catch (e) {
      console.error("Error parsing YAML:", e);
    }
  };

  const lockReflex = useRef(true);

  useEffect(() => {
    // Use workspaces module as a predefined data
    if (!vscode) {
      parseYaml(mockYaml);
    }

    window.addEventListener("message", (event) => {
      if (event.data.command === "updateContent") {
        lockReflex.current = true;
        parseYaml(event.data.content);

        setTimeout(() => {
          lockReflex.current = false;
        }, 50);
      }
    });
  }, []);

  useEffect(() => {
    if (lockReflex.current) {
      return;
    }
    sendUpdate(yaml.dump(content));
  }, [content]);

  return { content, setContent };
}

function sendUpdate(content: any) {
  if (typeof vscode == "undefined" || !vscode?.postMessage) {
    return;
  }
  vscode.postMessage({ command: "updateFile", content });
}
