package org.example
import okhttp3.*
import kotlinx.coroutines.*
import kotlinx.coroutines.flow.*

class WebSocketHook(
    private val url: String
) {
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())

    private val _messages = MutableSharedFlow<String>()
    val messages: SharedFlow<String> = _messages

    private val client = OkHttpClient()

    fun connect() {
        val request = Request.Builder()
            .url(FirebackConfig.host + url)
            .header("authorization", FirebackConfig.token)
            .header("workspace-id", FirebackConfig.workspaceId)
            .build()

        val listener = object : WebSocketListener() {
            override fun onOpen(ws: WebSocket, response: Response) {
                ws.send("Connected")
            }

            override fun onMessage(ws: WebSocket, text: String) {
                scope.launch { _messages.emit(text) }
            }

            override fun onClosed(ws: WebSocket, code: Int, reason: String) {
                println("Closed: $reason")
            }

            override fun onFailure(ws: WebSocket, t: Throwable, response: Response?) {
                println("Error: ${t.message}")
            }
        }

        client.newWebSocket(request, listener)
    }
}
