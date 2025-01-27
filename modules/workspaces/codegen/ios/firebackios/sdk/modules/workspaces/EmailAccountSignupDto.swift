import Foundation
struct EmailAccountSignupDto : Codable {
    var email: String? = nil
    var password: String? = nil
    var firstName: String? = nil
    var lastName: String? = nil
    var inviteId: String? = nil
    var publicJoinKeyId: String? = nil
    var workspaceTypeId: String? = nil
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
class EmailAccountSignupDtoViewModel: ObservableObject {
    // improve the fields here
    @Published var email: String? = nil
    @Published var password: String? = nil
    @Published var firstName: String? = nil
    @Published var lastName: String? = nil
    @Published var inviteId: String? = nil
    @Published var publicJoinKeyId: String? = nil
    @Published var workspaceTypeId: String? = nil
    func getDto() -> EmailAccountSignupDto {
        var dto = EmailAccountSignupDto()
        dto.email = self.email
        dto.password = self.password
        dto.firstName = self.firstName
        dto.lastName = self.lastName
        dto.inviteId = self.inviteId
        dto.publicJoinKeyId = self.publicJoinKeyId
        dto.workspaceTypeId = self.workspaceTypeId
        return dto
    }
}