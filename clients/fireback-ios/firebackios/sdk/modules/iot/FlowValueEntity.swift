import Foundation
class FlowValueEntity : Codable, Identifiable {
    var connectionId: String? = nil
    var valueInt: Int? = nil
    var valueString: String? = nil
    var valueFloat: Float64? = nil
    var valueType: Int? = nil
}
class FlowValueEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var connectionId: String? = nil
  @Published var connectionIdErrorMessage: String? = nil
  @Published var valueInt: Int? = nil
  @Published var valueIntErrorMessage: Int? = nil
  @Published var valueString: String? = nil
  @Published var valueStringErrorMessage: String? = nil
  @Published var valueType: Int? = nil
  @Published var valueTypeErrorMessage: Int? = nil
  func getDto() -> FlowValueEntity {
      var dto = FlowValueEntity()
    dto.connectionId = self.connectionId
    dto.valueInt = self.valueInt
    dto.valueString = self.valueString
    dto.valueType = self.valueType
      return dto
  }
}