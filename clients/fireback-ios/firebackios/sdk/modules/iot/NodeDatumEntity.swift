import Foundation
class NodeDatumEntity : Codable, Identifiable {
    var node: DataNodeEntity? = nil
    // var nodeId: String? = nil
    var valueFloat64: Float64? = nil
    var valueInt64: Int? = nil
    var valueString: String? = nil
    var valueBoolean: Bool? = nil
    var ingestedAt: String? = nil
}
class NodeDatumEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var valueInt64: Int? = nil
  @Published var valueInt64ErrorMessage: Int? = nil
  @Published var valueString: String? = nil
  @Published var valueStringErrorMessage: String? = nil
  @Published var valueBoolean: Bool? = nil
  @Published var valueBooleanErrorMessage: Bool? = nil
  func getDto() -> NodeDatumEntity {
      var dto = NodeDatumEntity()
    dto.valueInt64 = self.valueInt64
    dto.valueString = self.valueString
    dto.valueBoolean = self.valueBoolean
      return dto
  }
}