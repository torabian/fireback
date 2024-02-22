import Foundation
struct DnWriteMQTTTopicConfigDto : Codable {
    var topic: String? = nil
    var qos: String? = nil
    var message: String? = nil
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
class DnWriteMQTTTopicConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var topic: String? = nil
  @Published var topicErrorMessage: String? = nil
  @Published var qos: String? = nil
  @Published var qosErrorMessage: String? = nil
  @Published var message: String? = nil
  @Published var messageErrorMessage: String? = nil
  func getDto() -> DnWriteMQTTTopicConfigDto {
      var dto = DnWriteMQTTTopicConfigDto()
    dto.topic = self.topic
    dto.qos = self.qos
    dto.message = self.message
      return dto
  }
}
