import Foundation
class DeviceDeviceModbusConfig : Codable, Identifiable {
    var mode: ModbusTransmissionModeEntity? = nil
    // var modeId: String? = nil
    var baudRate: Int? = nil
    var dataBits: Int? = nil
    var parity: Int? = nil
    var stopBit: Int? = nil
    var timeout: Int? = nil
}
class DeviceEntity : Codable, Identifiable {
    var name: String? = nil
    var model: String? = nil
    var ip: String? = nil
    var wifiUser: String? = nil
    var wifiPassword: String? = nil
    var securityType: String? = nil
    var type: DeviceTypeEntity? = nil
    // var typeId: String? = nil
    var deviceModbusConfig: DeviceDeviceModbusConfig? = nil
}
class DeviceEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var model: String? = nil
  @Published var modelErrorMessage: String? = nil
  @Published var ip: String? = nil
  @Published var ipErrorMessage: String? = nil
  @Published var wifiUser: String? = nil
  @Published var wifiUserErrorMessage: String? = nil
  @Published var wifiPassword: String? = nil
  @Published var wifiPasswordErrorMessage: String? = nil
  @Published var securityType: String? = nil
  @Published var securityTypeErrorMessage: String? = nil
  func getDto() -> DeviceEntity {
      var dto = DeviceEntity()
    dto.name = self.name
    dto.model = self.model
    dto.ip = self.ip
    dto.wifiUser = self.wifiUser
    dto.wifiPassword = self.wifiPassword
    dto.securityType = self.securityType
      return dto
  }
}