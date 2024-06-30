package workspaces

func BackupTableMetaActionCreate(
	dto *BackupTableMetaEntity, query QueryDSL,
) (*BackupTableMetaEntity, *IError) {
	return BackupTableMetaActionCreateFn(dto, query)
}

func BackupTableMetaActionUpdate(
	query QueryDSL,
	fields *BackupTableMetaEntity,
) (*BackupTableMetaEntity, *IError) {
	return BackupTableMetaActionUpdateFn(query, fields)
}
