import Foundation
class RoleEntity : Codable, Identifiable {
    var name: String? = nil
    var capabilities: [CapabilityEntity]? = nil
    var capabilitiesListId: [String]? = nil
}
class RoleEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> RoleEntity {
      var dto = RoleEntity()
    dto.name = self.name
      return dto
  }
}