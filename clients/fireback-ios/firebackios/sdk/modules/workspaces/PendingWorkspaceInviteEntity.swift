import Foundation
class PendingWorkspaceInviteEntity : Codable, Identifiable {
    var value: String? = nil
    var type: String? = nil
    var coverLetter: String? = nil
    var workspaceName: String? = nil
    var role: RoleEntity? = nil
    // var roleId: String? = nil
}
class PendingWorkspaceInviteEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  @Published var coverLetter: String? = nil
  @Published var coverLetterErrorMessage: String? = nil
  @Published var workspaceName: String? = nil
  @Published var workspaceNameErrorMessage: String? = nil
  func getDto() -> PendingWorkspaceInviteEntity {
      var dto = PendingWorkspaceInviteEntity()
    dto.value = self.value
    dto.type = self.type
    dto.coverLetter = self.coverLetter
    dto.workspaceName = self.workspaceName
      return dto
  }
}