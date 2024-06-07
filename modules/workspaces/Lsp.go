package workspaces

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"

	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/ucarion/urlpath"
	"github.com/urfave/cli"
)

type stdioReadWriteCloser struct{}

var _ io.ReadWriteCloser = (*stdioReadWriteCloser)(nil)

func (c stdioReadWriteCloser) Read(p []byte) (n int, err error) {
	return os.Stdin.Read(p)
}

func (c stdioReadWriteCloser) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (c stdioReadWriteCloser) Close() error {
	return nil
}

func BeginLspServer(c *cli.Context) error {
	server := &Server{}

	handler := jsonrpc2.HandlerWithError(server.handle)

	<-jsonrpc2.NewConn(
		context.Background(),
		jsonrpc2.NewBufferedStream(stdioReadWriteCloser{}, jsonrpc2.VSCodeObjectCodec{}), handler,
	).DisconnectNotify()

	return nil
}

type Server struct {
	conn     *jsonrpc2.Conn
	document string
}

func (s *Server) completion(ctx context.Context, params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	completionItems := []lsp.CompletionItem{
		{
			Label:  "Item 1",
			Kind:   lsp.CIKKeyword,
			Detail: "Description of item 1",
		},
		{
			Label:  "Item 2",
			Kind:   lsp.CIKFunction,
			Detail: "Description of item 2",
		},
		{
			Label:  "Item 3",
			Kind:   lsp.CIKVariable,
			Detail: "Description of item 3",
		},
		{
			Label:  "Item 4",
			Kind:   lsp.CIKVariable,
			Detail: "Description of item 4",
		},
		{
			Label:  "Item 5",
			Kind:   lsp.CIKVariable,
			Detail: "Description of item 5",
		},
	}

	// Send the completion items to the client
	if err := s.conn.Notify(ctx, "textDocument/completion", completionItems); err != nil {
		return nil, err
	}

	return completionItems, nil
}

func HandleAutoCompletion(content []byte, line int, offset int) []lsp.CompletionItem {
	completionItems := []lsp.CompletionItem{}

	uri := GetContextFromYaml(content, line, offset)

	log.Println(line, offset, uri)
	for _, action := range actions {
		k := urlpath.New(action.Path)

		rawUrl, _ := removeQueryParams(uri)
		_, ok := k.Match(rawUrl)
		if ok {

			completionItems = append(completionItems, action.Handler()...)

			// Not sure, we might give multiple options
			break
		}

	}

	return completionItems
}
func (s *Server) handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {

	switch req.Method {
	case "initialize":
		var params lsp.InitializeParams

		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}

		return s.initialize(ctx, &params)
	case "textDocument/didOpen":
		var params lsp.DidOpenTextDocumentParams

		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return nil, s.didOpen(ctx, &params)
	case "textDocument/completion":

		var params lsp.TextDocumentPositionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}

		return HandleAutoCompletion([]byte(s.document), params.Position.Line+1, params.Position.Character), nil

	case "textDocument/didChange":

		var params lsp.DidChangeTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}

		s.document = params.ContentChanges[0].Text

		return nil, nil
	default:
		// conn.ReplyWithError(ctx, req.ID, &jsonrpc2.Error{
		// 	Code:    33,
		// 	Message: "23123",
		// })
		return nil, nil
	}
}

func (s *Server) initialize(ctx context.Context, params *lsp.InitializeParams) (*lsp.InitializeResult, error) {

	return &lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			CompletionProvider: &lsp.CompletionOptions{
				TriggerCharacters: []string{},
				ResolveProvider:   true,
			},
			TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
				Options: &lsp.TextDocumentSyncOptions{
					OpenClose: true,
					Change:    lsp.TDSKFull,
				},
			},
			// For incrementally
			// TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
			// 	Options: &lsp.TextDocumentSyncOptions{
			// 		OpenClose: true,
			// 		Change:    lsp.TDSKIncremental,
			// 	},
			// },
		},
	}, nil
}

func (s *Server) didOpen(ctx context.Context, params *lsp.DidOpenTextDocumentParams) error {
	content := params.TextDocument.Text

	return ValidateYAML(content)
}

func (s *Server) didChange(ctx context.Context, params *lsp.DidChangeTextDocumentParams) error {
	for _, change := range params.ContentChanges {
		if err := ValidateYAML(change.Text); err != nil {
			return err
		}
	}
	return nil
}

func ValidateYAML(content string) error {
	os.WriteFile("/tmp/4.log", []byte(content), 0644)
	if strings.TrimSpace(content) == "" {
		return errors.New("empty content")
	}
	// Simulate validation success
	return nil
}

// completionItems := []lsp.CompletionItem{
// 	{
// 		Label:  "entities:",
// 		Kind:   lsp.CIKProperty,
// 		Detail: "Set of entities",
// 	},
// }

// return completionItems, nil
// conn.Notify(ctx, "textDocument/completion", completionItems)
// return []lsp.CompletionItem{}, nil

// conn.Notify(ctx, "window/showMessage", lsp.ShowMessageParams{
// 	Type:    lsp.MTError,
// 	Message: "Something bad!",
// })

// Get the range of the line with the error
// start := lsp.Position{Line: 1, Character: 0} // Example line number (0-based)
// end := lsp.Position{Line: 1, Character: 5}   // Example line number (0-based)
// range2 := lsp.Range{Start: start, End: end}

// Create a diagnostic with an error message

// diagnostic := lsp.Diagnostic{
// 	Range:    range2,
// 	Severity: lsp.Error,
// 	Message:  "this is it",
// }
// Send a publish diagnostic notification to the client
// conn.Notify(ctx, "textDocument/publishDiagnostics", lsp.PublishDiagnosticsParams{
// 	URI:         "file:///Users/ali/test4/Test.yaml",
// 	Diagnostics: []lsp.Diagnostic{diagnostic},
// })
