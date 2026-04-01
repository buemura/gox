import * as path from "path";
import { workspace, ExtensionContext } from "vscode";
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient | undefined;

export function activate(context: ExtensionContext) {
  const config = workspace.getConfiguration("gox");
  const goxPath = config.get<string>("lsp.path") || "gox";

  const serverOptions: ServerOptions = {
    command: goxPath,
    args: ["lsp"],
    transport: TransportKind.stdio,
  };

  const clientOptions: LanguageClientOptions = {
    documentSelector: [{ scheme: "file", language: "gox" }],
    synchronize: {
      fileEvents: workspace.createFileSystemWatcher("**/*.gox"),
    },
  };

  client = new LanguageClient(
    "gox-lsp",
    "Gox Language Server",
    serverOptions,
    clientOptions
  );

  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
