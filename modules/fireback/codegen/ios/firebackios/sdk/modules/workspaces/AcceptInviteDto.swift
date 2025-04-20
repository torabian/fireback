import Foundation
struct AcceptInviteDto : Codable {
    var inviteUniqueId: String? = nil
    var visibility: String? = nil
    var updated: Int? = nil
    var created: Int? = nil
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
class AcceptInviteDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var inviteUniqueId: String? = nil
  @Published var inviteUniqueIdErrorMessage: String? = nil
  @Published var visibility: String? = nil
  @Published var visibilityErrorMessage: String? = nil
  @Published var updated: Int? = nil
  @Published var updatedErrorMessage: Int? = nil
  @Published var created: Int? = nil
  @Published var createdErrorMessage: Int? = nil
  func getDto() -> AcceptInviteDto {
      var dto = AcceptInviteDto()
    dto.inviteUniqueId = self.inviteUniqueId
    dto.visibility = self.visibility
    dto.updated = self.updated
    dto.created = self.created
      return dto
  }
}
