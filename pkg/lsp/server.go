package lsp

import (
	"context"
	"io"
	"net/url"
	"path/filepath"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

// Server implements protocol.Server for .gox files.
type Server struct {
	client  protocol.Client
	docs    *DocumentStore
	index   *ComponentIndex
	rootURI string
	logger  *zap.Logger
}

// Serve starts the LSP server over stdio.
func Serve(ctx context.Context, in io.ReadCloser, out io.WriteCloser) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	rwc := struct {
		io.ReadCloser
		io.Writer
	}{in, out}

	stream := jsonrpc2.NewStream(nopWriteCloser{rwc})
	s := &Server{
		docs:   NewDocumentStore(),
		index:  NewComponentIndex(),
		logger: logger,
	}

	ctx, conn, client := protocol.NewServer(ctx, s, stream, logger)
	s.client = client
	_ = conn

	<-ctx.Done()
	return nil
}

// nopWriteCloser wraps an io.ReadCloser+io.Writer into an io.ReadWriteCloser.
type nopWriteCloser struct {
	rw interface {
		io.ReadCloser
		io.Writer
	}
}

func (n nopWriteCloser) Read(p []byte) (int, error)  { return n.rw.Read(p) }
func (n nopWriteCloser) Write(p []byte) (int, error)  { return n.rw.Write(p) }
func (n nopWriteCloser) Close() error                 { return n.rw.Close() }

// uriToPath converts a file:// URI to a filesystem path.
func uriToPath(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return filepath.FromSlash(u.Path)
}

// --- protocol.Server implementation ---

func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	if params.RootURI != "" {
		s.rootURI = string(params.RootURI)
	}

	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: &protocol.TextDocumentSyncOptions{
				OpenClose: true,
				Change:    protocol.TextDocumentSyncKindFull,
				Save:      &protocol.SaveOptions{IncludeText: true},
			},
			DefinitionProvider: true,
			HoverProvider:      true,
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "gox-lsp",
			Version: "0.1.0",
		},
	}, nil
}

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	// Index all .gox files in the workspace.
	rootPath := uriToPath(s.rootURI)
	if rootPath != "" {
		s.index.Scan(rootPath)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

func (s *Server) Exit(ctx context.Context) error {
	return nil
}

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	content := params.TextDocument.Text
	s.docs.Open(uri, content)
	s.diagnose(ctx, uri, content)
	s.index.IndexFileContent(uri, content)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	// We use Full sync, so the last change contains the entire document.
	if len(params.ContentChanges) > 0 {
		content := params.ContentChanges[len(params.ContentChanges)-1].Text
		s.docs.Update(uri, content)
		s.diagnose(ctx, uri, content)
	}
	return nil
}

func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	s.docs.Close(uri)
	// Clear diagnostics.
	s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         params.TextDocument.URI,
		Diagnostics: []protocol.Diagnostic{},
	})
	return nil
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	if params.Text != "" {
		s.index.IndexFileContent(uri, params.Text)
	}
	return nil
}

func (s *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	uri := string(params.TextDocument.URI)
	content, ok := s.docs.Get(uri)
	if !ok {
		return nil, nil
	}
	return s.findDefinition(uri, content, params.Position), nil
}

func (s *Server) Hover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	uri := string(params.TextDocument.URI)
	content, ok := s.docs.Get(uri)
	if !ok {
		return nil, nil
	}
	return s.computeHover(uri, content, params.Position), nil
}

// --- Unimplemented methods (return nil/empty) ---

