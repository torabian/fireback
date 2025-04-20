import Foundation
class TableViewSizingEntity : Codable, Identifiable {
    var tableName: String? = nil
    var sizes: String? = nil
}
class TableViewSizingEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var tableName: String? = nil
  @Published var tableNameErrorMessage: String? = nil
  @Published var sizes: String? = nil
  @Published var sizesErrorMessage: String? = nil
  func getDto() -> TableViewSizingEntity {
      var dto = TableViewSizingEntity()
    dto.tableName = self.tableName
    dto.sizes = self.sizes
      return dto
  }
}