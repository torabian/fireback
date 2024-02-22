import Foundation
struct UserAccessLevelDto : Codable {
    var capabilities: [String]? = nil
    var workspaces: [String]? = nil
    var SQL: String? = nil
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
class UserAccessLevelDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var SQL: String? = nil
  @Published var SQLErrorMessage: String? = nil
  func getDto() -> UserAccessLevelDto {
      var dto = UserAccessLevelDto()
    dto.SQL = self.SQL
      return dto
  }
}
