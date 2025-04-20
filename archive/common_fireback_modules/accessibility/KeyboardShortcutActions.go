package accessibility

import "github.com/torabian/fireback/modules/fireback"

func KeyboardShortcutActionCreate(
	dto *KeyboardShortcutEntity, query fireback.QueryDSL,
) (*KeyboardShortcutEntity, *fireback.IError) {
	return KeyboardShortcutActionCreateFn(dto, query)
}

func KeyboardShortcutActionUpdate(
	query fireback.QueryDSL,
	fields *KeyboardShortcutEntity,
) (*KeyboardShortcutEntity, *fireback.IError) {
	return KeyboardShortcutActionUpdateFn(query, fields)
}
