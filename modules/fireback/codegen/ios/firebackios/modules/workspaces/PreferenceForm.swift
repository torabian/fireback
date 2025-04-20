import SwiftUI
struct PreferenceForm: View {
    @State private var formModel = PreferenceEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.Timezone.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}