import Foundation
struct UserRoleWorkspaceDto : Codable {
    var roleId: String? = nil
    var capabilities: [String]? = nil
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
class UserRoleWorkspaceDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var roleId: String? = nil
  @Published var roleIdErrorMessage: String? = nil
  func getDto() -> UserRoleWorkspaceDto {
      var dto = UserRoleWorkspaceDto()
    dto.roleId = self.roleId
      return dto
  }
}
