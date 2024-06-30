package workspaces

// func AuthorizeSocketConnection(so socketio.Socket) {
// 	token := so.Request().URL.Query()["access_token"][0]
// 	user, err := GetUserFromToken(token)

// 	workspacesList := GetUserWorkspaces(user.UniqueId)

// 	fmt.Println(workspacesList)
// 	so.On("disconnect", func() {})

// 	if err != nil {
// 		so.Emit("error", "Unauthorized access")
// 		return
// 	}

// 	userRoom := "user_" + user.UniqueId
// 	fmt.Println("Joining user room", userRoom)
// 	so.Join(userRoom)

// 	for _, workspace := range workspacesList {
// 		workspaceRoom := "workspace_" + *workspace.WorkspaceId
// 		fmt.Println("Workspace: " + workspaceRoom)
// 		so.Join(workspaceRoom)
// 	}
// }
