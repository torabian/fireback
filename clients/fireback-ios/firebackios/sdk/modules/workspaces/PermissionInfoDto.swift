import Foundation
struct PermissionInfoDto : Codable {
    var name: String? = nil
    var description: String? = nil
    var completeKey: String? = nil
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
class PermissionInfoDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  @Published var completeKey: String? = nil
  @Published var completeKeyErrorMessage: String? = nil
  func getDto() -> PermissionInfoDto {
      var dto = PermissionInfoDto()
    dto.name = self.name
    dto.description = self.description
    dto.completeKey = self.completeKey
      return dto
  }
}
