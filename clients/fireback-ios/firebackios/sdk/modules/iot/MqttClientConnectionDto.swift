import Foundation
struct MqttClientConnectionDto : Codable {
    var name: String? = nil
    var isConnected: Bool? = nil
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
class MqttClientConnectionDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var isConnected: Bool? = nil
  @Published var isConnectedErrorMessage: Bool? = nil
  func getDto() -> MqttClientConnectionDto {
      var dto = MqttClientConnectionDto()
    dto.name = self.name
    dto.isConnected = self.isConnected
      return dto
  }
}
