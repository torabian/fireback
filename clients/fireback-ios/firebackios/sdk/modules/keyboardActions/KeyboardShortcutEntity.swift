import Foundation
class KeyboardShortcutDefaultCombination : Codable, Identifiable {
    var altKey: Bool? = nil
    var key: String? = nil
    var metaKey: Bool? = nil
    var shiftKey: Bool? = nil
    var ctrlKey: Bool? = nil
}
class KeyboardShortcutUserCombination : Codable, Identifiable {
    var altKey: Bool? = nil
    var key: String? = nil
    var metaKey: Bool? = nil
    var shiftKey: Bool? = nil
    var ctrlKey: Bool? = nil
}
class KeyboardShortcutEntity : Codable, Identifiable {
    var os: String? = nil
    var host: String? = nil
    var defaultCombination: KeyboardShortcutDefaultCombination? = nil
    var userCombination: KeyboardShortcutUserCombination? = nil
    var action: String? = nil
    var actionKey: String? = nil
}
class KeyboardShortcutEntityViewModel: ObservableObject {
    // improve the fields here
    @Published var os: String? = nil
    @Published var host: String? = nil
    @Published var action: String? = nil
    @Published var actionKey: String? = nil
    func getDto() -> KeyboardShortcutEntity {
        var dto = KeyboardShortcutEntity()
        dto.os = self.os
        dto.host = self.host
        dto.action = self.action
        dto.actionKey = self.actionKey
        return dto
    }
}