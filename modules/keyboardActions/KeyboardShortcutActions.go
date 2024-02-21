package keyboardActions

import "github.com/torabian/fireback/modules/workspaces"

func KeyboardShortcutActionCreate(
	dto *KeyboardShortcutEntity, query workspaces.QueryDSL,
) (*KeyboardShortcutEntity, *workspaces.IError) {
	return KeyboardShortcutActionCreateFn(dto, query)
}

func KeyboardShortcutActionUpdate(
	query workspaces.QueryDSL,
	fields *KeyboardShortcutEntity,
) (*KeyboardShortcutEntity, *workspaces.IError) {
	return KeyboardShortcutActionUpdateFn(query, fields)
}
