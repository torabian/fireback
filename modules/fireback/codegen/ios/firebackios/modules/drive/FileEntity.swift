class FileEntity : Codable {
    private var _name: String?
    var `Name`: String? {
        set { _name = newValue }
        get { return _name }
    }
    private var _diskPath: String?
    var `DiskPath`: String? {
        set { _diskPath = newValue }
        get { return _diskPath }
    }
    private var _size: Int?
    var `Size`: Int? {
        set { _size = newValue }
        get { return _size }
    }
    private var _virtualPath: String?
    var `VirtualPath`: String? {
        set { _virtualPath = newValue }
        get { return _virtualPath }
    }
    private var _type: String?
    var `Type`: String? {
        set { _type = newValue }
        get { return _type }
    }
}