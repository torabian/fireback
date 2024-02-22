import Foundation
class WidgetEntity : Codable, Identifiable {
    var name: String? = nil
    var family: String? = nil
    var providerKey: String? = nil
}
class WidgetEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var family: String? = nil
  @Published var familyErrorMessage: String? = nil
  @Published var providerKey: String? = nil
  @Published var providerKeyErrorMessage: String? = nil
  func getDto() -> WidgetEntity {
      var dto = WidgetEntity()
    dto.name = self.name
    dto.family = self.family
    dto.providerKey = self.providerKey
      return dto
  }
}