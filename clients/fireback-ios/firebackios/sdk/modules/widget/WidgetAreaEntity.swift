import Foundation
class WidgetAreaWidgets : Codable, Identifiable {
    var title: String? = nil
    var widget: WidgetEntity? = nil
    // var widgetId: String? = nil
    var x: Int? = nil
    var y: Int? = nil
    var w: Int? = nil
    var h: Int? = nil
    var data: String? = nil
}
class WidgetAreaEntity : Codable, Identifiable {
    var name: String? = nil
    var layouts: String? = nil
    var widgets: [WidgetAreaWidgets]? = nil
}
class WidgetAreaEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var name: String? = nil
  @Published var nameErrorMessage: String? = nil
  @Published var layouts: String? = nil
  @Published var layoutsErrorMessage: String? = nil
  func getDto() -> WidgetAreaEntity {
      var dto = WidgetAreaEntity()
    dto.name = self.name
    dto.layouts = self.layouts
      return dto
  }
}