import Foundation
class GpioStateEntity : Codable, Identifiable {
    var gpioMode: GpioModeEntity? = nil
    // var gpioModeId: String? = nil
    var gpioIndexSnapshot: Int? = nil
    var gpioModeSnapshot: Int? = nil
    var gpioValueSnapshot: Int? = nil
    var gpio: GpioEntity? = nil
    // var gpioId: String? = nil
}
class GpioStateEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var gpioIndexSnapshot: Int? = nil
  @Published var gpioIndexSnapshotErrorMessage: Int? = nil
  @Published var gpioModeSnapshot: Int? = nil
  @Published var gpioModeSnapshotErrorMessage: Int? = nil
  @Published var gpioValueSnapshot: Int? = nil
  @Published var gpioValueSnapshotErrorMessage: Int? = nil
  func getDto() -> GpioStateEntity {
      var dto = GpioStateEntity()
    dto.gpioIndexSnapshot = self.gpioIndexSnapshot
    dto.gpioModeSnapshot = self.gpioModeSnapshot
    dto.gpioValueSnapshot = self.gpioValueSnapshot
      return dto
  }
}