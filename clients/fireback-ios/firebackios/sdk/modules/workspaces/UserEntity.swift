import Foundation
class UserEntity : Codable, Identifiable {
    var person: PersonEntity? = nil
    // var personId: String? = nil
    var passports: [PassportEntity]? = nil
    var passportsListId: [String]? = nil
}
class UserEntityViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> UserEntity {
      var dto = UserEntity()
      return dto
  }
}