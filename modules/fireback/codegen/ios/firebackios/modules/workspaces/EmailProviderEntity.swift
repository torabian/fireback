class EmailProviderEntity : Codable {
    private var _type: String?
    var `Type`: String? {
        set { _type = newValue }
        get { return _type }
    }
    private var _apiKey: String?
    var `ApiKey`: String? {
        set { _apiKey = newValue }
        get { return _apiKey }
    }
}