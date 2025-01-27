class GsmProviderEntity : Codable {
    private var _apiKey: String?
    var `ApiKey`: String? {
        set { _apiKey = newValue }
        get { return _apiKey }
    }
    private var _mainSenderNumber: String?
    var `MainSenderNumber`: String? {
        set { _mainSenderNumber = newValue }
        get { return _mainSenderNumber }
    }
    private var _type: String?
    var `Type`: String? {
        set { _type = newValue }
        get { return _type }
    }
    private var _invokeUrl: String?
    var `InvokeUrl`: String? {
        set { _invokeUrl = newValue }
        get { return _invokeUrl }
    }
    private var _invokeBody: String?
    var `InvokeBody`: String? {
        set { _invokeBody = newValue }
        get { return _invokeBody }
    }
}