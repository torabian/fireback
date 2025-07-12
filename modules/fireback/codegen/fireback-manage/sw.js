self.addEventListener("push", async (event) => {
  const data = event.data ? event.data.text() : "No payload";

  console.log("Push:", data);

  // Show the notification
  event.waitUntil(
    self.registration.showNotification("🔥 New Notification", {
      body: data,
      icon: "/icon.png",
      badge: "/badge.png",
    })
  );

  // Send data to all open tabs (clients)
  const allClients = await self.clients.matchAll({ includeUncontrolled: true });
  for (const client of allClients) {
    client.postMessage({ type: "PUSH_RECEIVED", payload: data });
  }
});
