import Foundation
class GpioModeEntity : Codable, Identifiable {
    var key: String? = nil
    var index: Int? = nil
    var description: String? = nil
}
class GpioModeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var key: String? = nil
  @Published var keyErrorMessage: String? = nil
  @Published var index: Int? = nil
  @Published var indexErrorMessage: Int? = nil
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  func getDto() -> GpioModeEntity {
      var dto = GpioModeEntity()
    dto.key = self.key
    dto.index = self.index
    dto.description = self.description
      return dto
  }
}