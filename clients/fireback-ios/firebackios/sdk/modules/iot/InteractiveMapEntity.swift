import Foundation
class InteractiveMapEntity : Codable, Identifiable {
    var name: String? = nil
}
class InteractiveMapEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> InteractiveMapEntity {
      var dto = InteractiveMapEntity()
    dto.name = self.name
      return dto
  }
}