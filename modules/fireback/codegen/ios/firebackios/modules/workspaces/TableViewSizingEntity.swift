class TableViewSizingEntity : Codable {
    private var _tableName: String?
    var `TableName`: String? {
        set { _tableName = newValue }
        get { return _tableName }
    }
    private var _sizes: String?
    var `Sizes`: String? {
        set { _sizes = newValue }
        get { return _sizes }
    }
}