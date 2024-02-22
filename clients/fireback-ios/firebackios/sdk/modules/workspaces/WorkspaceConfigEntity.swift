import Foundation
class WorkspaceConfigEntity : Codable, Identifiable {
    var disablePublicWorkspaceCreation: Int? = nil
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
    var zoomClientId: String? = nil
    var zoomClientSecret: String? = nil
    var allowPublicToJoinTheWorkspace: Bool? = nil
}
class WorkspaceConfigEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var disablePublicWorkspaceCreation: Int? = nil
  @Published var disablePublicWorkspaceCreationErrorMessage: Int? = nil
  @Published var zoomClientId: String? = nil
  @Published var zoomClientIdErrorMessage: String? = nil
  @Published var zoomClientSecret: String? = nil
  @Published var zoomClientSecretErrorMessage: String? = nil
  @Published var allowPublicToJoinTheWorkspace: Bool? = nil
  @Published var allowPublicToJoinTheWorkspaceErrorMessage: Bool? = nil
  func getDto() -> WorkspaceConfigEntity {
      var dto = WorkspaceConfigEntity()
    dto.disablePublicWorkspaceCreation = self.disablePublicWorkspaceCreation
    dto.zoomClientId = self.zoomClientId
    dto.zoomClientSecret = self.zoomClientSecret
    dto.allowPublicToJoinTheWorkspace = self.allowPublicToJoinTheWorkspace
      return dto
  }
}