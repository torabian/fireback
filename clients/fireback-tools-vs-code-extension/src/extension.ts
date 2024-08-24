// // The module 'vscode' contains the VS Code extensibility API
// // Import the module and reference it with the alias vscode in your code below
// import * as vscode from 'vscode';

// // This method is called when your extension is activated
// // Your extension is activated the very first time the command is executed
// export function activate(context: vscode.ExtensionContext) {

// 	// Use the console to output diagnostic information (console.log) and errors (console.error)
// 	// This line of code will only be executed once when your extension is activated
// 	console.log('Congratulations, your extension "fireback-tools" is now active!');

// 	// The command has been defined in the package.json file
// 	// Now provide the implementation of the command with registerCommand
// 	// The commandId parameter must match the command field in package.json
// 	let disposable = vscode.commands.registerCommand('fireback-tools.helloWorld', () => {
// 		// The code you place here will be executed every time your command is executed
// 		// Display a message box to the user
// 		vscode.window.showInformationMessage('Hello World from fireback-tools!');
// 	});

// 	context.subscriptions.push(disposable);
// }

// // This method is called when your extension is deactivated
// export function deactivate() {}

import * as path from "path";
import * as os from "os";
import { existsSync } from "fs";

import * as vscode from "vscode";
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";
import { ReactPanel } from "./panel";

let client: LanguageClient;

export function activate(context: vscode.ExtensionContext) {


  context.subscriptions.push(
    vscode.commands.registerCommand("module3-designer.start", () => {
      ReactPanel.activeEditor = vscode.window.activeTextEditor;
      ReactPanel.createOrShow(context.extensionPath);
    })
  );

  // Get the path to the LSP server executable
  let serverPath = vscode.workspace
    .getConfiguration("fireback")
    .get<string>("serverPath");

  if (serverPath === "[autodetect]" || serverPath === "") {
    const platform = os.platform();
    switch (platform) {
      case "win32":
        serverPath = "C:\\Program Files (x86)\\Torabian\\fireback.exe";
        break;
      case "darwin":
        serverPath = "/usr/local/bin/fireback";
        break;
      case "linux":
        serverPath = "/usr/local/bin/fireback";
        break;
      default:
        vscode.window.showErrorMessage(`Unsupported platform: ${platform}`);
        return;
    }
  }

  if (!serverPath) {
    vscode.window.showErrorMessage("Fireback path is not configured (" + os.platform() +")");
    return;
  }

  // Check if the server path exists
  if (!existsSync(serverPath)) {
    vscode.window.showErrorMessage(
      `Fireback executable not found at path: ${serverPath}. You can change it in settings.json in
      "fireback.serverPath". Set the value to [autodetect] to use default installation. If not installed,
      check https://github.com/torabian/fireback to download installers`
    );
    return;
  }

  // Define the server options
  let serverOptions: ServerOptions = {
    run: {
      command: serverPath,
      transport: TransportKind.stdio,
      options: {
        env: {
          ...process.env,
          LSP: "true",
        },
      },
    },
    debug: {
      command: serverPath,
      transport: TransportKind.stdio,
      options: {
        env: {
          ...process.env,
          LSP: "true",
        },
      },
    },
  };

  // Define the client options
  let clientOptions: LanguageClientOptions = {
    documentSelector: [{ scheme: "file", language: "yaml" , pattern: "**/*Module3.{yml,yaml}" }],
    synchronize: {
      fileEvents: vscode.workspace.createFileSystemWatcher("**/.clientrc"),
    },
  };

  // Create the language client and start the client.
  client = new LanguageClient(
    "fireback",
    "Fireback Language Server",
    serverOptions,
    clientOptions
  );

  // Start the client. This will also launch the server
  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
