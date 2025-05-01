package org.example
import okhttp3.*
import kotlinx.coroutines.*
import kotlinx.coroutines.flow.*

fun useEventBusSubscription(): Pair<SharedFlow<String>, () -> Unit> {
    val hook = WebSocketHook(
        url = "/audio",
    )
    return hook.messages to hook::connect
}