import Foundation
class DataNodeValues : Codable, Identifiable {
    var key: String? = nil
    var valueInt64: Int? = nil
    var valueFloat64: Float64? = nil
    var valueString: String? = nil
    var valueBoolean: Bool? = nil
    var valueType: String? = nil
    var value: String? = nil
    var readable: Bool? = nil
    var writable: Bool? = nil
}
class DataNodeReaders : Codable, Identifiable {
    var reader: NodeReaderEntity? = nil
    // var readerId: String? = nil
    var config: String? = nil
}
class DataNodeWriters : Codable, Identifiable {
    var writer: NodeWriterEntity? = nil
    // var writerId: String? = nil
    var config: String? = nil
}
class DataNodeEntity : Codable, Identifiable {
    var name: String? = nil
    var expanderFunction: ExpanderFunctionEntity? = nil
    // var expanderFunctionId: String? = nil
    var values: [DataNodeValues]? = nil
    var type: DataNodeTypeEntity? = nil
    // var typeId: String? = nil
    var mode: DataNodeModeEntity? = nil
    // var modeId: String? = nil
    var readers: [DataNodeReaders]? = nil
    var writers: [DataNodeWriters]? = nil
}
class DataNodeEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> DataNodeEntity {
      var dto = DataNodeEntity()
    dto.name = self.name
      return dto
  }
}