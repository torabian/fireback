import Foundation
struct DataNodeReaderFnDefDto : Codable {
    var fn: String? = nil
    var description: String? = nil
    var reads: String? = nil
    var writes: String? = nil
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
class DataNodeReaderFnDefDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var fn: String? = nil
  @Published var fnErrorMessage: String? = nil
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  @Published var reads: String? = nil
  @Published var readsErrorMessage: String? = nil
  @Published var writes: String? = nil
  @Published var writesErrorMessage: String? = nil
  func getDto() -> DataNodeReaderFnDefDto {
      var dto = DataNodeReaderFnDefDto()
    dto.fn = self.fn
    dto.description = self.description
    dto.reads = self.reads
    dto.writes = self.writes
      return dto
  }
}
