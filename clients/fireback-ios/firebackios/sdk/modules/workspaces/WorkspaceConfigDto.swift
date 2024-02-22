import Foundation
struct WorkspaceConfigDto : Codable {
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
    var workspaceId: String? = nil
    var zoomClientId: String? = nil
    var zoomClientSecret: String? = nil
    var allowPublicToJoinTheWorkspace: Bool? = nil
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
class WorkspaceConfigDtoViewModel: ObservableObject {
    // improve the fields here
    @Published var workspaceId: String? = nil
    @Published var zoomClientId: String? = nil
    @Published var zoomClientSecret: String? = nil
    @Published var allowPublicToJoinTheWorkspace: Bool? = nil
    @Published var visibility: String? = nil
    @Published var updated: Int? = nil
    @Published var created: Int? = nil
    func getDto() -> WorkspaceConfigDto {
        var dto = WorkspaceConfigDto()
        dto.workspaceId = self.workspaceId
        dto.zoomClientId = self.zoomClientId
        dto.zoomClientSecret = self.zoomClientSecret
        dto.allowPublicToJoinTheWorkspace = self.allowPublicToJoinTheWorkspace
        dto.visibility = self.visibility
        dto.updated = self.updated
        dto.created = self.created
        return dto
    }
}