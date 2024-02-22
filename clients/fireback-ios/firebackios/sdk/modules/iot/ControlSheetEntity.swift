import Foundation
class ControlSheetObjects : Codable, Identifiable {
    var width: Float64? = nil
    var height: Float64? = nil
    var type: String? = nil
    var selected: Bool? = nil
    var meta: String? = nil
    var dragging: Bool? = nil
    var id: String? = nil
    var connections: [ControlSheetObjectsConnections]? = nil
    var position: ControlSheetObjectsPosition? = nil
    var positionAbsolute: ControlSheetObjectsPositionAbsolute? = nil
}
class ControlSheetObjectsConnections : Codable, Identifiable {
    var type: String? = nil
    var data: String? = nil
}
class ControlSheetObjectsPosition : Codable, Identifiable {
    var x: Float64? = nil
    var y: Float64? = nil
}
class ControlSheetObjectsPositionAbsolute : Codable, Identifiable {
    var x: Float64? = nil
    var y: Float64? = nil
}
class ControlSheetEdges : Codable, Identifiable {
    var source: String? = nil
    var sourceHandle: String? = nil
    var target: String? = nil
    var targetHandle: String? = nil
    var id: String? = nil
}
class ControlSheetEntity : Codable, Identifiable {
    var isRunning: Bool? = nil
    var name: String? = nil
    var objects: [ControlSheetObjects]? = nil
    var edges: [ControlSheetEdges]? = nil
}
class ControlSheetEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var isRunning: Bool? = nil
  @Published var isRunningErrorMessage: Bool? = nil
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> ControlSheetEntity {
      var dto = ControlSheetEntity()
    dto.isRunning = self.isRunning
    dto.name = self.name
      return dto
  }
}