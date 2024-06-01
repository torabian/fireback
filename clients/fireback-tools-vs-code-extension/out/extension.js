"use strict";
// // The module 'vscode' contains the VS Code extensibility API
// // Import the module and reference it with the alias vscode in your code below
// import * as vscode from 'vscode';
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.deactivate = exports.activate = void 0;
const vscode = __importStar(require("vscode"));
const node_1 = require("vscode-languageclient/node");
let client;
function activate(context) {
    // Get the path to the LSP server executable
    const serverPath = vscode.workspace
        .getConfiguration("yamlLsp")
        .get("serverPath");
    if (!serverPath) {
        vscode.window.showErrorMessage("YAML LSP server path is not configured.");
        return;
    }
    // Define the server options
    let serverOptions = {
        run: {
            command: serverPath,
            transport: node_1.TransportKind.stdio,
            options: {
                env: {
                    ...process.env,
                    LSP: "true",
                },
            },
        },
        debug: {
            command: serverPath,
            transport: node_1.TransportKind.stdio,
            options: {
                env: {
                    ...process.env,
                    LSP: "true",
                },
            },
        },
    };
    // Define the client options
    let clientOptions = {
        documentSelector: [{ scheme: "file", language: "yaml" }],
        synchronize: {
            fileEvents: vscode.workspace.createFileSystemWatcher("**/.clientrc"),
        },
    };
    // Create the language client and start the client.
    client = new node_1.LanguageClient("yamlLsp", "YAML LSP", serverOptions, clientOptions);
    // Start the client. This will also launch the server
    client.start();
}
exports.activate = activate;
function deactivate() {
    if (!client) {
        return undefined;
    }
    return client.stop();
}
exports.deactivate = deactivate;
//# sourceMappingURL=extension.js.map