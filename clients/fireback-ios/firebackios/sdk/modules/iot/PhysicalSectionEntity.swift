import Foundation
class PhysicalSectionEntity : Codable, Identifiable {
    var name: String? = nil
}
class PhysicalSectionEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> PhysicalSectionEntity {
      var dto = PhysicalSectionEntity()
    dto.name = self.name
      return dto
  }
}