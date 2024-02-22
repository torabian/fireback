import Foundation
class WorkspaceInviteEntity : Codable, Identifiable {
    var coverLetter: String? = nil
    var targetUserLocale: String? = nil
    var value: String? = nil
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
    var firstName: String? = nil
    var lastName: String? = nil
    var used: Bool? = nil
    var role: RoleEntity? = nil
    // var roleId: String? = nil
}
class WorkspaceInviteEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var coverLetter: String? = nil
  @Published var coverLetterErrorMessage: String? = nil
  @Published var targetUserLocale: String? = nil
  @Published var targetUserLocaleErrorMessage: String? = nil
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  @Published var used: Bool? = nil
  @Published var usedErrorMessage: Bool? = nil
  func getDto() -> WorkspaceInviteEntity {
      var dto = WorkspaceInviteEntity()
    dto.coverLetter = self.coverLetter
    dto.targetUserLocale = self.targetUserLocale
    dto.value = self.value
    dto.firstName = self.firstName
    dto.lastName = self.lastName
    dto.used = self.used
      return dto
  }
}