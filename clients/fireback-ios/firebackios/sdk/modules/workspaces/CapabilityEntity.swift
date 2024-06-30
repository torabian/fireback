import Foundation
class CapabilityEntity : Codable, Identifiable {
    var name: String? = nil
    var description: String? = nil
}
class CapabilityEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  func getDto() -> CapabilityEntity {
      var dto = CapabilityEntity()
    dto.name = self.name
    dto.description = self.description
      return dto
  }
}