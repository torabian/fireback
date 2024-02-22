import Foundation
class GpioEntity : Codable, Identifiable {
    var name: String? = nil
    var index: Int? = nil
    var analogFunction: String? = nil
    var rtcGpio: String? = nil
    var comments: String? = nil
    var mode: GpioModeEntity? = nil
    // var modeId: String? = nil
}
class GpioEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var index: Int? = nil
  @Published var indexErrorMessage: Int? = nil
  @Published var analogFunction: String? = nil
  @Published var analogFunctionErrorMessage: String? = nil
  @Published var rtcGpio: String? = nil
  @Published var rtcGpioErrorMessage: String? = nil
  @Published var comments: String? = nil
  @Published var commentsErrorMessage: String? = nil
  func getDto() -> GpioEntity {
      var dto = GpioEntity()
    dto.name = self.name
    dto.index = self.index
    dto.analogFunction = self.analogFunction
    dto.rtcGpio = self.rtcGpio
    dto.comments = self.comments
      return dto
  }
}