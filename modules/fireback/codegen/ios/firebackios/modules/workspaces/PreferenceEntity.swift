class PreferenceEntity : Codable {
    private var _timezone: String?
    var `Timezone`: String? {
        set { _timezone = newValue }
        get { return _timezone }
    }
}