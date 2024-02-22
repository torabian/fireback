import Foundation
struct DnModbusTcpConfigDto : Codable {
    var timeOut: Int? = nil
    var slaveId: Int? = nil
    var host: String? = nil
    var port: String? = nil
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
class DnModbusTcpConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var timeOut: Int? = nil
  @Published var timeOutErrorMessage: Int? = nil
  @Published var slaveId: Int? = nil
  @Published var slaveIdErrorMessage: Int? = nil
  @Published var host: String? = nil
  @Published var hostErrorMessage: String? = nil
  @Published var port: String? = nil
  @Published var portErrorMessage: String? = nil
  func getDto() -> DnModbusTcpConfigDto {
      var dto = DnModbusTcpConfigDto()
    dto.timeOut = self.timeOut
    dto.slaveId = self.slaveId
    dto.host = self.host
    dto.port = self.port
      return dto
  }
}
