class BackupTableMetaEntity : Codable {
    private var _tableNameInDb: String?
    var `TableNameInDb`: String? {
        set { _tableNameInDb = newValue }
        get { return _tableNameInDb }
    }
}