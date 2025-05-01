import Foundation
struct WorkspaceDto : Codable {
    var relations: [UserRoleWorkspaceEntity]? = nil
    var relationsListId: [String]? = nil
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
class WorkspaceDtoViewModel: ObservableObject {
    // improve the fields here
    @Published var visibility: String? = nil
    @Published var updated: Int? = nil
    @Published var created: Int? = nil
    func getDto() -> WorkspaceDto {
        var dto = WorkspaceDto()
        dto.visibility = self.visibility
        dto.updated = self.updated
        dto.created = self.created
        return dto
    }
}