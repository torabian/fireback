import Foundation
class PersonEntity : Codable, Identifiable {
    var firstName: String? = nil
    var lastName: String? = nil
    var photo: String? = nil
    var gender: String? = nil
    var title: String? = nil
    var birthDate: Date? = nil
}
class PersonEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  @Published var photo: String? = nil
  @Published var photoErrorMessage: String? = nil
  @Published var gender: String? = nil
  @Published var genderErrorMessage: String? = nil
  @Published var title: String? = nil
  @Published var titleErrorMessage: String? = nil
  func getDto() -> PersonEntity {
      var dto = PersonEntity()
    dto.firstName = self.firstName
    dto.lastName = self.lastName
    dto.photo = self.photo
    dto.gender = self.gender
    dto.title = self.title
      return dto
  }
}