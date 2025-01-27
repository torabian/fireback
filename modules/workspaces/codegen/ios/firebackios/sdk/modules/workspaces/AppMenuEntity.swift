import Foundation
class AppMenuEntity : Codable, Identifiable {
    var href: String? = nil
    var icon: String? = nil
    var label: String? = nil
    var activeMatcher: String? = nil
    var applyType: String? = nil
    var capability: CapabilityEntity? = nil
    // var capabilityId: String? = nil
}
class AppMenuEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var href: String? = nil
  @Published var hrefErrorMessage: String? = nil
  @Published var icon: String? = nil
  @Published var iconErrorMessage: String? = nil
  @Published var label: String? = nil
  @Published var labelErrorMessage: String? = nil
  @Published var activeMatcher: String? = nil
  @Published var activeMatcherErrorMessage: String? = nil
  @Published var applyType: String? = nil
  @Published var applyTypeErrorMessage: String? = nil
  func getDto() -> AppMenuEntity {
      var dto = AppMenuEntity()
    dto.href = self.href
    dto.icon = self.icon
    dto.label = self.label
    dto.activeMatcher = self.activeMatcher
    dto.applyType = self.applyType
      return dto
  }
}