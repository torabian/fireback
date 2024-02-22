import Foundation
class ActivationKeyEntity : Codable, Identifiable {
    var series: String? = nil
    var used: Int? = nil
    var plan: ProductPlanEntity? = nil
    // var planId: String? = nil
}
class ActivationKeyEntityViewModel: ObservableObject {
    // improve the fields here
    @Published var series: String? = nil
    @Published var used: Int? = nil
    func getDto() -> ActivationKeyEntity {
        var dto = ActivationKeyEntity()
        dto.series = self.series
        dto.used = self.used
        return dto
    }
}