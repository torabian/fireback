import Foundation
class HmiComponentTypeEntity : Codable, Identifiable {
    var name: String? = nil
    var isDirectInteractable: Bool? = nil
}
class HmiComponentTypeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var isDirectInteractable: Bool? = nil
  @Published var isDirectInteractableErrorMessage: Bool? = nil
  func getDto() -> HmiComponentTypeEntity {
      var dto = HmiComponentTypeEntity()
    dto.name = self.name
    dto.isDirectInteractable = self.isDirectInteractable
      return dto
  }
}