//go:build wasm

package fireback

import "github.com/torabian/emi/emigo"

// ReactiveSearchActionReactiveHandlerWasm is the in-browser counterpart of
// ReactiveSearchActionReactiveHandler. It registers the same developer factory against an
// emigo.WasmReactor instead of a gin engine, so the business logic is shared
// verbatim across the real server and the wasm build.
//
// There is no gorilla, no gin, and no socket here: the reactor bridges the
// session's channels straight to the WebSocketWasm JS class.
//
func ReactiveSearchActionReactiveHandlerWasm(
	reactor *emigo.WasmReactor,
	factory func(session ReactiveSearchActionSession) (chan []byte, error),
) {
	reactor.Handle(ReactiveSearchActionMeta().URL, func(conn *emigo.WasmReactiveConn) error {
		session := ReactiveSearchActionSession{
			Done:        conn.Done,
			Read:        make(chan ReactiveSearchActionReadChan),
			QueryParams: ReactiveSearchActionQueryFromString(conn.Query.Encode()),
			// Socket and Ctx stay nil — meaningless in the browser.
		}
		// client -> server: adapt raw []byte frames into the typed read channel.
		go func() {
			for {
				select {
				case data, ok := <-conn.Read:
					if !ok {
						return
					}
					session.Read <- ReactiveSearchActionReadChan{Data: data}
				case <-conn.Done:
					return
				}
			}
		}()
		write, err := factory(session)
		if err != nil {
			return err
		}
		// server -> client: forward the factory's write channel to the bridge.
		// Closing write signals a server-initiated close to the reactor.
		go func() {
			for {
				select {
				case msg, ok := <-write:
					if !ok {
						close(conn.Write)
						return
					}
					conn.Write <- msg
				case <-conn.Done:
					return
				}
			}
		}()
		return nil
	})
}
