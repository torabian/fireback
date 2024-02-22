import Foundation
class ProductPlanPermissions : Codable, Identifiable {
    var capability: CapabilityEntity? = nil
    // var capabilityId: String? = nil
}
class ProductPlanEntity : Codable, Identifiable {
    var name: String? = nil
    var duration: Int? = nil
    var product: LicensableProductEntity? = nil
    // var productId: String? = nil
    var priceTag: PriceTagEntity? = nil
    // var priceTagId: String? = nil
    var permissions: [ProductPlanPermissions]? = nil
}
class ProductPlanEntityViewModel: ObservableObject {
    // improve the fields here
    @Published var name: String? = nil
    @Published var duration: Int? = nil
    func getDto() -> ProductPlanEntity {
        var dto = ProductPlanEntity()
        dto.name = self.name
        dto.duration = self.duration
        return dto
    }
}