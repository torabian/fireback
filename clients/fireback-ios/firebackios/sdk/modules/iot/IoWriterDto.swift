import Foundation
struct IoWriterDto : Codable {
    var content: String? = nil
    var type: String? = nil
    var host: String? = nil
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
class IoWriterDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var content: String? = nil
  @Published var contentErrorMessage: String? = nil
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  @Published var host: String? = nil
  @Published var hostErrorMessage: String? = nil
  @Published var path: String? = nil
  @Published var pathErrorMessage: String? = nil
  func getDto() -> IoWriterDto {
      var dto = IoWriterDto()
    dto.content = self.content
    dto.type = self.type
    dto.host = self.host
    dto.path = self.path
      return dto
  }
}