func (s *Server) WorkDoneProgressCancel(context.Context, *protocol.WorkDoneProgressCancelParams) error {
	return nil
}
func (s *Server) LogTrace(context.Context, *protocol.LogTraceParams) error   { return nil }
func (s *Server) SetTrace(context.Context, *protocol.SetTraceParams) error   { return nil }
func (s *Server) CodeAction(context.Context, *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	return nil, nil
}
func (s *Server) CodeLens(context.Context, *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	return nil, nil
}
func (s *Server) CodeLensResolve(context.Context, *protocol.CodeLens) (*protocol.CodeLens, error) {
	return nil, nil
}
func (s *Server) ColorPresentation(context.Context, *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	return nil, nil
}
func (s *Server) Completion(context.Context, *protocol.CompletionParams) (*protocol.CompletionList, error) {
	return nil, nil
}
func (s *Server) CompletionResolve(context.Context, *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	return nil, nil
}
func (s *Server) Declaration(context.Context, *protocol.DeclarationParams) ([]protocol.Location, error) {
	return nil, nil
}
func (s *Server) DidChangeConfiguration(context.Context, *protocol.DidChangeConfigurationParams) error {
	return nil
}
func (s *Server) DidChangeWatchedFiles(context.Context, *protocol.DidChangeWatchedFilesParams) error {
	return nil
}
func (s *Server) DidChangeWorkspaceFolders(context.Context, *protocol.DidChangeWorkspaceFoldersParams) error {
	return nil
}
func (s *Server) DocumentColor(context.Context, *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	return nil, nil
}
func (s *Server) DocumentHighlight(context.Context, *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	return nil, nil
}
func (s *Server) DocumentLink(context.Context, *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	return nil, nil
}
func (s *Server) DocumentLinkResolve(context.Context, *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	return nil, nil
}
func (s *Server) DocumentSymbol(context.Context, *protocol.DocumentSymbolParams) ([]interface{}, error) {
	return nil, nil
}
func (s *Server) ExecuteCommand(context.Context, *protocol.ExecuteCommandParams) (interface{}, error) {
	return nil, nil
}
func (s *Server) FoldingRanges(context.Context, *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	return nil, nil
}
func (s *Server) Formatting(context.Context, *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	return nil, nil
}
func (s *Server) Implementation(context.Context, *protocol.ImplementationParams) ([]protocol.Location, error) {
	return nil, nil
}
func (s *Server) OnTypeFormatting(context.Context, *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	return nil, nil
}
func (s *Server) PrepareRename(context.Context, *protocol.PrepareRenameParams) (*protocol.Range, error) {
	return nil, nil
}
func (s *Server) RangeFormatting(context.Context, *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	return nil, nil
}
func (s *Server) References(context.Context, *protocol.ReferenceParams) ([]protocol.Location, error) {
	return nil, nil
}
func (s *Server) Rename(context.Context, *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	return nil, nil
}
func (s *Server) SignatureHelp(context.Context, *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	return nil, nil
}
func (s *Server) Symbols(context.Context, *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	return nil, nil
}
func (s *Server) TypeDefinition(context.Context, *protocol.TypeDefinitionParams) ([]protocol.Location, error) {
	return nil, nil
}
func (s *Server) WillSave(context.Context, *protocol.WillSaveTextDocumentParams) error {
	return nil
}
func (s *Server) WillSaveWaitUntil(context.Context, *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	return nil, nil
}
func (s *Server) ShowDocument(context.Context, *protocol.ShowDocumentParams) (*protocol.ShowDocumentResult, error) {
	return nil, nil
}
func (s *Server) WillCreateFiles(context.Context, *protocol.CreateFilesParams) (*protocol.WorkspaceEdit, error) {
	return nil, nil
}
func (s *Server) DidCreateFiles(context.Context, *protocol.CreateFilesParams) error {
	return nil
}
func (s *Server) WillRenameFiles(context.Context, *protocol.RenameFilesParams) (*protocol.WorkspaceEdit, error) {
	return nil, nil
}
func (s *Server) DidRenameFiles(context.Context, *protocol.RenameFilesParams) error {
	return nil
}
func (s *Server) WillDeleteFiles(context.Context, *protocol.DeleteFilesParams) (*protocol.WorkspaceEdit, error) {
	return nil, nil
}
func (s *Server) DidDeleteFiles(context.Context, *protocol.DeleteFilesParams) error {
	return nil
}
func (s *Server) CodeLensRefresh(context.Context) error { return nil }
func (s *Server) PrepareCallHierarchy(context.Context, *protocol.CallHierarchyPrepareParams) ([]protocol.CallHierarchyItem, error) {
	return nil, nil
}
func (s *Server) IncomingCalls(context.Context, *protocol.CallHierarchyIncomingCallsParams) ([]protocol.CallHierarchyIncomingCall, error) {
	return nil, nil
}
func (s *Server) OutgoingCalls(context.Context, *protocol.CallHierarchyOutgoingCallsParams) ([]protocol.CallHierarchyOutgoingCall, error) {
	return nil, nil
}
func (s *Server) SemanticTokensFull(context.Context, *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	return nil, nil
}
func (s *Server) SemanticTokensFullDelta(context.Context, *protocol.SemanticTokensDeltaParams) (interface{}, error) {
	return nil, nil
}
func (s *Server) SemanticTokensRange(context.Context, *protocol.SemanticTokensRangeParams) (*protocol.SemanticTokens, error) {
	return nil, nil
}
func (s *Server) SemanticTokensRefresh(context.Context) error { return nil }
func (s *Server) LinkedEditingRange(context.Context, *protocol.LinkedEditingRangeParams) (*protocol.LinkedEditingRanges, error) {
	return nil, nil
}
func (s *Server) Moniker(context.Context, *protocol.MonikerParams) ([]protocol.Moniker, error) {
	return nil, nil
}
func (s *Server) Request(context.Context, string, interface{}) (interface{}, error) {
	return nil, nil
}
