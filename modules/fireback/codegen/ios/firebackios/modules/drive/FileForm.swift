import SwiftUI
struct FileForm: View {
    @State private var formModel = FileEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.Name.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.DiskPath.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.VirtualPath.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.Type.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}