import Foundation
class WorkspaceEntity : Codable, Identifiable {
    var description: String? = nil
    var name: String? = nil
    var type: WorkspaceTypeEntity? = nil
    // var typeId: String? = nil
}
class WorkspaceEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> WorkspaceEntity {
      var dto = WorkspaceEntity()
    dto.description = self.description
    dto.name = self.name
      return dto
  }
}