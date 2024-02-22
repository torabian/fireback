import Foundation
class MovableObjectEntity : Codable, Identifiable {
    var name: String? = nil
    var interactiveMaps: [InteractiveMapEntity]? = nil
    var interactiveMapsListId: [String]? = nil
}
class MovableObjectEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> MovableObjectEntity {
      var dto = MovableObjectEntity()
    dto.name = self.name
      return dto
  }
}