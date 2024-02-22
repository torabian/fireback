import Foundation
struct TestMailDto : Codable {
    var senderId: String? = nil
    var toName: String? = nil
    var toEmail: String? = nil
    var subject: String? = nil
    var content: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class TestMailDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var senderId: String? = nil
  @Published var senderIdErrorMessage: String? = nil
  @Published var toName: String? = nil
  @Published var toNameErrorMessage: String? = nil
  @Published var toEmail: String? = nil
  @Published var toEmailErrorMessage: String? = nil
  @Published var subject: String? = nil
  @Published var subjectErrorMessage: String? = nil
  @Published var content: String? = nil
  @Published var contentErrorMessage: String? = nil
  func getDto() -> TestMailDto {
      var dto = TestMailDto()
    dto.senderId = self.senderId
    dto.toName = self.toName
    dto.toEmail = self.toEmail
    dto.subject = self.subject
    dto.content = self.content
      return dto
  }
}
