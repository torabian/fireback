import Foundation
class CapabilityEntity : Codable, Identifiable {
    var name: String? = nil
}
class CapabilityEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> CapabilityEntity {
      var dto = CapabilityEntity()
    dto.name = self.name
      return dto
  }
}