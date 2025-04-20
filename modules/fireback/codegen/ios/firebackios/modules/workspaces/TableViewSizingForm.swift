import SwiftUI
struct TableViewSizingForm: View {
    @State private var formModel = TableViewSizingEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.TableName.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.Sizes.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}