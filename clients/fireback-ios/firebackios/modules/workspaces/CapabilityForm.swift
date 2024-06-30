import SwiftUI
struct CapabilityForm: View {
    @State private var formModel = CapabilityEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.Name.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}