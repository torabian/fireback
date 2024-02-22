import Foundation
struct DnReadWriteDataNodeDto : Codable {
    var nodeId: String? = nil
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
class DnReadWriteDataNodeDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var nodeId: String? = nil
  @Published var nodeIdErrorMessage: String? = nil
  func getDto() -> DnReadWriteDataNodeDto {
      var dto = DnReadWriteDataNodeDto()
    dto.nodeId = self.nodeId
      return dto
  }
}
