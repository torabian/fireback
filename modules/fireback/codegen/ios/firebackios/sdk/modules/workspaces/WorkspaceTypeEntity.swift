import Foundation
class WorkspaceTypeEntity : Codable, Identifiable {
    var title: String? = nil
    var description: String? = nil
    var slug: String? = nil
    var role: RoleEntity? = nil
    // var roleId: String? = nil
}
class WorkspaceTypeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var title: String? = nil
  @Published var titleErrorMessage: String? = nil
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  @Published var slug: String? = nil
  @Published var slugErrorMessage: String? = nil
  func getDto() -> WorkspaceTypeEntity {
      var dto = WorkspaceTypeEntity()
    dto.title = self.title
    dto.description = self.description
    dto.slug = self.slug
      return dto
  }
}