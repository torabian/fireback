import Foundation
class PersonEntity : Codable, Identifiable {
    var firstName: String? = nil
    var lastName: String? = nil
    var photo: String? = nil
}
class PersonEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  @Published var photo: String? = nil
  @Published var photoErrorMessage: String? = nil
  func getDto() -> PersonEntity {
      var dto = PersonEntity()
    dto.firstName = self.firstName
    dto.lastName = self.lastName
    dto.photo = self.photo
      return dto
  }
}