import Foundation
struct DnWriteToHostFileSystemConfigDto : Codable {
    var path: String? = nil
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
class DnWriteToHostFileSystemConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var path: String? = nil
  @Published var pathErrorMessage: String? = nil
  func getDto() -> DnWriteToHostFileSystemConfigDto {
      var dto = DnWriteToHostFileSystemConfigDto()
    dto.path = self.path
      return dto
  }
}
