class AppMenuEntity : Codable {
    private var _href: String?
    var `Href`: String? {
        set { _href = newValue }
        get { return _href }
    }
    private var _icon: String?
    var `Icon`: String? {
        set { _icon = newValue }
        get { return _icon }
    }
    private var _label: String?
    var `Label`: String? {
        set { _label = newValue }
        get { return _label }
    }
    private var _activeMatcher: String?
    var `ActiveMatcher`: String? {
        set { _activeMatcher = newValue }
        get { return _activeMatcher }
    }
    private var _applyType: String?
    var `ApplyType`: String? {
        set { _applyType = newValue }
        get { return _applyType }
    }
    private var _capability: CapabilityEntity?
    var `Capability`: CapabilityEntity? {
        set { _capability = newValue }
        get { return _capability }
    }
}