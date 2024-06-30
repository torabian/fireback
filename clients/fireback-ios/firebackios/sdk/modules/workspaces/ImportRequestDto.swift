import Foundation
struct ImportRequestDto : Codable {
    var file: String? = nil
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
class ImportRequestDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var file: String? = nil
  @Published var fileErrorMessage: String? = nil
  func getDto() -> ImportRequestDto {
      var dto = ImportRequestDto()
    dto.file = self.file
      return dto
  }
}
