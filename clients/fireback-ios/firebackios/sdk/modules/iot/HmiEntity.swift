import Foundation
class HmiComponents : Codable, Identifiable {
    var layoutMode: String? = nil
    var data: String? = nil
    var type: HmiComponentTypeEntity? = nil
    // var typeId: String? = nil
    var label: String? = nil
    var icon: String? = nil
    var readSubKey: String? = nil
    var read: DataNodeEntity? = nil
    // var readId: String? = nil
    var write: DataNodeEntity? = nil
    // var writeId: String? = nil
    var position: HmiComponentsPosition? = nil
    var states: [HmiComponentsStates]? = nil
}
class HmiComponentsPosition : Codable, Identifiable {
    var x: Int? = nil
    var y: Int? = nil
    var width: Int? = nil
    var height: Int? = nil
}
class HmiComponentsStates : Codable, Identifiable {
    var color: String? = nil
    var colorFilter: String? = nil
    var tag: String? = nil
    var label: String? = nil
    var value: String? = nil
}
class HmiEntity : Codable, Identifiable {
    var isRunning: Bool? = nil
    var name: String? = nil
    var layout: String? = nil
    var components: [HmiComponents]? = nil
}
class HmiEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var isRunning: Bool? = nil
  @Published var isRunningErrorMessage: Bool? = nil
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  func getDto() -> HmiEntity {
      var dto = HmiEntity()
    dto.isRunning = self.isRunning
    dto.name = self.name
      return dto
  }
}