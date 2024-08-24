import { readFileSync } from "fs";
import * as path from "path";
import * as vscode from "vscode";

export class ReactPanel {
  /**
   * Track the currently panel. Only allow a single panel to exist at a time.
   */
  public static currentPanel: ReactPanel | undefined;
  public static activeEditor: vscode.TextEditor | undefined;

  private static readonly viewType = "react";

  private readonly _panel: vscode.WebviewPanel;
  private readonly _extensionPath: string;
  private _disposables: vscode.Disposable[] = [];

  public static createOrShow(extensionPath: string) {
    // const column = vscode.window.activeTextEditor ? vscode.window.activeTextEditor.viewColumn : undefined;
    const column = vscode.ViewColumn.Two;
    // If we already have a panel, show it.
    // Otherwise, create a new panel.
    if (ReactPanel.currentPanel) {
      ReactPanel.currentPanel._panel.reveal(column);
    } else {
      ReactPanel.currentPanel = new ReactPanel(
        extensionPath,
        column || vscode.ViewColumn.One
      );
    }

    ReactPanel.currentPanel.updateWebviewContent();
  }

  private updateFileContent(content: string) {
    if (ReactPanel.activeEditor) {
      const edit = new vscode.WorkspaceEdit();
      const document = ReactPanel.activeEditor.document;

      // Update the entire document content
      edit.replace(
        document.uri,
        new vscode.Range(0, 0, document.lineCount, 0),
        content
      );

      vscode.workspace.applyEdit(edit);
    } else {
      // const editor = vscode.window.activeTextEditor;
      // if (editor && editor.document.languageId === 'yaml') {
      // 	const edit = new vscode.WorkspaceEdit();
      // 	edit.replace(editor.document.uri, new vscode.Range(0, 0, editor.document.lineCount, 0), content);
      // 	vscode.workspace.applyEdit(edit);
      // }
      vscode.window.showErrorMessage(
        "Unable to update the YAML file. No editor is active."
      );
    }
  }

  public updateWebviewContent(newData?: string) {
    if (vscode.window.activeTextEditor) {
      const document = vscode.window.activeTextEditor.document;
      if (document.languageId === "yaml") {
        this._panel.webview.postMessage({
          command: "updateContent",
          content: newData || document.getText(),
        });
      }
    }
  }

  private constructor(extensionPath: string, column: vscode.ViewColumn) {
    this._extensionPath = extensionPath;

    // Create and show a new webview panel
    this._panel = vscode.window.createWebviewPanel(
      ReactPanel.viewType,
      "Module3 Designer",
      column,
      {
        // Enable javascript in the webview
        enableScripts: true,

        // And restric the webview to only loading content from our extension's `media` directory.
        localResourceRoots: [
          vscode.Uri.file(path.join(this._extensionPath, "out")),
        ],
      }
    );

    if (vscode.window.activeTextEditor) {
      const document = vscode.window.activeTextEditor.document;
      // if (document.languageId === 'yaml') {
      this._panel.webview.postMessage({
        command: "updateContent",
        content: document.getText(),
      });
      // }
    }

    // Set the webview's initial html content
    this._panel.webview.html = this._getHtmlForWebview();

    // Listen for when the panel is disposed
    // This happens when the user closes the panel or when the panel is closed programatically
    this._panel.onDidDispose(() => this.dispose(), null, this._disposables);

    // Handle messages from the webview

    this._panel.webview.onDidReceiveMessage(
      (message) => {
        switch (message.command) {
          case "alert":
            vscode.window.showErrorMessage(message.text);
            return;

          case "updateFile":
            this.updateFileContent(message.content);
            break;
        }
      },
      null,
      this._disposables
    );
  }

  public doRefactor() {
    // Send a message to the webview webview.
    // You can send any JSON serializable data.
    this._panel.webview.postMessage({ command: "refactor" });
  }

  public dispose() {
    ReactPanel.currentPanel = undefined;

    // Clean up our resources
    this._panel.dispose();

    while (this._disposables.length) {
      const x = this._disposables.pop();
      if (x) {
        x.dispose();
      }
    }
  }

  private _getHtmlForWebview() {
    // const manifest = require(path.join(
    //   this._extensionPath,
    //   "out",
    //   "asset-manifest.json"
    // ));
    // const mainScript = manifest["files"]["main.js"];
    const mainScript = "static/js/main.c0aba94a.js";
    const mainStyle = "static/css/main.f7f07ed2.css";
    const indexFile = "index.html";

    const scriptPathOnDisk = vscode.Uri.file(
      path.join(this._extensionPath, "out", mainScript)
    );
    const indexOnDisk = vscode.Uri.file(
      path.join(this._extensionPath, "out", indexFile)
    );

    const filesPrefix = vscode.Uri.file(path.join(this._extensionPath, "out"))
      .toString()
      .replace("file://", "https://file%2B.vscode-resource.vscode-cdn.net");

    const scriptUri = scriptPathOnDisk
      .with({})
      .toString()
      .replace("file://", "https://file%2B.vscode-resource.vscode-cdn.net");

    const stylePathOnDisk = vscode.Uri.file(
      path.join(this._extensionPath, "out", mainStyle)
    );
    const styleUri = stylePathOnDisk
      .with({})
      .toString()
      .replace("file://", "https://file%2B.vscode-resource.vscode-cdn.net");

    // Use a nonce to whitelist which scripts can be run
    const nonce = getNonce();

    const basehref = `<base href="${vscode.Uri.file(
      path.join(this._extensionPath, "out")
    ).with({
      scheme: "vscode-resource",
    })}/">`;

    const htmlData = readFileSync(indexOnDisk.toString().replace("file://", ""))
      .toString()
      .replaceAll("/static", filesPrefix + "/static")
      .replaceAll("</title>", "</title>" + basehref)
      .replaceAll('defer="defer"', `defer="defer" nonce="${nonce}"`);

    return htmlData;

    return `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Security-Policy" content="default-src 'none'; img-src vscode-resource: https:; script-src 'nonce-${nonce}';style-src vscode-resource: 'unsafe-inline' http: https: data:;">
    <link rel="stylesheet" type="text/css" href="${styleUri}">
    
    <base href="${vscode.Uri.file(path.join(this._extensionPath, "out")).with({
      scheme: "vscode-resource",
    })}/">
  </head>
  <body>
    Hello ${nonce}
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <div id="root">No render</div>
    <script nonce="${nonce}" src="${scriptUri}"></script>
    
  </body>
</html>
`;

    // return `<!DOCTYPE html>
    //           <html lang="en">
    //           <head>
    //               <meta charset="utf-8">

    //               <meta name="theme-color" content="#000000">
    //               <title>Module3 Designer</title>
    //               <link rel="stylesheet" type="text/css" href="${styleUri}">
    //               <meta http-equiv="Content-Security-Policy" content="default-src 'none'; img-src vscode-resource: https:; script-src 'nonce-${nonce}';style-src vscode-resource: 'unsafe-inline' http: https: data:;">
    //               <base href="${vscode.Uri.file(
    //                 path.join(this._extensionPath, "out")
    //               ).with({
    //                 scheme: "vscode-resource",
    //               })}/">
    //           </head>

    //           <body>
    //               <noscript>You need to enable JavaScript to run this app.</noscript>
    //               <div id="root"></div>
    //               Hiiii

    //               <script nonce="${nonce}" src="${scriptUri}"></script>
    //           </body>
    //           </html>`;
  }
}

function getNonce() {
  let text = "";
  const possible =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  for (let i = 0; i < 32; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}
