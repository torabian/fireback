import Foundation
class BackupTableMetaEntity : Codable, Identifiable {
    var tableNameInDb: String? = nil
}
class BackupTableMetaEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var tableNameInDb: String? = nil
  @Published var tableNameInDbErrorMessage: String? = nil
  func getDto() -> BackupTableMetaEntity {
      var dto = BackupTableMetaEntity()
    dto.tableNameInDb = self.tableNameInDb
      return dto
  }
}