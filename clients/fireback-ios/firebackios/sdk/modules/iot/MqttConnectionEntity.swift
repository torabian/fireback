import Foundation
class MqttConnectionEntity : Codable, Identifiable {
    var ssl: Bool? = nil
    var autoReconnect: Bool? = nil
    var cleanSession: Bool? = nil
    var lastWillRetain: Bool? = nil
    var port: Int? = nil
    var keepAlive: Int? = nil
    var connectTimeout: Int? = nil
    var lastWillQos: Int? = nil
    var clientId: String? = nil
    var name: String? = nil
    var host: String? = nil
    var username: String? = nil
    var password: String? = nil
    var mqttVersion: MqttVersionEntity? = nil
    // var mqttVersionId: String? = nil
    var lastWillTopic: String? = nil
    var lastWillPayload: String? = nil
}
class MqttConnectionEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var ssl: Bool? = nil
  @Published var sslErrorMessage: Bool? = nil
  @Published var autoReconnect: Bool? = nil
  @Published var autoReconnectErrorMessage: Bool? = nil
  @Published var cleanSession: Bool? = nil
  @Published var cleanSessionErrorMessage: Bool? = nil
  @Published var lastWillRetain: Bool? = nil
  @Published var lastWillRetainErrorMessage: Bool? = nil
  @Published var port: Int? = nil
  @Published var portErrorMessage: Int? = nil
  @Published var keepAlive: Int? = nil
  @Published var keepAliveErrorMessage: Int? = nil
  @Published var connectTimeout: Int? = nil
  @Published var connectTimeoutErrorMessage: Int? = nil
  @Published var lastWillQos: Int? = nil
  @Published var lastWillQosErrorMessage: Int? = nil
  @Published var clientId: String? = nil
  @Published var clientIdErrorMessage: String? = nil
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var host: String? = nil
  @Published var hostErrorMessage: String? = nil
  @Published var username: String? = nil
  @Published var usernameErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  @Published var lastWillTopic: String? = nil
  @Published var lastWillTopicErrorMessage: String? = nil
  @Published var lastWillPayload: String? = nil
  @Published var lastWillPayloadErrorMessage: String? = nil
  func getDto() -> MqttConnectionEntity {
      var dto = MqttConnectionEntity()
    dto.ssl = self.ssl
    dto.autoReconnect = self.autoReconnect
    dto.cleanSession = self.cleanSession
    dto.lastWillRetain = self.lastWillRetain
    dto.port = self.port
    dto.keepAlive = self.keepAlive
    dto.connectTimeout = self.connectTimeout
    dto.lastWillQos = self.lastWillQos
    dto.clientId = self.clientId
    dto.name = self.name
    dto.host = self.host
    dto.username = self.username
    dto.password = self.password
    dto.lastWillTopic = self.lastWillTopic
    dto.lastWillPayload = self.lastWillPayload
      return dto
  }
}