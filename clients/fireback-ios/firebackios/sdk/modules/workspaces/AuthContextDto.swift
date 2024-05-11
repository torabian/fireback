import Foundation


struct AuthContextDto : Codable {
    var skipWorkspaceId: Bool? = nil
    var workspaceId: String? = nil
    var token: String? = nil
    var capabilities: [PermissionInfoDto]? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class AuthContextDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var skipWorkspaceId: Bool? = nil
  @Published var skipWorkspaceIdErrorMessage: Bool? = nil
  @Published var workspaceId: String? = nil
  @Published var workspaceIdErrorMessage: String? = nil
  @Published var token: String? = nil
  @Published var tokenErrorMessage: String? = nil
  func getDto() -> AuthContextDto {
      var dto = AuthContextDto()
    dto.skipWorkspaceId = self.skipWorkspaceId
    dto.workspaceId = self.workspaceId
    dto.token = self.token
      return dto
  }
}
