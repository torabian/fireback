import Foundation
class MqttVersionEntity : Codable, Identifiable {
    var version: String? = nil
}
class MqttVersionEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var version: String? = nil
  @Published var versionErrorMessage: String? = nil
  func getDto() -> MqttVersionEntity {
      var dto = MqttVersionEntity()
    dto.version = self.version
      return dto
  }
}